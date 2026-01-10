package compiler

import "strings"

type OutputType int

const (
	OutputUnknown OutputType = iota
	OutputExecutable
	OutputObject
	OutputIR
)

// Config holds all parameters required for a compilation session.
type Config struct {
	InputFile   string
	OutputFile  string
	OutputType  OutputType
	
	// Linker Options
	LinkShared   bool     // Create .so instead of executable (future use)
	LibraryPaths []string // -L flags
	Libraries    []string // -l flags (e.g. "c", "m")
	
	// Debugging
	Verbose     bool
}

// PostProcess infers default filenames if the user didn't specify -o
func (c *Config) PostProcess() {
	if c.OutputFile == "" {
		// Strip extension from input: main.arc -> main
		base := c.InputFile
		if idx := strings.LastIndex(base, "."); idx != -1 {
			base = base[:idx]
		}
		
		switch c.OutputType {
		case OutputObject:
			c.OutputFile = base + ".o"
		case OutputIR:
			c.OutputFile = base + ".ir"
		default:
			// Executable has no extension on Linux
			c.OutputFile = base 
		}
	}
}