// pkg/manager.go
package pkg

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/arc-language/upkg"
	"github.com/go-git/go-git/v5"
)

// PackageManager handles resolving and downloading dependencies
type PackageManager struct {
	Config *upkg.Config
}

// NewPackageManager creates a manager using upkg's default configuration
func NewPackageManager() *PackageManager {
	// This defaults to ~/.upkg (or ~/.cache/upkg depending on implementation)
	// We want to ensure we use the install path for stability
	cfg := upkg.DefaultConfig()
	
	// Ensure root dir exists
	os.MkdirAll(cfg.InstallPath, 0755)

	return &PackageManager{
		Config: cfg,
	}
}

// Ensure checks if a package exists, and downloads it if missing.
// lang: "c", "cpp" (implies system lib), or "" (implies arc module)
// importPath: "sqlite", "openssl", "github.com/user/repo"
// Returns the absolute path to the downloaded package root.
func (pm *PackageManager) Ensure(lang string, importPath string) (string, error) {
	// Clean the path
	importPath = strings.Trim(importPath, "\"")

	// Strategy 1: Arc Source Modules (Git)
	// Identified by having a domain structure (dot in the first segment) or specific protocol
	if lang == "" && (strings.Contains(importPath, ".") || strings.Contains(importPath, "/")) {
		return pm.ensureGitModule(importPath)
	}

	// Strategy 2: System Libraries (upkg)
	// Used for 'import c "sqlite"' or 'import cpp "openssl"'
	return pm.ensureSystemPackage(importPath)
}

// ensureGitModule handles "github.com/user/repo"
func (pm *PackageManager) ensureGitModule(importPath string) (string, error) {
	// Store source code in ~/.upkg/src/github.com/user/repo
	targetDir := filepath.Join(pm.Config.InstallPath, "src", importPath)

	if _, err := os.Stat(targetDir); err == nil {
		return targetDir, nil
	}

	fmt.Printf("[Pkg] Cloning %s...\n", importPath)
	
	// Assuming HTTPS. In a real scenario, might handle SSH or other protocols.
	url := "https://" + importPath
	if !strings.HasPrefix(importPath, "http") {
		url = "https://" + importPath
	}

	_, err := git.PlainClone(targetDir, false, &git.CloneOptions{
		URL:      url,
		Depth:    1,
		Progress: os.Stdout,
	})
	if err != nil {
		os.RemoveAll(targetDir) // Clean up partials
		return "", fmt.Errorf("git clone failed for %s: %w", importPath, err)
	}

	return targetDir, nil
}

// ensureSystemPackage handles "sqlite", "curl" via upkg
func (pm *PackageManager) ensureSystemPackage(name string) (string, error) {
	// Initialize upkg with Auto backend (detects apt, brew, winget, etc.)
	mgr, err := upkg.NewManager(upkg.BackendAuto, pm.Config)
	if err != nil {
		return "", fmt.Errorf("failed to initialize upkg: %w", err)
	}
	defer mgr.Close()

	// Check if already installed? 
	// For now, we rely on upkg.Download's idempotency or just try to download.
	// In the future, upkg.IsInstalled(name) would be better.

	ctx := context.Background()
	fmt.Printf("[Pkg] Resolving system dependency: %s\n", name)

	// Download/Install
	// Auto backend resolves "sqlite" -> "libsqlite3-dev" (apt) or "sqlite" (brew)
	pkg := &upkg.Package{Name: name}
	opts := &upkg.DownloadOptions{
		Extract:    boolPtr(true),
		VerifyHash: boolPtr(true),
	}

	if err := mgr.Download(ctx, pkg, opts); err != nil {
		return "", fmt.Errorf("upkg failed to install %s: %w", name, err)
	}

	// Return the install path root. 
	// Compiler/Linker will look for /lib and /include inside here or in system paths.
	return pm.Config.InstallPath, nil
}

func boolPtr(b bool) *bool {
	return &b
}