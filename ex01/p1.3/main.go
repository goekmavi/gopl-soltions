// TODO: benchmark tests
package main

import (
	"fmt"
	"os"
	"strings"
	//"time"
)

func main() {
	//start := time.Now()
	withRange(os.Args[1:])
	//end := time.Since(start).Seconds()
	//fmt.Println("After withRange: ", end)

	//start = time.Now()
	withJoin(os.Args[1:])
	//end = time.Since(start).Seconds()
	//fmt.Println("After withJoin: ", end)
}

func withRange(args []string) {
	result := ""
	tmp := ""

	for _, arg := range args {
		result += tmp + arg
		tmp = " "
	}

	fmt.Println(result)
}

func withJoin(args []string) {
	fmt.Println(strings.Join(args, " "))
}
