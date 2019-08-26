// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Echo3 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

//!+
func main() {
	start := time.Now()

	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	fmt.Printf("%5.fs elapsed\n", time.Since(start).Nanoseconds())

	s, sep = "", ""
	for i, arg := range os.Args[1:] {
		fmt.Println(strconv.Itoa(i) + " " + arg)
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	fmt.Printf("%5.fs elapsed\n", time.Since(start).Nanoseconds())

	fmt.Println(strings.Join(os.Args[1:], " "))

	fmt.Printf("%5.fs elapsed\n", time.Since(start).Nanoseconds())
}

//!-
