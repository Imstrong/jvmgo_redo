package heap
func getArrayClassName(className string) string {
	return "["+toDescriptor(className)
}
func toDescriptor(className string) string {
	//如果是数组类，直接返回类名
	if className[0]=='[' {
		return className
	}
	//如果是基本类型，返回对应类型描述符
	if d,ok:=primitiveTypes[className];ok{
		return d
	}
	return "L"+className+";"
}
//八种基本数据类型，以及void
var primitiveTypes=map[string]string {
	"void":"V",
	"boolean":"Z",
	"byte":"B",
	"short":"S",
	"int":"I",
	"long":"J",
	"float":"F",
	"double":"D",
	"char":"C",
}
func getComponentClassName(className string) string {
	if className[0]=='[' {
		componentTypeDescriptor:=className[1:]
		return toClassName(componentTypeDescriptor)
	}
	panic("Not array: "+className)
}
func toClassName(descriptor string) string {
	if descriptor[0]=='[' {
		return descriptor
	}
	if descriptor[0]=='L' {
		return descriptor[1:len(descriptor)-1]
	}
	for className,d:=range primitiveTypes {
		if d==descriptor {
			return className
		}
	}
	panic("Invalid descriptor: "+descriptor)
}