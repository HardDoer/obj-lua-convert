package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	someArgs := os.Args[1:]
	if (len(someArgs) == 0) {
		log.Fatalf("No arguments supplied to program. Please provide a path to the .obj file eg. \"/home/me/3dmodel.obj\"")
		os.Exit(1)
	}
	var fileName string 
	var vertices = "vertices={\n"
	var faces = "faces={\n"
	var result = ""
	file, err := os.Open(someArgs[0])
	if (err != nil) {
		log.Fatalf("Error reading file.")
		os.Exit(2)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		var split = strings.Split(fileScanner.Text(), " ")
		var prefix = split[0]
		if prefix == "mtllib" {
			fileName = strings.Split(split[1], ".")[0] + ".lua"
		} else if prefix == "v" {
			v_row := "{x=" + split[1] + "," + "y=" + split[2] + ",z=" + split[3] + ",w=1},\n" 
			vertices += v_row
		} else if prefix == "f" {
			f_row := "{"
			size := len(split)
			for i := 1; i < size; i ++ {
				f_row += strings.Split(split[i], "/")[0]
				if (i + 1 < size) {
					f_row += ","
				}
			}
			faces += f_row + "},\n"
		}
	}
	file.Close()
	result += "{\n" + vertices + "},\n"
	result += faces + "}\n}"
    os.WriteFile("./" + fileName, []byte(result), 0644)
}