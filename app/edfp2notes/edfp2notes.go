package main

import "os"
import "fmt"
import "github.com/ishiikurisu/edf"

func main() {
	edfContents := edf.ReadFile(os.Args[1])
	notes := edfContents.WriteNotes()
	fmt.Printf("%v", notes)
}
