package main

import "os"
import "fmt"
import "github.com/ishiikurisu/edf"

func main() {
    input := os.Args[1]
    edfContents := edf.ReadFile(input)
    fmt.Println(edfContents.WriteCSV())
}
