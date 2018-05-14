package classpath

import (
	"strings"
	"os"
)
const pathListSeparator=string(os.PathListSeparator)
type Entry interface{
	readClass(path string) ([]byte,Entry,error)
	String() string
}
func newEntry(path string) Entry {
	//多层目录
	if strings.Contains(path,pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path,"*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path,".jar")||strings.HasSuffix(path,".JAR")||strings.HasSuffix(path,".ZIP")||strings.HasSuffix(path,".zip") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}
