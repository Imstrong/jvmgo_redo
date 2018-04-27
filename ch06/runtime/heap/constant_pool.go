package heap

import (
	"jvmgo/ch06/classfile"
	"fmt"
)

type ConstantPool struct {
	class *Class
	consts []Constant
}
type Constant interface {

}
func (self *ConstantPool) GetConstant(index uint) Constant{
	if c:=self.consts[index];c!=nil {
		return c
	}
	panic(fmt.Sprintf("no constants at index %v",index))
}
func newConstantPool(class *Class,cp classfile.ConstantPool) *ConstantPool {
	cpCount:=len(cp)
	consts:=make([]Constant,cpCount)
	constantPool:=&ConstantPool{class,consts}
	for i:=1;i<cpCount;i++ {
		cpInfo:=cp[i]
		switch cpInfo.(type) {
		case *classfile.ConstantIntegerInfo:
			intInfo:=cpInfo.(*classfile.ConstantIntegerInfo)
			consts[i]=intInfo.Value()
		case *classfile.ConstantFloatInfo:
			floatInfo:=cpInfo.(*classfile.ConstantFloatInfo)
			consts[i]=floatInfo.Value()
		case *classfile.ConstantLongInfo:
			longInfo:=cpInfo.(*classfile.ConstantLongInfo)
			consts[i]=longInfo.Value()
			i++
		case *classfile.ConstantDoubleInfo:
			doubleInfo:=cpInfo.(*classfile.ConstantDoubleInfo)
			consts[i]=doubleInfo.Value()
			i++
		case *classfile.ConstantStringInfo:
			stringInfo:=cpInfo.(*classfile.ConstantStringInfo)
			consts[i]=stringInfo.String()
		case *classfile.ConstantClassInfo:
			classInfo:=cpInfo.(*classfile.ConstantClassInfo)
			consts[i]=newClassRef(cp,classInfo)
		case *classfile.ConstantFieldrefInfo:
			fieldrefInfo:=cpInfo.(*classfile.ConstantFieldrefInfo)
			consts[i]=newFieldref(cp,fieldrefInfo)
		case *classfile.ConstantMethodrefInfo:
			methodrefInfo:=cpInfo.(*classfile.ConstantMethodrefInfo)
			consts[i]=newMethodref(cp,methodrefInfo)
		case *classfile.ConstantInterfaceMethodrefInfo:
			interfacemethodrefInfo:=cpInfo.(*classfile.ConstantInterfaceMethodrefInfo)
			consts[i]=newInterfaceMethodref(cp,interfacemethodrefInfo)
		}
	}

}