package main

import (
	"fmt"
	"jvmgo_redo/ch02/classpath"
	"strings"
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
	classData,_,err:=cp.ReadClass(className)
	if err !=nil {
		fmt.Printf("Could not find or load main class %s\n, nested exception is: %v",cmd.class,err)
		return
	}
	fmt.Printf("class data :%v\n",classData)
}