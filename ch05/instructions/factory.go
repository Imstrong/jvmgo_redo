package instructions

import (
	"jvmgo/ch05/instructions/base"
	"fmt"
	. "jvmgo/ch05/instructions/constants"
)

func NewInstruction(opcode byte) base.Instruction {
	switch opcode {
	case 0x00:
		return &NOP{}
	case 0x01:
		return &ACONST_NULL{}
		...
	default:
		panic(fmt.Errorf("Unsupported opcode: 0x%x!", opcode))
	}
}
