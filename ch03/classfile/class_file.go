package classfile

import "fmt"

type ClassFile struct {
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags uint16
	thisClass uint16
	superClass uint16
	interfaces []uint16
	fields []*MemberInfo
	methods []*MemberInfo
	attributes []Attribute
}
func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		//如果捕获到异常，进入recover恢复
		if r:= recover();r!=nil {
			var ok bool
			err,ok =r.(error)
			if !ok {
				err=fmt.Errorf("%v",r)
			}
		}
	}()
	cr := &ClassReader{classData}
	cf=&ClassFile{}
	cf.read(cr)
	return
}
func (self *ClassFile) read(reader *ClassReader) {}
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {}
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {}
func (self *ClassFile) MinorVersion() uint16 {} // getter
func (self *ClassFile) MajorVersion() uint16 {} // getter
func (self *ClassFile) ConstantPool() ConstantPool {} // getter
func (self *ClassFile) AccessFlags() uint16 {} // getter
func (self *ClassFile) Fields() []*MemberInfo {} // getter
func (self *ClassFile) Methods() []*MemberInfo {} // getter
func (self *ClassFile) ClassName() string {}
func (self *ClassFile) SuperClassName() string {}
func (self *ClassFile) InterfaceNames() []string {}
