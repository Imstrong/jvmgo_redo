package main

import (
	"fmt"
	"strings"
	"jvmgo/ch07/classpath"
	"jvmgo/ch07/runtime/heap"
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
	classLoader := heap.NewClassLoader(cp,cmd.verboseClassFlag)
	className := strings.Replace(cmd.class, ".", "/", -1)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		interpret(mainMethod,cmd.verboseInstFlag)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}
}
