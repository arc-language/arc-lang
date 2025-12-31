package compiler

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Importer resolves paths and finds source files
type Importer struct {
	// Cache could go here in the future
}

func NewImporter() *Importer {
	return &Importer{}
}

// ResolveImport converts a relative import string to an absolute path
func (imp *Importer) ResolveImport(currentDir, importPath string) (string, error) {
	// 1. Handle local relative imports (./ or ../)
	if strings.HasPrefix(importPath, ".") {
		abs, err := filepath.Abs(filepath.Join(currentDir, importPath))
		if err != nil {
			return "", err
		}
		return abs, nil
	}

	// 2. Handle standard library or module imports (Future)
	// For now, treat everything as relative
	return filepath.Abs(filepath.Join(currentDir, importPath))
}

// GetSourceFiles returns all .arc files in a directory
func (imp *Importer) GetSourceFiles(dirPath string) ([]string, error) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read dir: %v", err)
	}

	var files []string
	for _, entry := range entries {
		if !entry.IsDir() && (strings.HasSuffix(entry.Name(), ".arc")) {
			files = append(files, filepath.Join(dirPath, entry.Name()))
		}
	}
	return files, nil
}