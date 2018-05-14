package runtime

import "jvmgo_redo/ch08/runtime/heap"

//表示局部变量表，局部变量可以为数字类型，因此会有一个属性是空的（0）
type Slot struct {
	num int32
	ref *heap.Object
}

