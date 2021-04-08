package mrz

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var sampleTD1 = []string{
	"I<SWE59000002<8198703142391<<<",
	"8703145M1701027SWE<<<<<<<<<<<8",
	"SPECIMEN<<SVEN<<<<<<<<<<<<<<<<",
}

func TestParse(t *testing.T) {
	data := strings.Join(sampleTD1, "\n")
	doc, _ := Parse(data)
	assert.Equal(t, "TD1", doc.Format)
}

func TestParseTD1(t *testing.T) {
	doc, _ := parseTD1(sampleTD1)
	for _, f := range doc.Fields {
		fmt.Println(f.Name, f.Value, f.Valid)
	}
	fmt.Println(doc.Valid)
	assert.Equal(t, "TD1", doc.Format)
}
