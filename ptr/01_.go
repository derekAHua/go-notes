package main

import (
	"fmt"
	"unsafe"
)

// About unsafe.Pointer

//1.任何类型的指针都可以被转化为 unsafe.Pointer
//2.unsafe.Pointer 可以被转化为任何类型的指针
//3.uintptr 可以被转化为 unsafe.Pointer
//4.unsafe.Pointer 可以被转化为 uintptr

func mainV1() {
	num := 10
	p := &num
	fmt.Printf("p-->%v	*p-->%v\n", p, *p)

	var f *float64 = (*float64)(unsafe.Pointer(p))
	fmt.Printf("f-->%v	*f-->%v\n", f, *f)
}

func main() {
	arr := [...]int{1, 2, 3}
	fmt.Println(arr)
	fmt.Println(unsafe.Sizeof(arr[0]))

	ap := &arr
	sp := uintptr(unsafe.Pointer(ap))
	now := unsafe.Pointer(sp + unsafe.Sizeof(arr[0]))

	arr2 := (*[2]int)(now)
	fmt.Println(*arr2)
	a2 := (*int)(now)
	fmt.Println(*a2)

	fmt.Printf("%p\n", &arr)
	fmt.Printf("%p\n", arr2)
	fmt.Printf("%p\n", a2)

	fmt.Println("--------")
	*a2 += 8
	fmt.Println(arr)
	*arr2 = [2]int{100, 100}
	arr2[1] = 11
	fmt.Println(arr)

	s := (*[20]int)(now)
	fmt.Println(s)
}
