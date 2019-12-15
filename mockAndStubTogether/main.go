package main

type IWebService interface {
	LogError(message string) error
}

type IEmailService interface {
	Send(to, subject, body string)
}

type logAnalyzer struct {
	service IWebService
	email   IEmailService
}

func NewLogAnalyzer(service IWebService, email IEmailService) *logAnalyzer {
	return &logAnalyzer{service, email}
}

func (l *logAnalyzer) Analize(filename string) {
	if len(filename) < 8 {
		err := l.service.LogError("Filename too short: " + filename)
		if err != nil {
			l.email.Send("someone@somewhere.com", "can’t log", err.Error())
		}
	}
}
