// EDLP - https://github.com/jh2119/edlp
//
package main

import (
	"fmt"
	"path/filepath"
)

const (
	ARCHIVE_EXTENSION = ".eat"
)

func generateFilenameOutput(filenameInput string, pathDestination string) string {
	// use the original file and path by default
	var filenameInputBase string = filenameInput

	// check for a change of path
	if len(pathDestination) != 0 {
		filenameInputBase = pathDestination + "/" + filepath.Base(filenameInput)
	}

	fileExt := filepath.Ext(filenameInput)
	fileNameLastIndex := len(filenameInputBase) - len(fileExt)
	return filenameInputBase[:fileNameLastIndex] + ARCHIVE_EXTENSION
}
