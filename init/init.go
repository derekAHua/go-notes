// 同一个文件可以有多个init函数
// init函数
package main

import "fmt"

func main() {

}

func init() {
	fmt.Println("init1...")
}

func init() {
	fmt.Println("init2...")
}
