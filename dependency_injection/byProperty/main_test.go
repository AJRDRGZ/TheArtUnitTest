package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type fakeExtensionManager struct {
	WillBeValid bool
}

func (this *fakeExtensionManager) IsValid(fileName string) bool {
	return this.WillBeValid
}

func Test_logAnalyzer_IsValidLogFileName_NameSupportedExtension_ReturnsTrue(t *testing.T) {
	// Sets up stub to return true
	myFakeManager := new(fakeExtensionManager)
	myFakeManager.WillBeValid = true

	log := NewLogAnalyzer()
	log.SetManager(myFakeManager) // Injects a stub

	got := log.IsValidLogFileName("short.ext")

	assert.True(t, got)
}
