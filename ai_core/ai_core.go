/**
 * AI核心代码, 估值1个亿
 */
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	for {
		line, err := r.ReadString('\n')
		if err == nil {
			line = strings.Replace(line, "吗", "", -1)
			line = strings.Replace(line, "? ", "!", -1)
			line = strings.Replace(line, "?", "!", -1)
			fmt.Println(line)
		}
	}
}
