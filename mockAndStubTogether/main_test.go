package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type fakeWebService struct {
	ToError error
}

func (f *fakeWebService) LogError(message string) error {
	if f.ToError != nil {
		return f.ToError
	}

	return nil
}

type fakeEmailService struct {
	to, subject, body string
}

func (f *fakeEmailService) Send(to, subject, body string) {
	f.to = to
	f.subject = subject
	f.body = body
}

func Test_logAnalyzer_Analyze_WebServiceThrows_SendsEmail(t *testing.T) {
	stubService := new(fakeWebService)
	stubService.ToError = errors.New("fake exception")

	mockEmail := new(fakeEmailService)

	log := NewLogAnalyzer(stubService, mockEmail)
	tooShortFileName := "abc.ext"
	log.Analize(tooShortFileName)

	assert.Equal(t, "someone@somewhere.com", mockEmail.to)
	assert.Equal(t, "fake exception", mockEmail.body)
	assert.Equal(t, "can’t log", mockEmail.subject)
}
