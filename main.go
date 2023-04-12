package main

import (
	"fmt"

	"github.com/timgluz/pkg/nydigits"
)

func main() {
	target := 62
	digits := []int{1, 2, 3, 4, 5, 10}
	solution, err := nydigits.Solve(target, digits)
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}

	fmt.Println("Solution: %d, %s", solution.Value, solution.Path)
}
