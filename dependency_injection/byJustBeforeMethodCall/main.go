package main

var (
	// ExtensionManagerFactory global variable for simulate static structure and allow monkey patch
	ExtensionManagerFactory = new(extensionManagerFactory)
)

type IExtensionManager interface {
	IsValid(fileName string) bool
}

type logAnalyzer struct {
	manager IExtensionManager
}

func NewLogAnalyzer() *logAnalyzer {
	return &logAnalyzer{
		manager: ExtensionManagerFactory.Create(), // Uses factory in production code
	}
}

func (l *logAnalyzer) IsValidLogFileName(fileName string) bool {
	return l.manager.IsValid(fileName)
}

// FileExtensionManager default dependency object
type FileExtensionManager struct{}

func (f *FileExtensionManager) IsValid(fileName string) bool {
	// Here logic
	return false
}

type extensionManagerFactory struct {
	customManager IExtensionManager
}

// Create defines factory that can use and return custom manager
func (e *extensionManagerFactory) Create() IExtensionManager {
	if e.customManager != nil {
		return e.customManager
	}

	return new(FileExtensionManager)
}

func (e *extensionManagerFactory) SetManager(mgr IExtensionManager) {
	e.customManager = mgr
}
