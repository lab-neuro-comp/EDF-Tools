package main

import (
	"os"
	"fmt"
	"github.com/ishiikurisu/edf"
)

func main() {
	edf1 := edf.ReadFile(os.Args[1])
	edf2 := edf.ReadFile(os.Args[2])
	result, oops := edf.Append(edf1, edf2)
	if oops != nil {
		fmt.Printf("%s\n", oops)
	} else {
		fmt.Printf("Writting %s...", os.Args[3])
		result.WriteEdf(os.Args[3])
		fmt.Println(" ok")
	}
}
