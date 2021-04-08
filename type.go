package mrz

type Position struct {
	Line  uint
	Start uint
	End   uint
}

type Range struct {
	Position *Position
	Raw      string
}

type Field struct {
	Name     string
	Value    string
	Valid    bool
	Ranges   []Range
	Position *Position
}

type Document struct {
	Format string
	Fields []*Field
	Valid  bool
}

type fieldParser func(lines []string) (*Field, error)
type textParser func(s string) (string, error)
