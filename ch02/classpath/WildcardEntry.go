package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

func newWildcardEntry(path string) CompositeEntry {
	//去掉路径末尾的*，得到基础路径
	baseDir := path[:len(path)-1]
	compositeEntry := []Entry{}
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		//以当前传入目录为根目录，如果根目录下该子目录为文件夹，则跳过
		if info.IsDir()&&path!=baseDir {
			return filepath.SkipDir
		}
		//如果根目录下文件为jar文件，创建ZipEntry并放到CompositeEntry中
		if strings.HasSuffix(path,".jar")||strings.HasSuffix(path,".JAR") {
			jarEntry:=newZipEntry(path)
			compositeEntry=append(compositeEntry,jarEntry)
		}
		return nil
	}
	//walk方法对一个树形结构目录的每个节点调用walkFn方法，如果是文件夹，则跳过；如果是jar文件，就创建Entry对象并放入CompositeEntry中
	filepath.Walk(baseDir,walkFn)
	return compositeEntry
}
