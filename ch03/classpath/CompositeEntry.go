package classpath

import (
	"strings"
	"errors"
)

type CompositeEntry []Entry
func newCompositeEntry(path string) CompositeEntry {
	compositeEntry:=[]Entry{}
	for _,path:=range strings.Split(path,pathListSeparator) {
		entry := newEntry(path)
		compositeEntry=append(compositeEntry,entry)
	}
	return compositeEntry
}
func (self CompositeEntry) readClass(className string) ([]byte,Entry,error) {
	for _,entry := range self {
		data,e,err := entry.readClass(className)
		if err == nil {
			return data,e,nil
		}
	}
	return nil,nil,errors.New("class not found :"+className)
}
func (self CompositeEntry) String() string {
	strs := make([]string, len(self))
	for i,entry:= range self {
		strs[i]=entry.String()
	}
	return strings.Join(strs,pathListSeparator)
}

