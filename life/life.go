package life

type Status bool

const (
	Alive Status = true
	Dead         = false
)

type Cell struct {
	Status Status
}

type MacroCell struct {
	NE, SE, SW, NW, Result MacroCell
}
