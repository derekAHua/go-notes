// 枚举
package main

import "fmt"

type ChipType uint8

const (
	None ChipType = iota
	Cpu
	Gpu
)

func (c ChipType) String() string {
	switch c {
	case None:
		return "none"
	case Cpu:
		return "CPU"
	case Gpu:
		return "GPU"
	default:
		return "DEFAULT"
	}
}

func main() {
	fmt.Println(Cpu.String())
}
