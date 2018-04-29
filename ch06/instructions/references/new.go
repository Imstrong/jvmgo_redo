package references

import (
	"jvmgo_redo/ch06/instructions/base"
	"jvmgo_redo/ch06/runtime"
	"jvmgo_redo/ch06/runtime/heap"
)

//new指令，创建对象
type NEW struct {
	base.Index16Instruction
}
//通过来自字节码的16位索引，到当前类的运行时常量池中找到一个类符号引用，并解析得到类数据，创建对象并把引用推入栈顶
func (self *NEW) Execute(frame *runtime.Frame) {
	//从当前线程的栈帧获取方法索引->类索引->类信息->常量池
	cp:=frame.Method().Class().ConstantPool()
	//从常量池中获取类索引
	classRef:=cp.GetConstant(self.Index).(*heap.ClassRef)
	//根据类索引封装Class对象
	class:=classRef.ResolvedClass()
	//如果是接口或抽象类，抛出异常
	if class.IsInterface()||class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}
	//创建对象
	ref:=class.NewObject()
	//将对象引用推入操作数栈顶
	frame.OperandStack().PushRef(ref)
}
