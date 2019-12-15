package main

type IExtensionManager interface {
	IsValid(fileName string) (bool, error)
}

type logAnalyzer struct {
	manager IExtensionManager
}

func NewLogAnalyzer(mgr IExtensionManager) *logAnalyzer {
	return &logAnalyzer{manager: mgr}
}

func (l *logAnalyzer) IsValidLogFileName(fileName string) (bool, error) {
	return l.manager.IsValid(fileName)
}
