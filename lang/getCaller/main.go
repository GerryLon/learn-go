package main

import (
	"fmt"
	"runtime"
)

func main() {
	funcName, file, line, ok := runtime.Caller(0)
	if ok {
		fmt.Println("func name: " + runtime.FuncForPC(funcName).Name())
		fmt.Printf("file: %s, line: %d\n", file, line)
	}
}
