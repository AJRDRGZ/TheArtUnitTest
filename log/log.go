package log

import (
	"errors"
	"strings"
)

var (
	ErrFileNameHasNotProvided = errors.New("filename has to be provided")
)

type Analizer struct {
	WasLastFileNameValid bool
}

func (this *Analizer) IsValidLogFileName(fileName string) bool {
	this.WasLastFileNameValid = false

	if !isValidExtension(fileName) {
		return false
	}

	this.WasLastFileNameValid = true
	return true
}

func isValidExtension(fileName string) bool {
	return strings.HasSuffix(strings.ToUpper(fileName), ".SLF")
}
