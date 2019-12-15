package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testableLogAnalyzer struct {
	logAnalyzerUsingFactoryMethod
}

func newTestableLogAnalyzer(mgr IExtensionManager) *testableLogAnalyzer {
	t := &testableLogAnalyzer{}
	t.GetManager = func() IExtensionManager {
		return mgr
	}

	return t
}

func Test_logAnalyzer_IsValidLogFileName_NameSupportedExtension_ReturnsTrue(t *testing.T) {
	// Sets up stub to return true
	stub := new(fakeExtensionManager)
	stub.WillBeValid = true

	logan := newTestableLogAnalyzer(stub)
	got := logan.IsValidLogFileName("file.ext")

	assert.True(t, got)
}

type fakeExtensionManager struct {
	WillBeValid bool
}

func (this *fakeExtensionManager) IsValid(fileName string) bool {
	return this.WillBeValid
}
