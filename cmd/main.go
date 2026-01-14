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

	// 1. Extract command and arguments
	var args []string
	command := os.Args[1]

	// Handle "./arc build ..." vs "./arc file.arc ..."
	if command == "build" {
		args = os.Args[2:]
	} else if command == "help" {
		printUsage()
		os.Exit(0)
	} else if strings.HasPrefix(command, "-") {
		// flags passed directly? Treat as build
		args = os.Args[1:]
	} else {
		// Assume it's an input file, implicit build
		args = os.Args[1:]
	}

	// 2. Parse Flags manually to support "arc build file.arc -o out" order
	config := compiler.Config{
		OutputType: compiler.OutputExecutable, // Default
	}

	var inputFile string
	
	for i := 0; i < len(args); i++ {
		arg := args[i]

		if strings.HasPrefix(arg, "-") {
			switch arg {
			case "-o":
				if i+1 < len(args) {
					config.OutputFile = args[i+1]
					i++
				}
			case "-emit":
				if i+1 < len(args) {
					val := args[i+1]
					i++
					switch val {
					case "exe": config.OutputType = compiler.OutputExecutable
					case "obj": config.OutputType = compiler.OutputObject
					case "ir":  config.OutputType = compiler.OutputIR
					}
				}
			case "-v":
				config.Verbose = true
			case "-L":
				if i+1 < len(args) {
					config.LibraryPaths = append(config.LibraryPaths, args[i+1])
					i++
				}
			case "-l":
				if i+1 < len(args) {
					config.Libraries = append(config.Libraries, args[i+1])
					i++
				}
			default:
				// Handle joined flags like -lc or -L/usr/lib
				if strings.HasPrefix(arg, "-l") {
					config.Libraries = append(config.Libraries, arg[2:])
				} else if strings.HasPrefix(arg, "-L") {
					config.LibraryPaths = append(config.LibraryPaths, arg[2:])
				} else {
					fmt.Fprintf(os.Stderr, "Unknown flag: %s\n", arg)
				}
			}
		} else {
			if inputFile == "" {
				inputFile = arg
			} else {
				// Ignore extra args for now
			}
		}
	}

	if inputFile == "" {
		fmt.Fprintln(os.Stderr, "Error: No input file specified")
		printUsage()
		os.Exit(1)
	}

	config.InputFile = inputFile

	// 3. Run Compiler
	comp := compiler.NewCompiler()
	if err := comp.Run(config); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: arc build <file> [options]")
	fmt.Println("Options:")
	fmt.Println("  -o <file>   Output file")
	fmt.Println("  -L <path>   Add library search path")
	fmt.Println("  -l <lib>    Link library (e.g. -l c)")
	fmt.Println("  -emit <fmt> Output format (exe, obj, ir)")
	fmt.Println("  -v          Verbose")
}