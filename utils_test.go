package mrz

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClearText(t *testing.T) {
	assert.Equal(t, " 1233", clearText("<1233<<"))
}

func TestGetText(t *testing.T) {
	data := []string{"abc", "xyz"}
	pos := Position{
		Line:  0,
		Start: 1,
		End:   3,
	}
	assert.Equal(t, "bc", getText(data, &pos))
}

func TestCheckDigitCode(t *testing.T) {
	assert.True(t, checkDigitCode("870314", "5"))
	assert.True(t, checkDigitCode("59000002", "8"))
}
