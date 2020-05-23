package main

import (
	"fmt"
)

type A interface {
	FA()
}

type B interface {
	FB()
}

type C interface {
	A
	B
	FC()
}

type STB struct {
}

func (s *STB) FB() {
	fmt.Println("fba")
}

func main() {
	var st = &ST{}
}
