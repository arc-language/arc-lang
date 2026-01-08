package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/arc-language/arc-lang/codegen/codegen"
	"github.com/arc-language/arc-lang/compiler"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "build":
		handleBuild(os.Args[2:])
	case "help":
		printUsage()
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n\n", command)
		printUsage()
		os.Exit(1)
	}
}

func handleBuild(args []string) {
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "Error: No input file specified\n\n")
		printUsage()
		os.Exit(1)
	}

	inputFile := args[0]
	outputFile := ""

	// Parse -o flag
	for i := 1; i < len(args); i++ {
		if args[i] == "-o" && i+1 < len(args) {
			outputFile = args[i+1]
			break
		}
	}

	// Default output filename if not specified
	if outputFile == "" {
		base := filepath.Base(inputFile)
		ext := filepath.Ext(base)
		outputFile = strings.TrimSuffix(base, ext)
	}

	// Check if input file exists
	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Error: File '%s' does not exist\n", inputFile)
		os.Exit(1)
	}

	// Create the compiler instance
	comp := compiler.NewCompiler()

	// Run the project compilation pipeline.
	// CompileProject handles finding imports, parsing all files, 
	// running semantic analysis on the whole set, and generating a single IR module.
	module, err := comp.CompileProject(inputFile)
	if err != nil {
		// Error logging is handled inside the compiler (printing to stderr)
		// We just exit with a failure code here.
		os.Exit(1)
	}

	fmt.Printf("✓ Module '%s' compiled (%d functions)\n", module.Name, len(module.Functions))

	// Determine output format based on extension
	ext := strings.ToLower(filepath.Ext(outputFile))

	switch ext {
	case ".ir":
		// Option 1: Output Textual IR
		err := os.WriteFile(outputFile, []byte(module.String()), 0644)
		if err != nil {
			die("Failed to write IR file: %v", err)
		}
		fmt.Printf("✓ IR written to %s\n", outputFile)

	case ".o":
		// Option 2: Output Object File
		objData, err := codegen.GenerateObject(module)
		if err != nil {
			die("Object generation failed: %v", err)
		}
		
		err = os.WriteFile(outputFile, objData, 0644)
		if err != nil {
			die("Failed to write object file: %v", err)
		}
		fmt.Printf("✓ Object file written to %s\n", outputFile)
		printLinkHint(outputFile)

	default:
		// Option 3: Output Executable (Default)
		exeData, err := codegen.GenerateExecutable(module)
		if err != nil {
			die("Executable generation failed: %v", err)
		}

		// Write with executable permissions (0755)
		err = os.WriteFile(outputFile, exeData, 0755)
		if err != nil {
			die("Failed to write executable: %v", err)
		}
		fmt.Printf("✓ Executable written to %s\n", outputFile)
	}
}

func printUsage() {
	fmt.Println("Arc Language Compiler")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  arc build <source-file> [-o <output-file>]")
	fmt.Println("  arc help")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  arc build main.arc                    # Compiles to './main' executable")
	fmt.Println("  arc build main.arc -o output.ir       # Compiles to IR text")
}

func die(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Error: "+format+"\n", args...)
	os.Exit(1)
}

func printLinkHint(objFile string) {
	exeName := strings.TrimSuffix(filepath.Base(objFile), filepath.Ext(objFile))
	fmt.Printf("\nTo link manually:\n")
	fmt.Printf("  gcc %s -o %s -no-pie\n", objFile, exeName)
}