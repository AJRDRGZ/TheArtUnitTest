package main

type IExtensionManager interface {
	IsValid(fileName string) bool
}

type logAnalyzerUsingFactoryMethod struct {
	GetManager func() IExtensionManager
}

func NewLogAnalyzerUsingFactoryMethod() *logAnalyzerUsingFactoryMethod {
	l := &logAnalyzerUsingFactoryMethod{}
	l.GetManager = func() IExtensionManager {
		return new(FileExtensionManager)
	}

	return l
}

func (l *logAnalyzerUsingFactoryMethod) IsValidLogFileName(fileName string) bool {
	return l.GetManager().IsValid(fileName)
}

type FileExtensionManager struct{}

func (f *FileExtensionManager) IsValid(fileName string) bool {
	// Here logic
	return false
}
