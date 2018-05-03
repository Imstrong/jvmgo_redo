package main

import (
	"fmt"
	"strings"
	"jvmgo_redo/ch06/classpath"
	"jvmgo_redo/ch06/classfile"
)

func main(){
	cmd:=ParseCmd()
	if cmd.versionFlag {
		fmt.Print("1.7.2u81")
	} else if cmd.helpFlag {
		fmt.Print("不会用拉倒")
	}else {
		startJVM(cmd)
	}
}
func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.jreOption, cmd.cpOption)
	className := strings.Replace(cmd.class, ".", "/", -1)
	cf := loadClass(className, cp)
	mainMethod := getMainMethod(cf)
	if mainMethod != nil {
		interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}
}
func loadClass(className string,cp *classpath.Classpath) *classfile.ClassFile {
	bytes, _, e := cp.ReadClass(className)
	if e!=nil {
		panic(e)
	}
	cf, err := classfile.Parse(bytes)
	if err!=nil {
		panic(err)
	}
	return cf
}
func getMainMethod(classfile *classfile.ClassFile) *classfile.AttrMethodInfo {
	for _,method:=range classfile.Methods() {
		if method.Name()=="main"&&method.Descriptor()=="([Ljava/lang/String;)V" {
			return method
		}
	}
	return nil
}