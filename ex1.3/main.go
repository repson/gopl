// ex1.3 Experiment to measure the difference in running
// time between our potentially inefficient versions and
// the one that uses strings.Join. (Section 1.6
// illustrates part of the time package, and Section
// 11.4 shows how to write benchmark tests for systematic
// performance evaluation.)

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func inefficient() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func efficient() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func main() {
	start1 := time.Now()
	inefficient()
	t1 := time.Now()
	elapsed1 := t1.Sub(start1)
	println(elapsed1.String())

	println()

	start2 := time.Now()
	efficient()
	t2 := time.Now()
	elapsed2 := t2.Sub(start2)
	println(elapsed2.String())
}
