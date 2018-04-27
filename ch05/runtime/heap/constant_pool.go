package heap
type Constant struct {
	class *Class
	consts []Constant
}
func (self *Constant) GetConstant(index uint) Constant{
	if c:=self.consts[index];c!=nil {

	}
	return c
}
type ConstantPool struct {

}
