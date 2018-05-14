package heap

import "unicode/utf16"

//字符串池，string为go字符串，object为java字符串
var internedStrings=map[string]*Object{}
//根据go字符串查询字符串池，如果存在则返回，否则先把go字符串转成java字符数组（utf16），
//然后创建一个String实例，把value设置为字符串数组，再把实例放入字符串池
func JString(loader *ClassLoader,goStr string) *Object {
	if internedStr,ok:=internedStrings[goStr];ok {
		return internedStr
	}
	chars:=stringToUtf16(goStr)
	jChars:=&Object{loader.LoadClass("[C"),chars}
	jStr:=loader.LoadClass("java/lang/String").NewObject()
	jStr.SetRefVal("value","[C",jChars)
	internedStrings[goStr]=jStr
	return jStr
}
func stringToUtf16(str string) []uint16{
	//type rune=int32
	runes:=[]rune(str)
	return utf16.Encode(runes)
}

