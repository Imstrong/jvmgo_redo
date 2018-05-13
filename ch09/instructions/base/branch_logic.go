package base

import "jvmgo/ch08/runtime"

func Branch(frame *runtime.Frame,offset int){
	pc:=frame.Thread().PC()
	nextPC:=pc+offset
	frame.SetNextPC(nextPC)
}