package runtime
type Frame struct {
	//下一帧的指针
	lower *Frame
	//局部变量表引用
	localVars LocalVars
	//	操作数栈指针
	operandStack *OperandStack
}
func NewFrame(maxLocals,maxStack uint) *Frame {
	return &Frame{
		//maxLocals定义了局部变量表大小maxStack定义操作数栈深度，由编译器计算并保存在方法表的Code属性中，在上一章已经能拿到了
		localVars:newLocalVars(maxLocals),
		operandStack:newOperandStack(maxStack),
	}
}
func (self Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self Frame) OperandStack() *OperandStack {
	return self.operandStack
}




