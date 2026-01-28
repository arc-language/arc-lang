package pkg

import (
	"fmt"
	"github.com/go-git/go-git/v5"
)

type GitProvider struct {
	URL string
}

func (g *GitProvider) Download(destPath string) error {
	// Simple git clone
	_, err := git.PlainClone(destPath, false, &git.CloneOptions{
		URL:      g.URL,
		Progress: nil, // Add os.Stdout to see progress
		Depth:    1,   // Shallow clone is faster
	})
	
	if err != nil {
		return fmt.Errorf("git clone failed: %v", err)
	}
	return nil
}