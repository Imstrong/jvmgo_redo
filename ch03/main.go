package main

import (
	"fmt"
	"jvmgo_redo/ch03/classpath"
	"strings"
	"jvmgo_redo/ch03/classfile"
)

func main(){
	cmd:=ParseCmd()
	if cmd.versionFlag {
		fmt.Print("1.7.2u81")
	} else if cmd.helpFlag {
		fmt.Print("不会用拉倒")
	}else {
		fmt.Printf("enter startJVM...:%v\n",cmd)
		startJVM(cmd)
	}
}
func startJVM(cmd *Cmd) {
	cp :=classpath.Parse(cmd.jreOption,cmd.cpOption)
	fmt.Print("classpath:%v class:%v args:%v\n",cmd.jreOption,cmd.cpOption,cmd.class,cmd.args)
	className:=strings.Replace(cmd.class,".","/",-1)
	cf:=loadClass(className,cp)
	printClassInfo(cf)
}
func loadClass(className string,cp *classpath.Classpath) *classfile.ClassFile{
	classData,_,err:=cp.ReadClass(className)
	if err !=nil {
		panic(err)
	}
	cf,err:=classfile.Parse(classData)
	if err!=nil {
		panic(err)
	}
	return cf
}
func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constants count: %v\n", len(cf.ConstantPool()))
	fmt.Printf("access flags: 0x%x\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf(" %s\n", f.Name())
	}
	fmt.Printf("methods count: %v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf(" %s\n", m.Name())
	}
}