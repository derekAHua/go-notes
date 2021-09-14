// os 获取输入参数
package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	for _, arg := range os.Args { // 第一个参数为程序的执行路径
		fmt.Println(arg)
	}

	fmt.Println("------")

	cmd := exec.Command(os.Args[1], os.Args[2])
	out, _ := cmd.CombinedOutput()
	fmt.Println(string(out))

}

// ahua@ahuadembp os % go run os.go go version
// /var/folders/l2/x2rvxx5s66s8msgczgl4762c0000gn/T/go-build1410208944/b001/exe/os
// go
// version
// ------
// go version go1.16.7 darwin/arm64
