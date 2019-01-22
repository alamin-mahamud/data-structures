package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/alamin-mahamud/go-data-structures/array/rotation"
)

func main() {
	d := flag.Int("d", 2, "rotate by this value")
	s := flag.Int("s", -1, "search this Term")
	flag.Parse()

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

	if *d > len(a) {
		log.Fatal("Rotation Point should not be greater than length")
		os.Exit(1)
	}

	if *s == -1 {
		log.Fatal("You must input the search Term")
		os.Exit(1)
	}

	// rotation.SimpleRotation([]int{1, 2, 3, 4, 5, 6}, 3)
	// rotation.RotateOneByOne(&a, *d)
	// rotation.JugglingRotation(&a, *d)
	rotation.ReversalAlgorithm(&a, *d)
	fmt.Println(a)
	fmt.Println(rotation.PivotBinarySearch(&a, *s))
}
