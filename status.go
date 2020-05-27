package main

import (
	"fmt"
	"sort"
)

func status() {
	grades := map[string]float64{"Alma": 0, "Rohit": 86.5}
	var names []string

	for name := range grades {
		names = append(names, name)
	}

	sort.Strings(names)

	for _, name := range names {
		fmt.Printf("%s has a grade of %0.1f%%\n", name, grades[name])
	}
}

func main() {
	status()
}
