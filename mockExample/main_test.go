package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type fakeWebService struct {
	LastError string
}

func (f *fakeWebService) LogError(message string) {
	f.LastError = message
}

func Test_logAnalyzer_Analyze_TooShortFileName_CallsWebService(t *testing.T) {
	mockService := new(fakeWebService)
	log := NewLogAnalyzer(mockService)
	tooShortFileName := "abc.ext"

	log.Analize(tooShortFileName)

	assert.Equal(t, "Filename too short: abc.ext", mockService.LastError)
}
