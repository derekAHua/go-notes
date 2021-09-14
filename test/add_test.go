// go test 使用
package test

import (
	"testing"
)

func TestAdd(t *testing.T) {
	r := Add(1, 1)
	if r != 2 {
		t.Error("Add Err..")
	}
}
func TestHelloWorld(t *testing.T) {
	t.Log("HelloWorld")
}
