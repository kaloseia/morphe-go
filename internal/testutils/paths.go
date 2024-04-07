package testutils

import (
	"path/filepath"
	"runtime"
)

func GetTestDirPath() string {
	_, fileName, _, _ := runtime.Caller(1)
	rootPath, absErr := filepath.Abs(filepath.Join(filepath.Dir(fileName), "../../"))
	if absErr != nil {
		panic(absErr)
	}
	return filepath.Join(rootPath, "testdata")
}
