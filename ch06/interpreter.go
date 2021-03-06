package main

import (
	"jvmgo_redo/ch06/runtime"
	"fmt"
	"jvmgo_redo/ch06/instructions/base"
	"jvmgo_redo/ch06/instructions"
	"jvmgo_redo/ch06/runtime/heap"
)

func interpret(method *heap.Method) {
	thread := runtime.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)
	defer catchErr(frame)
	loop(thread, method.Code())
}
func catchErr(frame *runtime.Frame) {
	if r:=recover();r!=nil {
		fmt.Printf("localVars;%v\n",frame.LocalVars())
		fmt.Printf("OperandStack:%v\n",frame.OperandStack())
		panic(r)
	}
}
func loop(thread *runtime.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}
	for {
		pc := frame.NextPC()
		thread.SetPC(pc)
		// decode
		reader.Reset(bytecode, pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())
		// execute
		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}
}