package references

import (
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/runtime"
	"jvmgo/ch08/runtime/heap"
)

type INVOKE_SPECIAL struct{ base.Index16Instruction }
// hack!
func (self *INVOKE_SPECIAL) Execute(frame *runtime.Frame) {
	currentClass:=frame.Method().Class()
	cp:=currentClass.ConstantPool()
	methodRef:=cp.GetConstant(self.Index).(*heap.MethodRef)
	resolvedClass:=methodRef.ResolvedClass()
	resolvedMethod:=methodRef.ResolvedMethod()
	//如果是构造方法，
	if resolvedMethod.Name()=="<init>"&&resolvedMethod.Class()!=resolvedClass {
		panic("java.lang.NoSuchMethodError")
	}
	if resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	//从操作数栈弹出this引用，如果为空，抛出NullPointerException
	ref:=frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount())
	if ref==nil {
		panic("java.lang.NullPointerException")
	}
	if resolvedMethod.IsProtected() && resolvedMethod.Class().IsSuperClassOf(currentClass)&&resolvedMethod.Class().GetPackageName()!=currentClass.GetPackageName()&&ref.Class()!=currentClass&&!ref.Class().IsSubClassOf(currentClass) {
		panic("java.lang.IllegalAccessError")
	}
	methodTobeInvoked:=resolvedMethod
	if currentClass.IsSuper()&&resolvedClass.IsSuperClassOf(currentClass)&&resolvedMethod.Name()!="<init>" {
		methodTobeInvoked=heap.LookupMethodInClass(currentClass.SuperClass(),methodRef.Name(),methodRef.Descriptor())
	}
	if methodTobeInvoked==nil||methodTobeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	base.InvokeMethod(frame,methodTobeInvoked)
}
