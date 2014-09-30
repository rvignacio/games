package permutations

import (
	"testing"
)

func Test2Permutation(t *testing.T) {

	r := Generate("A", "B")

	if len(r) != 2 {
		t.Errorf("Expected 2 permutations, but got %d.", len(r))
	}

	// Test r[0] == ["A", "B"]
	if r[0][0] != "A" && r[0][1] != "B" {
		t.Errorf("Expected first permutation to be [A, B], but got %+v", r[0])
	}

	// Test r[1] == ["B", "A"]
	if r[1][0] != "B" && r[1][1] != "A" {
		t.Errorf("Expected first permutation to be [B, A], but got %+v", r[0])
	}

}

func TestNPermutation(t *testing.T) {

	var fact func(i int) int

	fact = func(i int) int {
		if i == 1 {
			return 1
		}
		return i * fact(i-1)
	}

	values := make([]interface{}, 0)

	for _, v := range []interface{}{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"} {
		values = append(values, v)
		r := Generate(values...)
		if len(r) != fact(len(values)) {
			t.Errorf("%d", len(r))
		}
	}

}
