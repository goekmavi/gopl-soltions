// Print out the command-line arguments and the name of the command that invoked it.
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0:])
}
