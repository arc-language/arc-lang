package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type TestResult struct {
	Name       string
	Passed     bool
	Error      string
	Duration   time.Duration
	ExitCode   int
	Expected   int
	Compiled   bool
	Linked     bool
	CompileLog string
	LinkLog    string
	RunLog     string
}

var logFile *os.File

func main() {
	// Open log file
	var err error
	logFile, err = os.Create("test.log")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating log file: %v\n", err)
		os.Exit(1)
	}
	defer logFile.Close()

	// First, make sure arc binary exists
	arcBinary := "./arc"
	if _, err := os.Stat(arcBinary); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Error: 'arc' binary not found. Please build it first:\n")
		fmt.Fprintf(os.Stderr, "  go build -o arc main.go\n")
		logToFile("Error: 'arc' binary not found\n")
		os.Exit(1)
	}

	testDir := "tests"
	
	// Find all .arc test files
	arcFiles, err := filepath.Glob(filepath.Join(testDir, "test_*.arc"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error finding test files: %v\n", err)
		logToFile(fmt.Sprintf("Error finding test files: %v\n", err))
		os.Exit(1)
	}

	if len(arcFiles) == 0 {
		fmt.Fprintf(os.Stderr, "No .arc test files found in %s/\n", testDir)
		logToFile(fmt.Sprintf("No .arc test files found in %s/\n", testDir))
		os.Exit(1)
	}

	header := "Arc Language Test Runner\n" + strings.Repeat("=", 70) + "\n"
	fmt.Print(header)
	logToFile(header)
	
	msg := fmt.Sprintf("Found %d test files\n\n", len(arcFiles))
	fmt.Print(msg)
	logToFile(msg)

	results := make([]TestResult, 0, len(arcFiles))
	tempDir := filepath.Join(os.TempDir(), "arc_tests")
	os.RemoveAll(tempDir) // Clean up from previous runs
	os.MkdirAll(tempDir, 0755)
	defer os.RemoveAll(tempDir)

	for i, arcFile := range arcFiles {
		testName := strings.TrimSuffix(filepath.Base(arcFile), ".arc")
		msg := fmt.Sprintf("[%d/%d] %-45s ", i+1, len(arcFiles), testName)
		fmt.Print(msg)
		logToFile(msg)
		
		start := time.Now()
		result := runTest(arcBinary, arcFile, testName, tempDir)
		result.Duration = time.Since(start)
		results = append(results, result)

		if result.Passed {
			msg := fmt.Sprintf("✅ PASS (%.2fs)\n", result.Duration.Seconds())
			fmt.Print(msg)
			logToFile(msg)
		} else {
			msg := fmt.Sprintf("❌ FAIL (%.2fs)\n", result.Duration.Seconds())
			fmt.Print(msg)
			logToFile(msg)
			
			if !result.Compiled {
				msg := "    Compilation failed\n"
				fmt.Print(msg)
				logToFile(msg)
				logToFile(fmt.Sprintf("    Compile output:\n%s\n", indent(result.CompileLog)))
			} else if !result.Linked {
				msg := "    Linking failed\n"
				fmt.Print(msg)
				logToFile(msg)
				logToFile(fmt.Sprintf("    Link output:\n%s\n", indent(result.LinkLog)))
			} else {
				msg := fmt.Sprintf("    Expected exit=%d, got exit=%d\n", result.Expected, result.ExitCode)
				fmt.Print(msg)
				logToFile(msg)
				if result.RunLog != "" {
					logToFile(fmt.Sprintf("    Run output:\n%s\n", indent(result.RunLog)))
				}
			}
			if result.Error != "" && len(result.Error) < 200 {
				msg := fmt.Sprintf("    %s\n", result.Error)
				fmt.Print(msg)
				logToFile(msg)
			}
		}
	}

	// Print summary
	fmt.Println()
	logToFile("\n")
	separator := strings.Repeat("=", 70) + "\n"
	fmt.Print(separator)
	logToFile(separator)
	printSummary(results)
	
	logToFile("\nLog file written to: test.log\n")
}

