package nydigits

import "fmt"

type Operator int

const (
	NoOp Operator = iota
	Plus
	Minus
	Times
	Divide
)

var ops = [4]Operator{Plus, Minus, Times, Divide}

func (op Operator) String() string {
	switch op {
	case Plus:
		return "+"
	case Minus:
		return "-"
	case Times:
		return "*"
	case Divide:
		return "/"
	default:
		return " "
	}
}

func (op Operator) Apply(a, b int) (int, error) {
	switch op {
	case Plus:
		return a + b, nil
	case Minus:
		if a < b {
			return 0, fmt.Errorf("Can't subtract %d from %d", b, a)
		}

		return a - b, nil
	case Times:
		return a * b, nil
	case Divide:
		if b == 0 {
			return 0, fmt.Errorf("Can't divide by zero")
		}

		if a%b != 0 {
			return 0, fmt.Errorf("Can't divide %d by %d", a, b)
		}

		return a / b, nil
	default:
		return 0, fmt.Errorf("Unknown operator")
	}
}

type OperationStep struct {
	Op        Operator
	Digit     int
	PrevValue int
	Value     int
}

func (s OperationStep) String() string {
	return fmt.Sprintf("%d %s %d = %d", s.PrevValue, s.Op, s.Digit, s.Value)
}

type Solution struct {
	Value      int
	Operations []OperationStep
}

type Node struct {
	Parent *Node
	Op     Operator
	Value  int
	Digit  int

	UnusedDigits []int
	Children     []*Node
}

func NewNode(value int, op Operator) *Node {
	return &Node{
		Parent:       nil,
		Op:           op,
		Value:        value,
		UnusedDigits: []int{},
	}
}

func (n *Node) AddChild(op Operator, digit int) error {
	if n.Digit == digit {
		return fmt.Errorf("Can't add child with same digit")
	}

	newValue, err := op.Apply(n.Value, digit)
	if err != nil {
		return err
	}

	if newValue < 1 {
		return fmt.Errorf("Only positive values are allowed")
	}

	newUnusedDigits := cloneWithout(n.UnusedDigits, digit)

	child := &Node{
		Parent:       n,
		Op:           op,
		Value:        newValue,
		Digit:        digit,
		UnusedDigits: newUnusedDigits,
	}

	n.Children = append(n.Children, child)
	return nil
}

func (n *Node) Steps() []OperationStep {
	steps := []OperationStep{}

	currentNode := n
	for currentNode.Parent != nil {
		step := OperationStep{
			Op:        currentNode.Op,
			Digit:     currentNode.Digit,
			PrevValue: currentNode.Parent.Value,
			Value:     currentNode.Value,
		}

		steps = append([]OperationStep{step}, steps...)

		currentNode = currentNode.Parent
	}

	return steps
}

// cloneWithout returns a new slice with the given value removed
func cloneWithout(slice []int, value int) []int {
	if len(slice) == 0 {
		return []int{}
	}

	newSlice := make([]int, 0, len(slice)-1)
	for _, v := range slice {
		if v != value {
			newSlice = append(newSlice, v)
		}
	}

	return newSlice
}
func Solve(target int, digits []int) (Solution, error) {
	if target < 1 {
		return Solution{}, fmt.Errorf("Target must be a positive integer")
	}

	if len(digits) == 0 {
		return Solution{}, fmt.Errorf("No digits given")
	}

	root := NewNode(0, NoOp)
	root.UnusedDigits = digits

	bestSolution := Solution{
		Value: 0,
	}

	frontier := []*Node{root}

	var solutionNode *Node
	finished := false
	for !finished && len(frontier) > 0 {
		currentNode := frontier[0]
		frontier = frontier[1:]

		// add childs to new node
		for _, op := range ops {
			for _, digit := range currentNode.UnusedDigits {
				if err := currentNode.AddChild(op, digit); err != nil {
					continue
				}
			}
		}

		// add new nodes to queue
		for _, child := range currentNode.Children {
			if child.Value == target {
				solutionNode = child
				finished = true
				break
			}

			frontier = append(frontier, child)
		}
	}

	if solutionNode != nil {
		bestSolution.Value = solutionNode.Value
		bestSolution.Operations = solutionNode.Steps()
	}

	return bestSolution, nil
}
