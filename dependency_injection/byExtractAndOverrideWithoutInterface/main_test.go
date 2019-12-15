package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testableLogAnalyzer struct {
	isSupport bool
	logAnalyzerUsingFactoryMethod
}

func newTestableLogAnalyzer() *testableLogAnalyzer {
	t := &testableLogAnalyzer{}

	// override logAnalyzerUsingFactoryMethod.IsValid()
	t.isValid = func(filename string) bool {
		return t.isSupport // Returns fake value that was set by the test
	}

	return t
}

func Test_logAnalyzer_IsValidLogFileName_NameSupportedExtension_ReturnsTrue(t *testing.T) {
	logan := newTestableLogAnalyzer()
	logan.isSupport = true // Sets fake result value

	got := logan.IsValidLogFileName("file.ext")

	assert.True(t, got)
}
