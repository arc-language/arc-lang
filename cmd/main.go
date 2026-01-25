package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/arc-language/arc-lang/compiler"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	args := os.Args[1:]
	if args[0] == "build" { args = args[1:] }
	if args[0] == "help" { printUsage(); os.Exit(0) }

	config := compiler.Config{OutputType: compiler.OutputExecutable}
	var inputFile string

	for i := 0; i < len(args); i++ {
		arg := args[i]
		if strings.HasPrefix(arg, "-") {
			switch arg {
			case "-o":
				if i+1 < len(args) { config.OutputFile = args[i+1]; i++ }
			case "-emit":
				if i+1 < len(args) {
					switch args[i+1] {
					case "exe": config.OutputType = compiler.OutputExecutable
					case "obj": config.OutputType = compiler.OutputObject
					case "ir":  config.OutputType = compiler.OutputIR
					}
					i++
				}
			case "-v": config.Verbose = true
			case "-L":
				if i+1 < len(args) { config.LibraryPaths = append(config.LibraryPaths, args[i+1]); i++ }
			case "-l":
				if i+1 < len(args) { config.Libraries = append(config.Libraries, args[i+1]); i++ }
			}
		} else {
			if inputFile == "" { inputFile = arg }
		}
	}

	if inputFile == "" { inputFile = "." }
	config.InputFile = inputFile

	comp := compiler.NewCompiler()
	if err := comp.Run(config); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: arc [build] <path> [options]")
	fmt.Println("Options:")
	fmt.Println("  -o <file>   Output file")
	fmt.Println("  -emit <fmt> Output format (exe, obj, ir)")
	fmt.Println("  -L <path>   Add library search path")
	fmt.Println("  -l <lib>    Link library")
}