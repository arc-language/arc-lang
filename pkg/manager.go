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
	cfg := upkg.DefaultConfig()
	os.MkdirAll(cfg.InstallPath, 0755)

	return &PackageManager{
		Config: cfg,
	}
}

// Ensure checks if a package exists, and downloads it if missing.
// lang: "c", "cpp" (system/wrapper), or "" (arc module)
// importPath: "sqlite", "github.com/user/repo", "github.com/user/repo/folder"
// Returns the absolute path to the downloaded package root (or subdir).
func (pm *PackageManager) Ensure(lang string, importPath string) (string, error) {
	importPath = strings.Trim(importPath, "\"")

	// Detect if this is a Remote/Git import (contains domain or slash)
	// Works for both Arc modules and C wrappers (import c "github.com/...")
	isRemote := strings.Contains(importPath, ".") || strings.Contains(importPath, "/")

	if isRemote {
		return pm.ensureGitModule(importPath)
	}

	// Strategy 2: System Libraries (upkg)
	// Used for 'import c "sqlite"' or 'import c "libc"'
	return pm.ensureSystemPackage(importPath)
}

// ensureGitModule handles "github.com/user/repo" and "github.com/user/repo/subdir"
func (pm *PackageManager) ensureGitModule(importPath string) (string, error) {
	// Heuristic: Assume first 3 components are the repo (host/user/repo)
	// e.g. github.com/johndoe/wrapper/lib -> repo: github.com/johndoe/wrapper
	parts := strings.Split(importPath, "/")
	
	var repoPath, subDir string

	if len(parts) >= 3 {
		repoPath = strings.Join(parts[:3], "/")
		if len(parts) > 3 {
			subDir = filepath.Join(parts[3:]...)
		}
	} else {
		// Fallback for things like "localhost/repo"
		repoPath = importPath
	}

	// Store source code in ~/.upkg/src/github.com/user/repo
	targetRepoDir := filepath.Join(pm.Config.InstallPath, "src", repoPath)

	// Check if repo exists
	if _, err := os.Stat(targetRepoDir); os.IsNotExist(err) {
		fmt.Printf("[Pkg] Cloning %s...\n", repoPath)
		
		url := "https://" + repoPath
		if !strings.HasPrefix(repoPath, "http") {
			url = "https://" + repoPath
		}

		_, err := git.PlainClone(targetRepoDir, false, &git.CloneOptions{
			URL:      url,
			Depth:    1,
			Progress: os.Stdout,
		})
		if err != nil {
			os.RemoveAll(targetRepoDir) // Clean up partials
			return "", fmt.Errorf("git clone failed for %s: %w", repoPath, err)
		}
	}

	// Return the specific subdirectory inside the cloned repo
	// If subDir is empty, this just returns the repo root
	fullPath := filepath.Join(targetRepoDir, subDir)
	return fullPath, nil
}

// ensureSystemPackage handles "sqlite", "curl" via upkg
func (pm *PackageManager) ensureSystemPackage(name string) (string, error) {
	// Initialize upkg with Auto backend
	mgr, err := upkg.NewManager(upkg.BackendAuto, pm.Config)
	if err != nil {
		return "", fmt.Errorf("failed to initialize upkg: %w", err)
	}
	defer mgr.Close()

	ctx := context.Background()
	fmt.Printf("[Pkg] Resolving system dependency: %s\n", name)

	// Download/Install
	pkg := &upkg.Package{Name: name}
	opts := &upkg.DownloadOptions{
		Extract:    boolPtr(true),
		VerifyHash: boolPtr(true),
	}

	if err := mgr.Download(ctx, pkg, opts); err != nil {
		return "", fmt.Errorf("upkg failed to install %s: %w", name, err)
	}

	return pm.Config.InstallPath, nil
}

func boolPtr(b bool) *bool {
	return &b
}