// Package ir - coroutine instruction definitions
package ir

import (
    "fmt"
)

// CoroIdInst represents llvm.coro.id intrinsic
type CoroIdInst struct {
    BaseInstruction
}

func (i *CoroIdInst) String() string {
    return fmt.Sprintf("%%%s = call token @llvm.coro.id(i32 0, ptr null, ptr null, ptr null)", i.ValName)
}

// CoroBeginInst represents llvm.coro.begin intrinsic
type CoroBeginInst struct {
    BaseInstruction
}

func (i *CoroBeginInst) String() string {
    id := i.Ops[0]
    return fmt.Sprintf("%%%s = call ptr @llvm.coro.begin(token %s, ptr null)", i.ValName, formatOp(id))
}

// CoroSuspendInst represents llvm.coro.suspend intrinsic (await point)
type CoroSuspendInst struct {
    BaseInstruction
    Final bool // true for final suspend
}

func (i *CoroSuspendInst) String() string {
    finalFlag := "false"
    if i.Final {
        finalFlag = "true"
    }
    return fmt.Sprintf("%%%s = call i8 @llvm.coro.suspend(token none, i1 %s)", i.ValName, finalFlag)
}

// CoroEndInst represents llvm.coro.end intrinsic
type CoroEndInst struct {
    BaseInstruction
}

func (i *CoroEndInst) String() string {
    handle := i.Ops[0]
    return fmt.Sprintf("call i1 @llvm.coro.end(ptr %s, i1 false)", formatOp(handle))
}

// CoroFreeInst represents llvm.coro.free intrinsic
type CoroFreeInst struct {
    BaseInstruction
}

func (i *CoroFreeInst) String() string {
    id := i.Ops[0]
    handle := i.Ops[1]
    return fmt.Sprintf("%%%s = call ptr @llvm.coro.free(token %s, ptr %s)", i.ValName, formatOp(id), formatOp(handle))
}