package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	ErrFake = errors.New("this is fake")
)

type fakeExtensionManager struct {
	WillBeValid bool
	WillError   error
}

func (this *fakeExtensionManager) IsValid(fileName string) (bool, error) {
	if this.WillError != nil {
		return false, this.WillError
	}

	return this.WillBeValid, nil
}

func Test_logAnalyzer_IsValidLogFileName_NameSupportedExtension_ReturnsTrue(t *testing.T) {
	// Sets up stub to return true
	myFakeManager := new(fakeExtensionManager)
	myFakeManager.WillBeValid = true

	log := NewLogAnalyzer(myFakeManager) // Sends in stub
	got, _ := log.IsValidLogFileName("short.ext")

	assert.True(t, got)
}

func Test_logAnalyzer_IsValidLogFileName_ExtManagerError_ReturnsFalse(t *testing.T) {
	myFakeManager := new(fakeExtensionManager)
	myFakeManager.WillError = ErrFake

	log := NewLogAnalyzer(myFakeManager)
	got, gotErr := log.IsValidLogFileName("anything.anyextension")

	assert.False(t, got)
	assert.EqualError(t, gotErr, ErrFake.Error())
}
