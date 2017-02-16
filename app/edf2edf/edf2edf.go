package main

import "fmt"
import "os"
import "github.com/ishiikurisu/edf"

func main() {
	input := os.Args[1]
	output := os.Args[2]

	fmt.Println("---")
	fmt.Println("# EDF to EDF #")
	fmt.Println("Input: " + input)
	fmt.Println("Output: " + output)

	// Reading inlet
	inlet := edf.ReadFile(input)
	// Updating data
	inlet.Header["patient"] = "Johnny"
	// Writting to outlet
	inlet.WriteEdf(output)

	fmt.Println("...")
}
