package classfile
type DeprecatedAttribute struct {
	MarkerAttribute
}
type SyntheticAttribute struct {
	MarkerAttribute
}
type MarkerAttribute struct{}
func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	//这两个属性使用时只判断其有没有，值是空的，因此实现方法是空的
}
