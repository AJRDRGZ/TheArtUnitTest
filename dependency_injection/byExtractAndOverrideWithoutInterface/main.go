package main

type logAnalyzerUsingFactoryMethod struct {
	isValid func(filename string) bool
}

func NewLogAnalyzerUsingFactoryMethod() *logAnalyzerUsingFactoryMethod {
	return &logAnalyzerUsingFactoryMethod{isValid: isValidExtManager}
}

func (l *logAnalyzerUsingFactoryMethod) IsValidLogFileName(fileName string) bool {
	return l.isValid(fileName)
}

func isValidExtManager(filename string) bool {
	// Returns result from real dependency
	mgr := new(FileExtensionManager)
	return mgr.IsValid(filename)
}

type FileExtensionManager struct{}

func (f *FileExtensionManager) IsValid(fileName string) bool {
	// Here logic
	return false
}
