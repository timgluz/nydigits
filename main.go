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

func run(target int, digits []int) (nydigits.Solution, error) {
	if target < 1 {
		return nydigits.Solution{}, fmt.Errorf("Target must be a positive integer")
	}

	if len(digits) == 0 {
		return nydigits.Solution{}, fmt.Errorf("No digits provided")
	}

	solution, err := nydigits.Solve(target, digits)
	if err != nil {
		return nydigits.Solution{}, err
	}

	return solution, nil
}

func showSolution(solution nydigits.Solution, target int, digits []int) {
	fmt.Println("----------------------------------")
	fmt.Println("NY-Digits Solver")
	fmt.Printf("Target:   %d\n", target)
	fmt.Printf("Digits:   %v\n", digits)
	fmt.Printf("Solution: %d\n", solution.Value)
	fmt.Println("----------------------------------")
	fmt.Printf("Operations:\n")
	for i, step := range solution.Operations {
		if i == 0 {
			continue // Skip the first step
		}

		fmt.Printf("\t%d: %s\n", i, step)
	}
}

func main() {
	var target = flag.Int("target", 0, "Target number")
	flag.Parse()

	digits := parseDigits(flag.Args())
	if len(digits) == 0 {
		fmt.Println("No digits provided")
		return
	}

	solution, err := run(*target, digits)
	if err != nil {
		fmt.Println("Error:", err)
	}

	showSolution(solution, *target, digits)
}
