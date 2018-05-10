package classpath

import (
	"path/filepath"
	"io/ioutil"
)

type DirEntry struct {
	absPath string
}
func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}
func (self DirEntry) readClass(className string) ([]byte,Entry,error) {
	fullPath := filepath.Join(self.absPath, className)
	data,err:=ioutil.ReadFile(fullPath)
	if err!=nil {
		return nil,nil,err
	}
	return data,self,nil
}
func (self DirEntry) String() string {
	return self.absPath
}