package heap

import "jvmgo_redo/ch08/classfile"

type Field struct {
	//标识该字段在字段表中的位置，字段表为一个[]*Slot，该值在计算字段个数时设置
	slotId uint
	constValueIndex uint
	ClassMember
}
func newFields(class *Class,cfFields []*classfile.AttrMethodInfo) []*Field {
	fields:=make([]*Field,len(cfFields))
	for i,cfField:=range cfFields {
		fields[i]=&Field{}
		fields[i].class=class
		fields[i].copyMemberInfo(cfField)
		fields[i].copyAttributes(cfField)
	}
	return fields
}
func (self *Field) isLongOrDouble() bool {
	return self.descriptor=="J"||self.descriptor=="D"
}
func (self *Field) copyAttributes(cfField *classfile.AttrMethodInfo) {
	if valAttr:=cfField.ConstantValueIndex();valAttr!=nil {
		self.constValueIndex=uint(valAttr.ConstantValueIndex())
	}
}
func (self *Field) ConstantValueIndex() uint {
	return self.constValueIndex
}
func (self *Field) IsStatic() bool {
	if self.accessFlags&ACC_STATIC!=0 {
		return true
	}
	return false
}
func (self *Field) Class() *Class {
	return self.class
}
func (self *Field) SlotId() uint {
	return self.slotId
}