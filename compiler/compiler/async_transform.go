package compiler

import (
	"fmt"
	
	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
)

// AsyncTransformer converts async functions into state machines
type AsyncTransformer struct {
	module *ir.Module
	logger *Logger
}

// NewAsyncTransformer creates a new async transformer
func NewAsyncTransformer(module *ir.Module) *AsyncTransformer {
	return &AsyncTransformer{
		module: module,
		logger: NewLogger("[AsyncTransform]"),
	}
}

// Transform processes all async functions in the module
func (t *AsyncTransformer) Transform() error {
	asyncFunctions := 0
	
	for _, fn := range t.module.Functions {
		// Check if function has coroutine attribute
		if t.isAsyncFunction(fn) {
			asyncFunctions++
			t.logger.Info("Transforming async function: %s", fn.Name())
			if err := t.transformAsyncFunction(fn); err != nil {
				return fmt.Errorf("failed to transform %s: %w", fn.Name(), err)
			}
		}
	}
	
	if asyncFunctions == 0 {
		t.logger.Debug("No async functions found in module")
	} else {
		t.logger.Info("Transformed %d async function(s)", asyncFunctions)
	}
	
	return nil
}

// isAsyncFunction checks if a function is async
func (t *AsyncTransformer) isAsyncFunction(fn *ir.Function) bool {
	for _, attr := range fn.Attributes {
		if attr == ir.AttrCoroutine {
			return true
		}
	}
	return false
}

// transformAsyncFunction converts an async function to a state machine
func (t *AsyncTransformer) transformAsyncFunction(fn *ir.Function) error {
	t.logger.Debug("Analyzing function %s for suspend points", fn.Name())
	
	// 1. Find all suspend points (OpCoroSuspend instructions)
	suspendPoints := t.findSuspendPoints(fn)
	if len(suspendPoints) == 0 {
		t.logger.Debug("No suspend points found in %s (async function without await)", fn.Name())
		// Even without suspend points, we need to set up the coroutine frame
		// This allows the function to return a future/promise
		return t.transformAsyncFunctionWithoutSuspend(fn)
	}
	
	t.logger.Debug("Found %d suspend point(s)", len(suspendPoints))
	
	// 2. Analyze live variables at each suspend point
	liveVars := t.analyzeLiveVariables(fn, suspendPoints)
	
	// 3. Create coroutine frame structure
	frameType := t.createFrameType(fn, liveVars)
	t.logger.Debug("Created frame type with %d fields", len(frameType.Fields))
	
	// 4. Split function into resume functions
	t.splitIntoResumeFunctions(fn, suspendPoints, frameType)
	
	// 5. Generate dispatch logic
	t.generateDispatchLogic(fn, len(suspendPoints))
	
	return nil
}

// transformAsyncFunctionWithoutSuspend handles async functions that don't await
func (t *AsyncTransformer) transformAsyncFunctionWithoutSuspend(fn *ir.Function) error {
	t.logger.Debug("Transforming async function %s without suspend points", fn.Name())
	
	// This function is async but doesn't call await
	// It should still return a completed future/promise
	// For now, we just mark it and let it execute synchronously
	
	t.logger.Warning("Async function %s has no await calls - will execute synchronously", fn.Name())
	return nil
}

// findSuspendPoints locates all CoroSuspend instructions
func (t *AsyncTransformer) findSuspendPoints(fn *ir.Function) []*ir.CoroSuspendInst {
	var points []*ir.CoroSuspendInst
	
	for _, block := range fn.Blocks {
		for _, inst := range block.Instructions {
			if suspend, ok := inst.(*ir.CoroSuspendInst); ok {
				points = append(points, suspend)
				t.logger.Debug("Found suspend point in block %s", block.Name())
			}
		}
	}
	
	return points
}