func runTest(arcBinary, arcFile, testName, tempDir string) TestResult {
	result := TestResult{
		Name:     testName,
		Expected: 0, // Default expected exit code
	}

	logToFile(fmt.Sprintf("\n--- Testing: %s ---\n", testName))
	logToFile(fmt.Sprintf("Source file: %s\n", arcFile))

	// Check for expected exit code in filename (e.g., test_name_exit42.arc)
	if strings.Contains(testName, "_exit") {
		parts := strings.Split(testName, "_exit")
		if len(parts) == 2 {
			fmt.Sscanf(parts[1], "%d", &result.Expected)
			logToFile(fmt.Sprintf("Expected exit code: %d\n", result.Expected))
		}
	}

	// Paths
	objFile := filepath.Join(tempDir, testName+".o")
	exeFile := filepath.Join(tempDir, testName)

	logToFile(fmt.Sprintf("Object file: %s\n", objFile))
	logToFile(fmt.Sprintf("Executable: %s\n", exeFile))

	// Step 1: Compile .arc to .o using arc binary
	logToFile("\nStep 1: Compiling to object file...\n")
	cmd := exec.Command(arcBinary, "build", arcFile, "-o", objFile)
	output, err := cmd.CombinedOutput()
	result.CompileLog = string(output)
	
	if err != nil {
		result.Error = strings.TrimSpace(string(output))
		logToFile(fmt.Sprintf("Compilation failed: %v\n", err))
		logToFile(fmt.Sprintf("Output:\n%s\n", output))
		return result
	}
	
	result.Compiled = true
	logToFile("Compilation successful\n")
	if len(output) > 0 {
		logToFile(fmt.Sprintf("Compile output:\n%s\n", output))
	}

	// Step 2: Link with gcc
	logToFile("\nStep 2: Linking with gcc...\n")
	cmd = exec.Command("gcc", objFile, "-o", exeFile, "-no-pie")
	output, err = cmd.CombinedOutput()
	result.LinkLog = string(output)
	
	if err != nil {
		result.Error = strings.TrimSpace(string(output))
		logToFile(fmt.Sprintf("Linking failed: %v\n", err))
		logToFile(fmt.Sprintf("Output:\n%s\n", output))
		return result
	}
	
	result.Linked = true
	logToFile("Linking successful\n")
	if len(output) > 0 {
		logToFile(fmt.Sprintf("Link output:\n%s\n", output))
	}

	// Step 3: Run the executable
	logToFile("\nStep 3: Running executable...\n")
	cmd = exec.Command(exeFile)
	output, err = cmd.CombinedOutput()
	result.RunLog = string(output)
	result.ExitCode = 0
	
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			result.ExitCode = exitErr.ExitCode()
			logToFile(fmt.Sprintf("Process exited with code: %d\n", result.ExitCode))
		} else {
			result.Error = fmt.Sprintf("Execution failed: %v", err)
			logToFile(fmt.Sprintf("Execution error: %v\n", err))
			return result
		}
	} else {
		logToFile("Process exited with code: 0\n")
	}
	
	if len(output) > 0 {
		logToFile(fmt.Sprintf("Program output:\n%s\n", output))
	}

	// Step 4: Check exit code matches expected
	if result.ExitCode == result.Expected {
		result.Passed = true
		logToFile("✅ Test PASSED\n")
	} else {
		logToFile(fmt.Sprintf("❌ Test FAILED: expected exit=%d, got exit=%d\n", result.Expected, result.ExitCode))
		if len(output) > 0 {
			result.Error = fmt.Sprintf("Output: %s", strings.TrimSpace(string(output)))
		}
	}

	return result
}

func printSummary(results []TestResult) {
	passed := 0
	failed := 0
	compileErrors := 0
	linkErrors := 0
	runtimeErrors := 0
	totalDuration := time.Duration(0)

	for _, r := range results {
		totalDuration += r.Duration
		if r.Passed {
			passed++
		} else {
			failed++
			if !r.Compiled {
				compileErrors++
			} else if !r.Linked {
				linkErrors++
			} else {
				runtimeErrors++
			}
		}
	}

	summary := fmt.Sprintf("Total Tests:      %d\n", len(results))
	summary += fmt.Sprintf("Passed:           %d (%.1f%%)\n", passed, float64(passed)/float64(len(results))*100)
	summary += fmt.Sprintf("Failed:           %d\n", failed)
	if failed > 0 {
		summary += fmt.Sprintf("  Compile errors: %d\n", compileErrors)
		summary += fmt.Sprintf("  Link errors:    %d\n", linkErrors)
		summary += fmt.Sprintf("  Runtime errors: %d\n", runtimeErrors)
	}
	summary += fmt.Sprintf("Total Duration:   %.2fs\n", totalDuration.Seconds())

	fmt.Print(summary)
	logToFile(summary)

	if failed > 0 {
		failList := "\nFailed Tests:\n"
		for _, r := range results {
			if !r.Passed {
				status := "runtime"
				if !r.Compiled {
					status = "compile"
				} else if !r.Linked {
					status = "link"
				}
				failList += fmt.Sprintf("  ❌ %-40s (%s)\n", r.Name, status)
			}
		}
		failList += "\n"
		fmt.Print(failList)
		logToFile(failList)
		os.Exit(1)
	} else {
		success := "\n🎉 All tests passed!\n"
		fmt.Print(success)
		logToFile(success)
	}
}

func logToFile(msg string) {
	if logFile != nil {
		logFile.WriteString(msg)
	}
}

func indent(s string) string {
	if s == "" {
		return ""
	}
	lines := strings.Split(strings.TrimRight(s, "\n"), "\n")
	for i, line := range lines {
		lines[i] = "      " + line
	}
	return strings.Join(lines, "\n")
}