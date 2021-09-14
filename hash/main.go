package main

import (
	"crypto/md5"
	"crypto/sha512"
	"fmt"
)

func main() {
	s1 := "jknjkabfjkdbjkf"
	s2 := "jknjkabfjkdbjkf"

	m := md5.New()
	m.Write([]byte(s1))
	m.Write([]byte(s2))

	fmt.Printf("%x\n", m.Sum(nil))

	s := sha512.New()
	s.Write([]byte(s1))
	fmt.Printf("%x", s.Sum(nil))

}
