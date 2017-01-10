package main

import "os"
import "io/ioutil"
import "sync"
import "fmt"
import "github.com/ishiikurisu/edf"

// Processes a folder containing many EDF files
func main() {
		var wg sync.WaitGroup
		var source string

		// Deciding source folder
		if len(os.Args) > 1 {
				source = os.Args[1]
		} else {
				source = "."
		}
		source += "/"

		// Processing each file on directory
		files, _ := ioutil.ReadDir(source)
		for _, file := range files {
				fileName := file.Name()
				if validFile(fileName) {
						wg.Add(1)
						go process(source + fileName, &wg)
				}
		}
		wg.Wait()
}

// Checks if the file name is validFile
func validFile(inlet string) bool {
		isEdf := true
		ext := ".edf"

		for i := 1; i <= 4 && isEdf; i++ {
				if inlet[len(inlet)-i] != ext[len(ext)-i] {
						 isEdf = false
				}
		}

		return isEdf
}

// Processes an EDF file, extracting the subject name and the sampling rate
func process(inlet string, wg *sync.WaitGroup) {
		edfHeader, _ := edf.ReadFile(inlet)
		fmt.Printf("%s\t%s\t%s\n", inlet, edfHeader["patient"], edf.GetSampling(edfHeader))
		wg.Done()
}
