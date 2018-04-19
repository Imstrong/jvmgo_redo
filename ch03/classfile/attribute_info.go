package classfile
type AttributeInfo interface {
	readInfo(reader *ClassReader)
}
func readAttributes(reader *ClassReader,constantPool ConstantPool) []AttributeInfo {
	attributesCount:=reader.readUint16()
	attributes:=make([]AttributeInfo,attributesCount)
	for i := range attributes {
		attributes[i]=readAttribute(reader,constantPool)
	}
	return attributes
}
func readAttribute(reader *ClassReader,constantPool ConstantPool) AttributeInfo {
	attrNameIndex := reader.readUint16()
	attrName:= constantPool.getUtf8(attrNameIndex)
	attrLength:=reader.readUint32()
	attrInfo:=newAttributeInfo(attrName,attrLength,constantPool)
	return attrInfo
}
func newAttributeInfo(attrName string,attrLength uint32,constantPool ConstantPool) AttributeInfo{
	switch attrName {
	case "Code":
		return &CodeAttribute{cp:constantPool}
	case "ConstantValue":
		return &ConstantValueAttr{}
	case "Deprecated":
		return &DeprecatedAttribute{}
	case "Exceptions":
		return &ExceptionAttribute{}
	case "LineNumberTable":
		return &LineNumTableAttribute{}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{}
	case "SourceFile":
		return &SourcefileAttribute{}
	case "Synthetic":
		return &SyntheticAttribute{}
	default:
		return &UnparsedAttribute{}
	}
}
