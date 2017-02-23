package main

import "os"
import "github.com/ishiikurisu/edf"

func main() {
	input := os.Args[1]
	output := os.Args[2]
	stuff := edf.ReadFile(input)
	stuff.WriteCsvToFile(output)
}
