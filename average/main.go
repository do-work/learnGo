// readfile calculates the average of several numbers from a text file.
package main

import (
	"fmt"
	"github.com/do-work/learnGo/datafile"
	"log"
)

func main() {
	results, err := datafile.GetFloats("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	var sum float64 = 0
	for _, number := range results {
		sum += number
	}
	sampleCount := float64(len(results))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)
}
