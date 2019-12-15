package main

type IWebService interface {
	LogError(message string)
}

type logAnalyzer struct {
	service IWebService
}

func NewLogAnalyzer(service IWebService) *logAnalyzer {
	return &logAnalyzer{service: service}
}

func (l *logAnalyzer) Analize(filename string) {
	if len(filename) < 8 {
		// Logs error in production code
		l.service.LogError("Filename too short: " + filename)
	}
}
