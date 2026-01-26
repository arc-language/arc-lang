package compiler

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Importer resolves paths and finds source files
type Importer struct{}

func NewImporter() *Importer {
	return &Importer{}
}

// ResolveImport converts an import string to an absolute directory path.
// baseFile is the absolute path of the file containing the import statement.
func (imp *Importer) ResolveImport(baseFile, importPath string) (string, error) {
	currentDir := filepath.Dir(baseFile)

	// 1. Handle local relative imports (starts with . or ..)
	// Example: import "../socket" or import "./io"
	if strings.HasPrefix(importPath, ".") {
		abs, err := filepath.Abs(filepath.Join(currentDir, importPath))
		if err != nil {
			return "", err
		}
		return abs, nil
	}

	// 2. Default: Treat non-relative imports as relative to current directory for now.
	// In a full implementation, you would check GOROOT/ARCROOT or module paths here.
	abs, err := filepath.Abs(filepath.Join(currentDir, importPath))
	if err != nil {
		return "", err
	}
	return abs, nil
}

// GetSourceFiles returns all .ax files in a specific directory
func (imp *Importer) GetSourceFiles(dirPath string) ([]string, error) {
	info, err := os.Stat(dirPath)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("import path '%s' is not a directory", dirPath)
	}

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read dir: %v", err)
	}

	var files []string
	for _, entry := range entries {
		// FIXED: Check for .ax extension instead of .arc
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".ax") {
			files = append(files, filepath.Join(dirPath, entry.Name()))
		}
	}
	return files, nil
}