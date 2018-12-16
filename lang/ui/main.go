package main

import (
	"github.com/GerryLon/learn-go/lang/ui/messagebox"
)

func main() {
	yes := messagebox.MessageBox4windows("真心话大冒险", "我帅吗?")
	for !yes {
		yes = messagebox.MessageBox4windows("真心话大冒险", "我帅吗?")
	}
	messagebox.MessageBox4windows("恭喜你!", "答对了.")
}
