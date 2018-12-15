package main

import (
	"fmt"
	"github.com/GerryLon/learn-go/lang/ui/messagebox"
)

func main() {
	yes := messagebox.MessageBox4windows("Are You Sure?", "Press Yes Or No")
	if yes {
		fmt.Println("you selected yes")
	}
}
