package life

import (
	"fmt"
	"testing"
)

func newLeaf(s Status) *Cell {
	return &Cell{
		Status: s,
	}
}

func swap(r []*Cell, i, j int) {
	tmp := r[i]
	r[i] = r[j]
	r[j] = tmp
}

func variants(alive, dead int) (variants [][]*Cell, e error) {
	if r := alive + dead; r != 9 {
		e = fmt.Errorf("Wrong number of cells %d", r)
		return
	}

	result := make([]*Cell, 9)

	for i := 0; i < 9; i++ {
		if i < alive {
			result[i] = newLeaf(Alive)
		} else {
			result[i] = newLeaf(Dead)
		}
	}

	variants = append(variants, result)

	counter := make([]int, 9)

	for i := 0; i < 9; {
		if counter[i] < i {
			if i%2 == 0 {
				swap(result, 0, i)
			} else {
				swap(result, counter[i], i)
			}

			counter[i] += 1

			i = 0

			variants = append(variants, result)
		} else {
			counter[i] = 0
			i++
		}
	}

	return
}

func TestNextStatus_0Alive_9Dead(t *testing.T) {
	var s Status

	alive0_dead8 := []*Cell{
		newLeaf(Dead), newLeaf(Dead), newLeaf(Dead),
		newLeaf(Dead), newLeaf(Dead), newLeaf(Dead),
		newLeaf(Dead), newLeaf(Dead), newLeaf(Dead),
	}

	s = NextStatus(
		alive0_dead8[0], alive0_dead8[1], alive0_dead8[2],
		alive0_dead8[3], alive0_dead8[4], alive0_dead8[5],
		alive0_dead8[6], alive0_dead8[7], alive0_dead8[8],
	)

	if s != Dead {
		t.Errorf("Expected 'Dead', but got %v.", s)
	}
}

func TestNextStatus_1Alive_8Dead(t *testing.T) {
	var s Status

	variants := [][]Status{
		[]Status{Alive, Dead, Dead, Dead, Dead, Dead, Dead, Dead, Dead},
		[]Status{Dead, Alive, Dead, Dead, Dead, Dead, Dead, Dead, Dead},
		[]Status{Dead, Dead, Alive, Dead, Dead, Dead, Dead, Dead, Dead},
		[]Status{Dead, Dead, Dead, Alive, Dead, Dead, Dead, Dead, Dead},
		[]Status{Dead, Dead, Dead, Dead, Alive, Dead, Dead, Dead, Dead},
		[]Status{Dead, Dead, Dead, Dead, Dead, Alive, Dead, Dead, Dead},
		[]Status{Dead, Dead, Dead, Dead, Dead, Dead, Alive, Dead, Dead},
		[]Status{Dead, Dead, Dead, Dead, Dead, Dead, Dead, Alive, Dead},
		[]Status{Dead, Dead, Dead, Dead, Dead, Dead, Dead, Dead, Alive},
	}

	for _, variant := range variants {
		alive1_dead7 := []*Cell{
			newLeaf(variant[0]), newLeaf(variant[1]), newLeaf(variant[2]),
			newLeaf(variant[3]), newLeaf(variant[4]), newLeaf(variant[5]),
			newLeaf(variant[6]), newLeaf(variant[7]), newLeaf(variant[8]),
		}

		s = NextStatus(
			alive1_dead7[0], alive1_dead7[1], alive1_dead7[2],
			alive1_dead7[3], alive1_dead7[4], alive1_dead7[5],
			alive1_dead7[6], alive1_dead7[7], alive1_dead7[8],
		)

		if s != Dead {
			t.Errorf("Expected 'Dead', but got %v.", s)
		}
	}
}

func TestNextStatus_2Alive_7Dead(t *testing.T) {
	var s Status

	variants := [][]Status{
		[]Status{Alive, Dead, Dead, Dead, Dead, Dead, Dead, Dead, Dead},
		[]Status{Dead, Alive, Dead, Dead, Dead, Dead, Dead, Dead, Dead},
		[]Status{Dead, Dead, Alive, Dead, Dead, Dead, Dead, Dead, Dead},
		[]Status{Dead, Dead, Dead, Alive, Dead, Dead, Dead, Dead, Dead},
		[]Status{Dead, Dead, Dead, Dead, Alive, Dead, Dead, Dead, Dead},
		[]Status{Dead, Dead, Dead, Dead, Dead, Alive, Dead, Dead, Dead},
		[]Status{Dead, Dead, Dead, Dead, Dead, Dead, Alive, Dead, Dead},
		[]Status{Dead, Dead, Dead, Dead, Dead, Dead, Dead, Alive, Dead},
		[]Status{Dead, Dead, Dead, Dead, Dead, Dead, Dead, Dead, Alive},
	}

	for _, variant := range variants {
		alive1_dead7 := []*Cell{
			newLeaf(variant[0]), newLeaf(variant[1]), newLeaf(variant[2]),
			newLeaf(variant[3]), newLeaf(variant[4]), newLeaf(variant[5]),
			newLeaf(variant[6]), newLeaf(variant[7]), newLeaf(variant[8]),
		}

		s = NextStatus(
			alive1_dead7[0], alive1_dead7[1], alive1_dead7[2],
			alive1_dead7[3], alive1_dead7[4], alive1_dead7[5],
			alive1_dead7[6], alive1_dead7[7], alive1_dead7[8],
		)

		if s != Dead {
			t.Errorf("Expected 'Dead', but got %v.", s)
		}
	}
}
