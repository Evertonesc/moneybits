//go:build integration

package dirs

import (
	"os"
	"path/filepath"
	"runtime"
)

func BuildFixturePath(fixturePath string) string {
	wdir, _ := os.Getwd()
	resolvedPath := filepath.Join(wdir, fixturePath)

	return resolvedPath
}

func RootProjectDir() string {
	_, filename, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(filename)

	rootDir := filepath.Join(currentDir, "..", "..", "..", "..")

	return rootDir
}
