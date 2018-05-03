package base

import "jvmgo/ch07/runtime"

func Branch(frame *runtime.Frame,offset int){
	pc:=frame.Thread().PC()
	nextPC:=pc+offset
	frame.SetNextPC(nextPC)
}