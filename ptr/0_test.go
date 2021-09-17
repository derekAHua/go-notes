package main

import (
	"fmt"
	"testing"
	"unsafe"
)

// 通过指针获取slice长度
func TestGetSliceLen(t *testing.T) {
	//type Slice struct {
	//	Data unsafe.Pointer
	//	Len  int
	//	Cap  int
	//}
	s := make([]int, 9, 20)
	var Len = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(8)))
	fmt.Println(Len, len(s)) // 9 9
	var Cap = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(16)))
	fmt.Println(Cap, cap(s)) // 20 20

}
