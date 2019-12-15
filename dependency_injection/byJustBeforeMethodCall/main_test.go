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
	defer func(original *extensionManagerFactory) {
		// restore after use (monkey patch)
		ExtensionManagerFactory = original
	}(ExtensionManagerFactory)

	// Sets up stub to return true
	myFakeManager := new(fakeExtensionManager)
	myFakeManager.WillBeValid = true

	// Sets stub into factory struct for this test
	ExtensionManagerFactory.SetManager(myFakeManager)

	log := NewLogAnalyzer()
	got := log.IsValidLogFileName("short.ext")

	assert.True(t, got)
}
