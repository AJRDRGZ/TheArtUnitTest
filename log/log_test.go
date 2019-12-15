package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Testing a struct by calling a method and checking the value of a property
func TestAnalizer_IsValidLogFileName_WhenCalled_ChangesWasLastFileNameValid(t *testing.T) {
	tests := []struct {
		fileName string
		expected bool
	}{
		{"badfile.foo", false},
		{"goodfile.slf", true},
	}
	for _, tt := range tests {
		la := makeAnalizer()
		la.IsValidLogFileName(tt.fileName)

		assert.Equal(t, tt.expected, la.WasLastFileNameValid,
			"WasLastFileNameValid was not updated")
	}
}

func makeAnalizer() *Analizer {
	return &Analizer{}
}
