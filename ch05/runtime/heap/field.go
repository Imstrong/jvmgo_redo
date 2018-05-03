package heap

import "jvmgo_redo_redo/ch05/classfile"

type Field struct {
	ClassMember
}
func newFields(class *Class,cfFields []*classfile.AttrMethodInfo) []*Field {
	fields:=make([]*Field,len(cfFields))
	for i,cfField:=range cfFields {
		fields[i]=&Field{}
		fields[i].class=class
		fields[i].copyMemberInfo(cfField)
	}
	return fields
}
