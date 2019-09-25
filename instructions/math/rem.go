package math

import (
	"math"

	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Remainder double
type DREM struct{ base.NoOperandsInstruction }

func (instr *DREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := math.Mod(v1, v2) // todo
	stack.PushDouble(result)
}

// Remainder float
type FREM struct{ base.NoOperandsInstruction }

func (instr *FREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := float32(math.Mod(float64(v1), float64(v2))) // todo
	stack.PushFloat(result)
}

// Remainder int
type IREM struct{ base.NoOperandsInstruction }

func (instr *IREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()

	if v2 == 0 {
		frame.Thread().ThrowDivByZero()
	} else {
		result := v1 % v2
		stack.PushInt(result)
	}
}

// Remainder long
type LREM struct{ base.NoOperandsInstruction }

func (instr *LREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()

	if v2 == 0 {
		frame.Thread().ThrowDivByZero()
	} else {
		result := v1 % v2
		stack.PushLong(result)
	}
}