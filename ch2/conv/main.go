package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/tommyp/gopl.io/ch2/tempconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()

		s, _ := strconv.ParseFloat(text, 64)

		c := tempconv.CToF(tempconv.Celsius(s))

		fmt.Println(c.String())
	}
}
