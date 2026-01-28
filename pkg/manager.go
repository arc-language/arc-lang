package pkg

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// PackageManager handles resolving and downloading dependencies
type PackageManager struct {
	RootPath string // usually ~/.arc/
}

// NewPackageManager creates a manager pointing to ~/.arc/
func NewPackageManager() *PackageManager {
	home, err := os.UserHomeDir()
	if err != nil {
		home = "/tmp"
	}
	root := filepath.Join(home, ".arc")
	return &PackageManager{
		RootPath: root,
	}
}

// Ensure checks if a package exists, and downloads it if missing.
// lang: "c", "cpp", "objc", or "" (empty for normal arc imports)
// importPath: "github.com/user/repo", "nix.org/lib", "brew.sh/lib"
// Returns the absolute path to the downloaded package/library.
func (pm *PackageManager) Ensure(lang string, importPath string) (string, error) {
	// Clean the path (remove quotes if passed raw)
	importPath = strings.Trim(importPath, "\"")

	var provider PackageProvider
	var targetDir string

	// 1. Detect Provider based on string prefix
	if strings.HasPrefix(importPath, "nix.org/") {
		pkgName := strings.TrimPrefix(importPath, "nix.org/")
		targetDir = filepath.Join(pm.RootPath, "nix", pkgName)
		provider = &NixProvider{Name: pkgName}
	} else if strings.HasPrefix(importPath, "brew.sh/") {
		pkgName := strings.TrimPrefix(importPath, "brew.sh/")
		targetDir = filepath.Join(pm.RootPath, "brew", pkgName)
		provider = &BrewProvider{Formula: pkgName}
	} else if strings.Contains(importPath, "github.com/") || strings.Contains(importPath, "gitlab.com/") {
		// Standard Arc Package
		targetDir = filepath.Join(pm.RootPath, "src", importPath)
		provider = &GitProvider{URL: "https://" + importPath}
	} else {
		// Local import or system library (e.g., "sqlite3" without prefix)
		// We return empty string to let the generic importer handle it as a system path
		return "", nil
	}

	// 2. Check if already exists
	if _, err := os.Stat(targetDir); err == nil {
		// Already exists, return path
		return targetDir, nil
	}

	// 3. Create directory
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create package dir: %v", err)
	}

	fmt.Printf("[Pkg] Downloading %s to %s...\n", importPath, targetDir)

	// 4. Download
	if err := provider.Download(targetDir); err != nil {
		// Cleanup on failure
		os.RemoveAll(targetDir)
		return "", fmt.Errorf("failed to download %s: %v", importPath, err)
	}

	return targetDir, nil
}

// PackageProvider interface for different download strategies
type PackageProvider interface {
	Download(destPath string) error
}