// analyzeLiveVariables determines which variables are live at each suspend point
func (t *AsyncTransformer) analyzeLiveVariables(fn *ir.Function, suspendPoints []*ir.CoroSuspendInst) map[*ir.CoroSuspendInst][]ir.Value {
	liveVars := make(map[*ir.CoroSuspendInst][]ir.Value)
	
	// Simple analysis: find all allocas and consider them potentially live
	// A full implementation would do dataflow analysis
	var allocas []ir.Value
	
	if len(fn.Blocks) > 0 {
		entryBlock := fn.Blocks[0]
		for _, inst := range entryBlock.Instructions {
			if alloca, ok := inst.(*ir.AllocaInst); ok {
				allocas = append(allocas, alloca)
				t.logger.Debug("Found alloca: %s (type: %v)", alloca.Name(), alloca.AllocatedType)
			}
		}
	}
	
	// For now, assume all allocas are live at all suspend points
	// A better implementation would do proper liveness analysis
	for _, suspend := range suspendPoints {
		liveVars[suspend] = allocas
		t.logger.Debug("Suspend point has %d live variables", len(allocas))
	}
	
	return liveVars
}

// createFrameType creates the coroutine frame structure
func (t *AsyncTransformer) createFrameType(fn *ir.Function, liveVars map[*ir.CoroSuspendInst][]ir.Value) *types.StructType {
	// Frame structure:
	// struct CoroFrame {
	//   void* resume_fn;     // Function pointer to resume (offset 0)
	//   void* destroy_fn;    // Cleanup function (offset 8)
	//   int32 state;         // Current state (offset 16) - 0=initial, 1+=resume points
	//   int32 padding;       // Alignment padding (offset 20)
	//   ... saved variables ...  (offset 24+)
	// }
	
	fields := []types.Type{
		types.NewPointer(types.Void), // resume_fn (8 bytes)
		types.NewPointer(types.Void), // destroy_fn (8 bytes)
		types.I32,                     // state (4 bytes)
		types.I32,                     // padding for alignment (4 bytes)
	}
	
	// Add fields for each unique live variable
	seenVars := make(map[ir.Value]bool)
	varCount := 0
	
	for _, vars := range liveVars {
		for _, v := range vars {
			if !seenVars[v] {
				if alloca, ok := v.(*ir.AllocaInst); ok {
					fields = append(fields, alloca.AllocatedType)
					seenVars[v] = true
					varCount++
					t.logger.Debug("Added field for variable %s (type: %v)", alloca.Name(), alloca.AllocatedType)
				}
			}
		}
	}
	
	frameName := fn.Name() + ".frame"
	frameType := types.NewStruct(frameName, fields, false)
	
	t.logger.Info("Created coroutine frame type '%s' with %d variable field(s)", frameName, varCount)
	
	return frameType
}

// splitIntoResumeFunctions splits the function at each suspend point
func (t *AsyncTransformer) splitIntoResumeFunctions(fn *ir.Function, suspendPoints []*ir.CoroSuspendInst, frameType *types.StructType) {
	// This is the most complex part of the transformation
	// For each suspend point, we need to:
	// 1. Save all live variables to the frame
	// 2. Update the state field
	// 3. Return control to caller
	// 4. Create a resume entry point that restores state
	
	t.logger.Debug("Splitting function %s into state machine with %d states", fn.Name(), len(suspendPoints)+1)
	
	// State 0: initial entry
	// State 1: resume after first suspend
	// State 2: resume after second suspend
	// ... etc
	
	for i, suspend := range suspendPoints {
		stateNum := i + 1
		t.logger.Debug("Processing suspend point %d (state %d)", i, stateNum)
		
		// In a full implementation:
		// - Insert code before suspend to save locals to frame
		// - Insert code after suspend to restore locals from frame
		// - Create resume blocks
		
		// For now, just log that we identified the split point
		block := suspend.Parent()
		t.logger.Debug("Suspend point in block %s", block.Name())
	}
	
	t.logger.Warning("State machine code generation not fully implemented - coroutine will use simplified runtime model")
}

// generateDispatchLogic creates the state machine dispatcher
func (t *AsyncTransformer) generateDispatchLogic(fn *ir.Function, numStates int) {
	// Generate a switch statement at function entry that jumps to the right resume point
	// based on frame->state
	//
	// Pseudocode:
	// entry:
	//   frame = load_or_allocate_frame()
	//   state = frame->state
	//   switch(state) {
	//     case 0: goto initial
	//     case 1: goto resume1
	//     case 2: goto resume2
	//     default: unreachable
	//   }
	
	t.logger.Debug("Generating dispatch logic for %d states", numStates+1)
	t.logger.Warning("Dispatch logic generation not fully implemented")
}