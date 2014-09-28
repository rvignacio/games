package permutations

import (
	"testing"
)

func TestHeapPermutation(t *testing.T) {

	r := Generate("A", "B", "C", "D")

	t.Errorf("%d\n%+v", len(r), r)

}
