package classpath

import (
	"path/filepath"
	"os"
	"fmt"
)

type Classpath struct{
	//系统classpath
	bootClasspath Entry
	//扩展类路径
	extendClasspath Entry
	//用户自定义类路径
	userClasspath Entry
}
//解析用户输入的jre参数和classpath
func Parse(jreOption,classpathOption string) *Classpath {
	cp := &Classpath{}
	//分别加载系统类路径和用户类路径的class文件并返回
	cp.ParseCpAndExtendClasspath(jreOption)
	cp.ParseUserClasspath(classpathOption)
	return cp
}
//加载扩展类路径的class文件
func (self *Classpath) ParseCpAndExtendClasspath(jreOption string) {
	jreDir:=getJreDir(jreOption)
	jreLibPath:=filepath.Join(jreDir,"lib","*")
	fmt.Printf("jre path is : %s\n",jreLibPath)
	self.bootClasspath=newWildcardEntry(jreLibPath)//基础类路径
	extLibPath:=filepath.Join(jreDir,"lib","ext","*")
	fmt.Printf("extend path is : %s\n",extLibPath)
	self.extendClasspath=newWildcardEntry(extLibPath)
}
//根据用户输入的jreOption获得基础类路径
func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME");jh!= "" {
		fmt.Printf("JAVA_HOME:%s\n",jh)
		return filepath.Join(jh,"jre")
	}
	panic("Can not find jre folder")
}
func exists(dir string) bool {
	//Stat方法返回参数中dir文件的信息，如果有error，则为*PathError
	if _,err := os.Stat(dir);err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
//根据userJreOption解析用户类路径，如果用户没有输入则默认为当前路径
func (self *Classpath) ParseUserClasspath(classpath string) {
	if classpath=="" {
		classpath="."
	}
	self.userClasspath=newEntry(classpath)
}
func (self *Classpath) ReadClass(className string) ([]byte,Entry,error) {
	className=className+".class"
	if data,entry,err:=self.bootClasspath.readClass(className);err==nil {
		return data,entry,err
	}
	if data,entry,err := self.extendClasspath.readClass(className) ;err==nil {
		return data,entry,err
	}
	return self.userClasspath.readClass(className)
}
func (self *Classpath) String() string {
	return self.userClasspath.String()
}