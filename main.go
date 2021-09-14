package main

import (
	"fmt"
	"github.com/dullgiulio/pingo"
	"log"
	"math"
	"unsafe"
)

const E = math.Pi

func main() {
	//con() // 无类型常量的使用，可以赋值给不同类型的变量
	//float() // float32精确度为6位，float64为14位
	//stringAndChar() // diff of string and char

}

func callPlu() {
	// Make a new plugin from the executable we created. Connect to it via TCP
	p := pingo.NewPlugin("tcp", "pingo/hello")
	// Actually start the plugin
	p.Start()
	// Remember to stop the plugin when done using it
	defer p.Stop()

	var resp string

	// Call a function from the object we created previously
	if err := p.Call("MyPlugin.SayHello", "Go developer", &resp); err != nil {
		log.Print(err)
	} else {
		log.Print(resp)
	}
}

func stringAndChar() {
	ch := 'A'
	str := "A"
	b := byte('A')
	by := []byte("A")
	fmt.Println(unsafe.Sizeof(ch))     //4  int32
	fmt.Println(unsafe.Sizeof(str))    //16 复合结构体指针
	fmt.Println(unsafe.Sizeof(str[0])) //1  str[0]只读，不能写
	fmt.Println(unsafe.Sizeof(b))      //1  uint8
	fmt.Println(unsafe.Sizeof(by))     //24
}

func float() {
	var a float32 = 1.00000001
	var b float32 = 1.00000002
	fmt.Println(a == b)
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(float64(a)))
}

func con() {
	var a float32 = E
	var b float64 = E
	fmt.Println(a)
	fmt.Println(b)
}
