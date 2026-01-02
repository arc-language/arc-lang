package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// TestCase defines a unit of work for the runner.
// Tests are defined in separate files (e.g., tests_foundation.go) using init().
type TestCase struct {
	Name     string
	Globals  string // Structs, global functions, imports
	Body     string // Code inserted into the main() function
	Expected string // Substring expected in standard output
}

// AllTests is populated by init() functions in other files in this package.
var AllTests []TestCase

func main() {
	// 1. Setup Logging
	logFile, err := os.Create("tests.log")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating log file: %v\n", err)
		os.Exit(1)
	}
	defer logFile.Close()

	// 2. Check for Compiler Binary
	arcBinary := "./arc"
	if _, err := os.Stat(arcBinary); os.IsNotExist(err) {
		fail(logFile, "Error: './arc' binary not found. Please build the compiler first (go build -o arc main.go).")
	}

	// 3. Prepare Temp Directory
	tempDir := filepath.Join(os.TempDir(), "arc_test_env")
	os.RemoveAll(tempDir) // Clean start
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		fail(logFile, fmt.Sprintf("Error creating temp dir: %v", err))
	}
	defer os.RemoveAll(tempDir)

	// 4. Run Header
	header := fmt.Sprintf("Running %d tests...\n", len(AllTests))
	fmt.Print(header)
	writeLog(logFile, header)

	passed := 0
	failed := 0
	startTotal := time.Now()

	// 5. Execute Tests
	for i, test := range AllTests {
		prefix := fmt.Sprintf("[%d/%d] %-35s ", i+1, len(AllTests), test.Name)
		fmt.Print(prefix)
		writeLog(logFile, fmt.Sprintf("\n--- Test: %s ---\n", test.Name))

		start := time.Now()
		err := runSingleTest(test, arcBinary, tempDir, logFile)
		duration := time.Since(start)

		if err == nil {
			msg := fmt.Sprintf("✅ PASS (%.3fs)\n", duration.Seconds())
			fmt.Print(msg)
			writeLog(logFile, "Result: PASS\n")
			passed++
		} else {
			msg := fmt.Sprintf("❌ FAIL (%.3fs)\n", duration.Seconds())
			fmt.Print(msg)
			writeLog(logFile, fmt.Sprintf("Result: FAIL\nError: %v\n", err))
			// Print error to console lightly so we don't spam if it's huge
			fmt.Printf("      Error: %v\n", err)
			failed++
		}
	}

	// 6. Summary
	fmt.Println(strings.Repeat("-", 60))
	summary := fmt.Sprintf("Passed: %d | Failed: %d | Total Time: %.2fs\n", 
		passed, failed, time.Since(startTotal).Seconds())
	fmt.Print(summary)
	writeLog(logFile, "\n"+strings.Repeat("=", 60)+"\n")
	writeLog(logFile, summary)

	if failed > 0 {
		os.Exit(1)
	}
}

func runSingleTest(tc TestCase, arcBinary, tempDir string, log *os.File) error {
	// 1. Construct Source Code
	// We inject standard libc definitions automatically to keep test cases clean.
	fullSource := fmt.Sprintf(`
extern libc {
    func printf(fmt: *byte, ...) int32
    func malloc(size: usize) *void
    func free(ptr: *void)
    func memcpy(dest: *void, src: *void, count: usize) *void
    func memset(dest: *void, val: int32, count: usize) *void
    func strlen(s: *byte) usize
}

// --- Test Globals ---
%s

func main() int32 {
    // --- Test Body ---
%s
    return 0
}
`, tc.Globals, tc.Body)

	srcPath := filepath.Join(tempDir, "test.arc")
	objPath := filepath.Join(tempDir, "test.o")
	exePath := filepath.Join(tempDir, "test_exe")

	// 2. Write Source File
	if err := os.WriteFile(srcPath, []byte(fullSource), 0644); err != nil {
		return fmt.Errorf("write source failed: %v", err)
	}
	writeLog(log, "Source Code:\n"+fullSource+"\n")

	// 3. Compile (Arc -> Object)
	cmdCompile := exec.Command(arcBinary, "build", srcPath, "-o", objPath)
	outCompile, err := cmdCompile.CombinedOutput()
	if err != nil {
		writeLog(log, "Compile Output:\n"+string(outCompile))
		return fmt.Errorf("compile failed: %s", strings.TrimSpace(string(outCompile)))
	}

	// 4. Link (Object -> Executable using GCC)
	cmdLink := exec.Command("gcc", objPath, "-o", exePath, "-no-pie")
	outLink, err := cmdLink.CombinedOutput()
	if err != nil {
		writeLog(log, "Link Output:\n"+string(outLink))
		return fmt.Errorf("link failed: %s", strings.TrimSpace(string(outLink)))
	}

	// 5. Run Executable
	cmdRun := exec.Command(exePath)
	outRun, err := cmdRun.CombinedOutput()
	output := string(outRun)
	writeLog(log, "Run Output:\n"+output)

	if err != nil {
		// Differentiate between crash and non-zero exit if possible, 
		// though usually we return 0 in main template.
		return fmt.Errorf("runtime error: %v\nOutput: %s", err, output)
	}

	// 6. Assert Expectations
	if !strings.Contains(output, tc.Expected) {
		return fmt.Errorf("assertion failed.\nExpected substring: %q\nGot output: %q", tc.Expected, output)
	}

	return nil
}

func writeLog(f *os.File, msg string) {
	if f != nil {
		f.WriteString(msg)
	}
}

func fail(log *os.File, msg string) {
	fmt.Fprintln(os.Stderr, msg)
	writeLog(log, msg+"\n")
	os.Exit(1)
}