package life

type Status int

const (
	Alive Status = 1
	Dead         = 0
)

// ul ur
// bl br
type Cell struct {
	Status                 Status
	ul, ur, bl, br, Result *Cell
	Level                  int
}

func (c Cell) IsAlive() bool {
	return c.Status == Alive
}

func (c *Cell) nextGen() *Cell {
	r := &Cell{}

	switch {
	// Level 0: 1 by 1
	// x
	// Level 1: 2 by 2
	// x x
	// x x
	case c.Level == 0, c.Level == 1:
		return nil
	// Level 2: 4 by 4
	// Return center as result (1 generation forward)
	case c.Level == 2:
		return l2Result(c)
	case c.Level > 2:

	}

	return r
}

func NewCell(ul, ur, bl, br *Cell, level int) *Cell {

	c := &Cell{
		ul:    ul,
		ur:    ul,
		bl:    bl,
		br:    br,
		Level: level,
	}

	c.Result = c.nextGen()

	return c

}

func NextStatus(ul, uc, ur, cl, cc, cr, bl, bc, br *Cell) Status {
	n := ul.Status + uc.Status + ur.Status +
		cl.Status + cr.Status +
		bl.Status + bc.Status + br.Status

	switch {
	case n == 3:
	case n == 2 && cc.IsAlive():
		return Alive
	}

	return Dead
}

// Level 2: 4 by 4
// x x x x
// x o o x
// x o o x
// x x x x
// Return center as result (1 generation forward)
func l2Result(c *Cell) (r *Cell) {
	r.Level = 2

	// Result upper left
	// x x x
	// x ? o
	// x o o
	r.ul.Status = NextStatus(
		c.ul.ul, c.ul.ur, c.ur.ul,
		c.ul.bl, c.ul.br, c.ur.bl,
		c.bl.ul, c.bl.ur, c.br.ul,
	)

	// Result upper right
	// x x x
	// o ? x
	// o o x
	r.ur.Status = NextStatus(
		c.ul.ur, c.ur.ul, c.ur.ur,
		c.ul.br, c.ur.bl, c.ur.br,
		c.bl.ur, c.br.bl, c.br.br,
	)

	// Result bottom left
	// x o o
	// x ? o
	// x x z
	r.bl.Status = NextStatus(
		c.ul.bl, c.ul.br, c.ur.bl,
		c.bl.ul, c.bl.ur, c.br.ul,
		c.bl.bl, c.bl.br, c.br.bl,
	)

	// Result bottom right
	// o o x
	// o ? x
	// x x x
	r.br.Status = NextStatus(
		c.ul.br, c.ur.bl, c.ur.br,
		c.bl.ur, c.br.ul, c.br.ur,
		c.bl.br, c.br.bl, c.br.br,
	)

	return
}
