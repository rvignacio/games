package life

var RPentomino = &Game{
	b: NewBoard(3, 3,
		Point{0, 1},
		Point{0, 2},
		Point{1, 0},
		Point{1, 1},
		Point{2, 1},
	),
	lastMovements: make([]Movement, 3),
}
