package main

import (
	"flag"
	"fmt"

	nydigits "github.com/timgluz/nydigits/pkg"
)

func parseDigits(args []string) []int {
	digits := make([]int, len(args))
	for i, arg := range args {
		fmt.Sscanf(arg, "%d", &digits[i])
	}
	return digits
}

func main() {
	var target = flag.Int("target", 0, "Target number")
	flag.Parse()

	if *target < 1 {
		fmt.Println("Target must be a positive integer")
		return
	}

	digits := parseDigits(flag.Args())
	if len(digits) == 0 {
		fmt.Println("No digits provided")
		return
	}

	solution, err := nydigits.Solve(*target, digits)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Solution:  %s\n", solution.Path)
}
