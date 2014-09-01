package life

type Board [][]Cell

func (b Board) Rows() int {
	return len(b)
}

func (b Board) Columns() int {
	return len(b[0])
}

func (b Board) SumNeighbours(x, y int) (n int) {

	for row := -1; row <= 1; row++ {

		for column := -1; column <= 1; column++ {

			if row == 0 && column == 0 {
				continue
			}

			nx := x + row
			ny := y + column
			switch {
			case nx < 0 || nx >= b.Rows():
			case ny < 0 || ny >= b.Columns():
				break
			default:
				cell := b[nx][ny]
				if cell {
					n += 1
				}
			}

		}

	}

	return
}

func (b Board) String() (out string) {

	rows := b.Rows()
	columns := b.Columns()

	for x := 0; x < rows; x++ {
		for y := 0; y < columns; y++ {
			out += b[x][y].String()
		}
		out += "\n"
	}

	return
}

// Create a new board.
func NewBoard(rows, columns int, startingCells ...Point) (b Board) {

	b = Board(make([][]Cell, rows))
	for i := range b {
		b[i] = make([]Cell, columns)
	}

	for _, p := range startingCells {
		b[p.X][p.Y].Live()
	}

	return
}
