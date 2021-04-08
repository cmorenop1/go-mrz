package mrz

type Position struct {
	Line  uint `json:"line"`
	Start uint `json:"start"`
	End   uint `json:"end"`
}

type Range struct {
	Position *Position `json:"position"`
	Raw      string    `json:"raw"`
}

type Field struct {
	Name     string    `json:"name"`
	Value    string    `json:"value"`
	Valid    bool      `json:"valid"`
	Ranges   []Range   `json:"ranges"`
	Position *Position `json:"position"`
}

type Document struct {
	Format  string                 `json:"format"`
	Fields  []*Field               `json:"details"`
	Summary map[string]interface{} `json:"fields"`
	Valid   bool                   `json:"valid"`
}

type fieldParser func(lines []string) (*Field, error)
type textParser func(s string) (string, error)
