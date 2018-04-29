package heap

import "jvmgo/ch06/classfile"

type Field struct {
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
	if valAttr:=cfField.ConstantValueAttribute();valAttr!=nil {
		self.constValueIndex=uint(valAttr.ConstantValueIndex())
	}
}