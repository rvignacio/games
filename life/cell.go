package life

type Cell bool

func (c *Cell) Live() {
	*c = Cell(true)
}

func (c *Cell) Die() {
	*c = Cell(false)
}

func (c Cell) NextState(n int) Cell {

	if c {

		if n != 2 && n != 3 {
			c.Die()
		}

	} else {

		if n == 3 {
			c.Live()
		}

	}

	return c
}

func (c Cell) String() string {
	if c {
		return "*"
	} else {
		return " "
	}
}
