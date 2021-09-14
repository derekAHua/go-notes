// go:generate 的使用
package main

import "fmt"

//go:generate go run main.go
func main() {
	fmt.Println("Generate")

	p := B
	fmt.Println(p.String())
}

//go:generate stringer -type=Pill
type Pill int

const (
	A Pill = iota
	B
	C
	D
)
