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
		return ""
	}
}

type Solution struct {
	Value  int
	Target int
	Path   string
}

type Node struct {
	Parent *Node
	Op     Operator
	Value  int
	Digit  int

	Children []*Node
}

func (n *Node) AddChild(op Operator, digit int) bool {
	if n.Digit == digit {
		return false
	}

	switch op {
	case Plus:
		return n.AddPlusChild(digit)
	case Minus:
		return n.AddMinusChild(digit)
	case Times:
		return n.AddTimesChild(digit)
	case Divide:
		return n.AddDivideChild(digit)
	default:
		return false
	}
}

func (n *Node) AddPlusChild(digit int) bool {
	child := &Node{
		Parent: n,
		Op:     Plus,
		Value:  n.Value + digit,
		Digit:  digit,
	}

	n.Children = append(n.Children, child)
	return true
}

func (n *Node) AddMinusChild(digit int) bool {
	if n.Value <= digit {
		return false
	}

	child := &Node{
		Parent: n,
		Op:     Minus,
		Value:  n.Value - digit,
		Digit:  digit,
	}

	n.Children = append(n.Children, child)
	return true
}

func (n *Node) AddTimesChild(digit int) bool {
	if n.Value == 0 || digit == 0 {
		return false
	}

	child := &Node{
		Parent: n,
		Op:     Times,
		Value:  n.Value * digit,
		Digit:  digit,
	}

	n.Children = append(n.Children, child)
	return true
}

func (n *Node) AddDivideChild(digit int) bool {
	if n.Value == 0 || digit == 0 {
		return false
	}

	if n.Value < digit {
		return false
	}

	if n.Value%digit != 0 {
		return false
	}

	child := &Node{
		Parent: n,
		Op:     Divide,
		Value:  n.Value / digit,
		Digit:  digit,
	}

	n.Children = append(n.Children, child)
	return true
}

func Solve(target int, digits []int) (Solution, error) {
	fmt.Println("Solving NYDigits")

	root := &Node{
		Parent: nil,
		Op:     NoOp,
		Value:  0,
	}

	ops := []Operator{Plus, Minus, Times, Divide}
	bestSolution := Solution{
		Value:  0,
		Target: target,
		Path:   "",
	}

	queue := []*Node{root}

	var solutionNode *Node
	finished := false
	for !finished && len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]

		// add childs to new node
		for _, op := range ops {
			for _, digit := range digits {
				currentNode.AddChild(op, digit)
			}
		}

		// add new nodes to queue
		for _, child := range currentNode.Children {
			if child.Value == target {
				fmt.Println("Found solution: ", child.Value)
				solutionNode = child
				finished = true
				break
			}

			queue = append(queue, child)
		}
	}

	if solutionNode != nil {
		bestSolution.Value = solutionNode.Value
		path := ""
		for node := solutionNode; node != nil; node = node.Parent {
			path = fmt.Sprintf("%s %d %s", node.Op, node.Digit, path)
		}

		bestSolution.Path = path
	}

	return bestSolution, nil
}
