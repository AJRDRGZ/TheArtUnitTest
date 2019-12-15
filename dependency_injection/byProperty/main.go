package main

type IExtensionManager interface {
	IsValid(fileName string) bool
}

type logAnalyzer struct {
	manager IExtensionManager
}

func NewLogAnalyzer() *logAnalyzer {
	return &logAnalyzer{
		manager: &FileExtensionManager{}, // default dependency
	}
}

func (l *logAnalyzer) SetManager(manager IExtensionManager) {
	l.manager = manager
}

func (l *logAnalyzer) GetManager() IExtensionManager {
	return l.manager
}

func (l *logAnalyzer) IsValidLogFileName(fileName string) bool {
	return l.manager.IsValid(fileName)
}

type FileExtensionManager struct{}

func (f *FileExtensionManager) IsValid(fileName string) bool {
	// Here logic
	return false
}
