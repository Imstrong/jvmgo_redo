package base

import (
	"jvmgo_redo/ch08/runtime"
	"jvmgo_redo/ch08/runtime/heap"
	"fmt"
)

//方法执行逻辑
//不管执行的是什么方法，是类方法还是实例方法，执行逻辑都是一样的，只是操作数不同，实例方法要传入this指针作为其中一个操作数而已
func InvokeMethod(invokerFrame *runtime.Frame,method *heap.Method) {
	//为这个函数创建新的栈帧并推入虚拟机栈
	thread:=invokerFrame.Thread()
	newFrame:=thread.NewFrame(method)
	thread.PushFrame(newFrame)
	//传递参数
	argSlot:=int(method.ArgSlotCount())
	if argSlot>0 {
		for i:=argSlot-1;i>=0;i-- {
			slot:=invokerFrame.OperandStack().PopSlot()
			newFrame.LocalVars().SetSlot(uint(i),slot)
		}
	}
	// hack!
	if method.IsNative() {
		if method.Name() == "registerNatives" {
			thread.PopFrame()
		} else {
			panic(fmt.Sprintf("native method: %v.%v%v\n",
				method.Class().Name(), method.Name(), method.Descriptor()))
		}
	}
}
