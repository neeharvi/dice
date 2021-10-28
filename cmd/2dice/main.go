package main

import (
	"flag"
	"fmt"

	"github.com/neeharvi/dice"
)

func main() {
	var verbose bool
	flag.BoolVar(&verbose, "v", verbose, "")
	flag.BoolVar(&verbose, "verbose", verbose, "show separate dice rolls")
	flag.Parse()

	n1 := dice.Roll()
	n2 := dice.Roll()

	if verbose {
		fmt.Println(n1, "+", n2, "=", n1+n2)
	} else {
		fmt.Println(n1 + n2)
	}
}
