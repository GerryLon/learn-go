package main

import (
	"fmt"
	"github.com/GerryLon/learn-go/lang/gogenerate/painkiller"
)

//go:generate echo hello
//go:generate go run main.go
//go:generate  echo file=$GOFILE pkg=$GOPACKAGE
func main() {
	var pill painkiller.Pill
	pill = painkiller.Aspirin

	fmt.Printf("%s", pill)
}
