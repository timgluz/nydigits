package nydigits

import (
	"reflect"
	"testing"
)

func TestOperatorString(t *testing.T) {
	var tests = []struct {
		op   Operator
		want string
	}{
		{NoOp, " "},
		{Plus, "+"},
		{Minus, "-"},
		{Times, "*"},
		{Divide, "/"},
	}

	for _, test := range tests {
		if got := test.op.String(); got != test.want {
			t.Errorf("Operator.String() = %q, want %q", got, test.want)
		}
	}
}

func TestOperatorApply(t *testing.T) {
	var tests = []struct {
		op   Operator
		a    int
		b    int
		want int
		err  bool
	}{
		{NoOp, 0, 0, 0, true},
		{Plus, 1, 1, 2, false},
		{Minus, 1, 1, 0, false},
		{Times, 2, 2, 4, false},
		{Divide, 4, 2, 2, false},
		{Divide, 4, 3, 0, true},
	}

	for _, test := range tests {
		got, err := test.op.Apply(test.a, test.b)
		if err != nil && !test.err {
			t.Errorf("Operator.Apply() = %q, want %q", err, test.want)
		}

		if got != test.want {
			t.Errorf("Operator.Apply() = %q, want %q", got, test.want)
		}
	}
}

func TestRootNodeAddChild(t *testing.T) {
	testNode := NewNode(0, NoOp)

	var tests = []struct {
		title    string
		operator Operator
		value    int
		err      bool
	}{
		{"NoOp is not allowed for child nodes", NoOp, 0, true},
		{"Child and Parent values are same", Plus, 0, true},
		{"Child has null value", Plus, 0, true},
		{"Child has negative value", Plus, -1, true},
		{"Child and Parent values are different", Plus, 1, false},
	}

	for _, test := range tests {
		err := testNode.AddChild(test.operator, test.value)
		if test.err == false && err != nil {
			t.Errorf("%s: expected no error when adding (%s, %d), but got %v", test.title, test.operator, test.value, err)
		}

		if test.err == true && err == nil {
			t.Errorf("%s: expected error when adding (%s, %d), but got nil", test.title, test.operator, test.value)
		}
	}
}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func operationsEqual(a, b []OperationStep) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if !reflect.DeepEqual(a[i], b[i]) {
			return false
		}
	}

	return true
}

func TestCloneWithout(t *testing.T) {
	var tests = map[string]struct {
		slice []int
		value int
		want  []int
	}{
		"Empty slice": {
			slice: []int{},
			value: 0,
			want:  []int{},
		},
		"Slice with one element": {
			slice: []int{1},
			value: 1,
			want:  []int{},
		},
		"Slice with two elements": {
			slice: []int{1, 2},
			value: 1,
			want:  []int{2},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := cloneWithout(test.slice, test.value)
			if !slicesEqual(got, test.want) {
				t.Errorf("cloneWithout(%v, %d) = %v, want %v", test.slice, test.value, got, test.want)
			}
		})
	}
}

func TestSolve(t *testing.T) {
	var tests = []struct {
		target   int
		digits   []int
		expected Solution
		err      bool
	}{
		{0, []int{0}, Solution{}, true},
		{1, []int{}, Solution{}, true},
		{1, []int{1}, Solution{Operations: []OperationStep{{Op: Plus, Digit: 1, PrevValue: 0, Value: 1}}}, false},
	}

	for _, test := range tests {
		got, err := Solve(test.target, test.digits)
		if test.err == false && err != nil {
			t.Errorf("Solve(%d, %v) = %v, want %v", test.target, test.digits, err, test.expected)
		}

		if test.err == true && err == nil {
			t.Errorf("Solve(%d, %v) = %v, want %v", test.target, test.digits, err, test.expected)
		}

		if !operationsEqual(got.Operations, test.expected.Operations) {
			t.Errorf(
				"Solve(%d, %v) = %v, want %v",
				test.target, test.digits, got.Operations, test.expected.Operations,
			)
		}
	}
}
