// --- START OF FILE builder/ir/coroutines.go ---
package ir

import (
    "fmt"
    "github.com/arc-language/arc-lang/builder/types"
)

// CoroIdInst represents coro.id intrinsic
type CoroIdInst struct {
    BaseInstruction
}

func (i *CoroIdInst) String() string {
    return fmt.Sprintf("%%%s = call token coro.id(i32 0, ptr null, ptr null, ptr null)", i.ValName)
}

// CoroBeginInst represents coro.begin intrinsic
type CoroBeginInst struct {
    BaseInstruction
}

func (i *CoroBeginInst) String() string {
    id := i.Ops[0]
    return fmt.Sprintf("%%%s = call ptr coro.begin(token %s, ptr null)", i.ValName, formatOp(id))
}

// CoroSuspendInst represents coro.suspend intrinsic (await point)
type CoroSuspendInst struct {
    BaseInstruction
    Final bool // true for final suspend
}

func (i *CoroSuspendInst) String() string {
    finalFlag := "false"
    if i.Final {
        finalFlag = "true"
    }
    return fmt.Sprintf("%%%s = call i8 coro.suspend(token none, i1 %s)", i.ValName, finalFlag)
}

// CoroEndInst represents coro.end intrinsic
type CoroEndInst struct {
    BaseInstruction
}

func (i *CoroEndInst) String() string {
    handle := i.Ops[0]
    return fmt.Sprintf("call i1 coro.end(ptr %s, i1 false)", formatOp(handle))
}

// CoroFreeInst represents coro.free intrinsic
type CoroFreeInst struct {
    BaseInstruction
}

func (i *CoroFreeInst) String() string {
    id := i.Ops[0]
    handle := i.Ops[1]
    return fmt.Sprintf("%%%s = call ptr coro.free(token %s, ptr %s)", i.ValName, formatOp(id), formatOp(handle))
}

// --- New Instructions ---

// CoroResumeInst represents coro.resume intrinsic
type CoroResumeInst struct {
    BaseInstruction
}

func (i *CoroResumeInst) String() string {
    handle := i.Ops[0]
    return fmt.Sprintf("call void coro.resume(ptr %s)", formatOp(handle))
}

// CoroDestroyInst represents coro.destroy intrinsic
type CoroDestroyInst struct {
    BaseInstruction
}

func (i *CoroDestroyInst) String() string {
    handle := i.Ops[0]
    return fmt.Sprintf("call void coro.destroy(ptr %s)", formatOp(handle))
}

// CoroPromiseInst represents coro.promise intrinsic (get promise from handle)
type CoroPromiseInst struct {
    BaseInstruction
}

func (i *CoroPromiseInst) String() string {
    handle := i.Ops[0]
    align := i.Ops[1]
    fromStart := i.Ops[2]
    return fmt.Sprintf("%%%s = call ptr coro.promise(ptr %s, i32 %s, i1 %s)", 
        i.ValName, formatOp(handle), formatOp(align), formatOp(fromStart))
}

// CoroDoneInst represents coro.done intrinsic
type CoroDoneInst struct {
    BaseInstruction
}

func (i *CoroDoneInst) String() string {
    handle := i.Ops[0]
    return fmt.Sprintf("%%%s = call i1 coro.done(ptr %s)", i.ValName, formatOp(handle))
}