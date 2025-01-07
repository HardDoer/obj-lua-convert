package main

import (
	"os"
	"fmt"
)

func main() {
	someArgs := os.Args
	fmt.Print(someArgs)
}