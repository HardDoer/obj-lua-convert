package main

import (
	"os"
	"log"
)

func main() {
	someArgs := os.Args[1:]
	if (len(someArgs) == 0) {
		log.Fatalf("No arguments supplied to program. Please provide a path to the .obj file eg. \"/home/me/3dmodel.obj\"")
		os.Exit(1)
	}
	log.Print(someArgs)
}