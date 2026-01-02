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

	// 1. Check for compiler binary
	arcBinary := "./arc"
	if _, err := os.Stat(arcBinary); os.IsNotExist(err) {
		fail("Error: 'arc' binary not found. Run: go build -o arc main.go")
	}

	// 2. Setup Test Directory
	testDir := filepath.Join("tests", "foundation")
	if _, err := os.Stat(testDir); os.IsNotExist(err) {
		os.MkdirAll(testDir, 0755)
		fmt.Printf("Created test directory: %s\n", testDir)
		fmt.Println("Please populate it with .arc files.")
		return
	}

	// 3. Find Test Files
	arcFiles, err := filepath.Glob(filepath.Join(testDir, "*.arc"))
	if err != nil {
		fail(fmt.Sprintf("Error searching for files: %v", err))
	}

	if len(arcFiles) == 0 {
		fail(fmt.Sprintf("No .arc files found in %s", testDir))
	}

	// 4. Header
	printHeader(len(arcFiles))

	// 5. Run Tests
	results := make([]TestResult, 0, len(arcFiles))
	tempDir := filepath.Join(os.TempDir(), "arc_tests")
	os.RemoveAll(tempDir)
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

		printResult(result)
	}

	// 6. Summary
	printSummary(results)
}

func runTest(arcBinary, arcFile, testName, tempDir string) TestResult {
	result := TestResult{Name: testName, Expected: 0}

	// Check filename for expected exit code (e.g. test_fail_exit1.arc)
	if strings.Contains(testName, "_exit") {
		parts := strings.Split(testName, "_exit")
		if len(parts) == 2 {
			fmt.Sscanf(parts[1], "%d", &result.Expected)
		}
	}

	objFile := filepath.Join(tempDir, testName+".o")
	exeFile := filepath.Join(tempDir, testName)

	// A. Compile (Arc -> Object)
	cmd := exec.Command(arcBinary, "build", arcFile, "-o", objFile)
	out, err := cmd.CombinedOutput()
	result.CompileLog = string(out)
	if err != nil {
		result.Error = "Compilation Error"
		return result
	}
	result.Compiled = true

	// B. Link (Object -> Exe using GCC/Clang for libc)
	cmd = exec.Command("gcc", objFile, "-o", exeFile, "-no-pie")
	out, err = cmd.CombinedOutput()
	result.LinkLog = string(out)
	if err != nil {
		result.Error = "Linking Error"
		return result
	}
	result.Linked = true

	// C. Run
	cmd = exec.Command(exeFile)
	out, err = cmd.CombinedOutput()
	result.RunLog = string(out)
	
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			result.ExitCode = exitErr.ExitCode()
		} else {
			result.Error = fmt.Sprintf("Execution Error: %v", err)
			return result
		}
	}

	if result.ExitCode == result.Expected {
		result.Passed = true
	} else {
		result.Error = fmt.Sprintf("Wrong Exit Code: expected %d, got %d", result.Expected, result.ExitCode)
	}

	return result
}

func fail(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	logToFile(msg + "\n")
	os.Exit(1)
}

func printHeader(count int) {
	header := "\nArc Foundation Test Runner\n" + strings.Repeat("=", 60) + "\n"
	fmt.Print(header)
	logToFile(header)
	fmt.Printf("Found %d tests in tests/foundation/\n\n", count)
}

func printResult(r TestResult) {
	if r.Passed {
		msg := fmt.Sprintf("✅ PASS (%.3fs)\n", r.Duration.Seconds())
		fmt.Print(msg)
		logToFile(msg)
	} else {
		msg := fmt.Sprintf("❌ FAIL (%.3fs)\n", r.Duration.Seconds())
		fmt.Print(msg)
		logToFile(msg)
		
		detail := ""
		if !r.Compiled {
			detail = fmt.Sprintf("   Compilation Failed:\n%s\n", indent(r.CompileLog))
		} else if !r.Linked {
			detail = fmt.Sprintf("   Linking Failed:\n%s\n", indent(r.LinkLog))
		} else {
			detail = fmt.Sprintf("   %s\n", r.Error)
			if r.RunLog != "" {
				detail += fmt.Sprintf("   Output:\n%s\n", indent(r.RunLog))
			}
		}
		fmt.Print(detail)
		logToFile(detail)
	}
}

func printSummary(results []TestResult) {
	passed, failed := 0, 0
	for _, r := range results {
		if r.Passed { passed++ } else { failed++ }
	}
	
	line := strings.Repeat("-", 60) + "\n"
	fmt.Print(line)
	logToFile(line)
	
	summary := fmt.Sprintf("Passed: %d | Failed: %d | Total: %d\n", passed, failed, len(results))
	fmt.Print(summary)
	logToFile(summary)
	
	if failed > 0 {
		os.Exit(1)
	}
}

func logToFile(msg string) {
	if logFile != nil { logFile.WriteString(msg) }
}

func indent(s string) string {
	return "      " + strings.ReplaceAll(strings.TrimSpace(s), "\n", "\n      ")
}