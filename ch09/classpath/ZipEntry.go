package classpath

import (
	"path/filepath"
	"io/ioutil"
	"archive/zip"
	"errors"
)

type ZipEntry struct {
	fullPath string
}
func newZipEntry(path string) *ZipEntry {
	absPath,err :=filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}
func (self ZipEntry) readClass(className string) ([]byte,Entry,error) {
	reader,err := zip.OpenReader(self.fullPath)
	if err !=nil {
		return nil,nil,err
	}
	defer reader.Close()
	for _,f:=range reader.File {
		if f.Name==className {
			rc,err := f.Open()
			if err != nil {
				return nil,nil,err
			}
			defer rc.Close()
			data,err := ioutil.ReadAll(rc)
			if err != nil {
				return nil,nil,err
			}
			return data,self,nil
		}
	}
	return nil,nil,errors.New("class not found: "+className)
}
func (self ZipEntry) String() string {
	return self.fullPath
}

