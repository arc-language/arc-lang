package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// TestCase defines a unit of work for the runner.
type TestCase struct {
	Name     string
	Phase    int    // 1=Foundation, 2=Operators, 3=Control, etc.
	Globals  string // Structs, global functions, imports
	Body     string // Code inserted into the main() function
	Expected string // Substring expected in standard output
}

// AllTests is populated by init() functions in other test files
var AllTests []TestCase

// Test execution timeout (per test)
const TEST_TIMEOUT = 10 * time.Second

// LibPath defines where libc is located.
// Your provided path: /usr/lib/gcc/x86_64-linux-gnu/13/../../../x86_64-linux-gnu/
// resolves to: /usr/lib/x86_64-linux-gnu/
const LIB_PATH = "/usr/lib/x86_64-linux-gnu"

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
		fail(logFile, "Error: './arc' binary not found. Please build the compiler first.")
	}

	// 3. Prepare Temp Directory
	tempDir := filepath.Join(os.TempDir(), "arc_test_env")
	os.RemoveAll(tempDir) // Clean start
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		fail(logFile, fmt.Sprintf("Error creating temp dir: %v", err))
	}
	defer os.RemoveAll(tempDir)

	// 4. Sort tests by Phase first, then Name within each phase
	sort.Slice(AllTests, func(i, j int) bool {
		if AllTests[i].Phase != AllTests[j].Phase {
			return AllTests[i].Phase < AllTests[j].Phase
		}
		return AllTests[i].Name < AllTests[j].Name
	})

	// 5. Run Header
	header := fmt.Sprintf("Arc Compiler Test Suite - Running %d tests...\n", len(AllTests))
	fmt.Print(header)
	fmt.Println(strings.Repeat("=", 60))
	writeLog(logFile, header)

	passed := 0
	failed := 0
	startTotal := time.Now()

	// 6. Execute Tests
	for i, test := range AllTests {
		prefix := fmt.Sprintf("[%d/%d] %-40s ", i+1, len(AllTests), test.Name)
		fmt.Print(prefix)
		writeLog(logFile, fmt.Sprintf("\n--- Test: %s ---\n", test.Name))

		start := time.Now()
		output, err := runSingleTest(test, arcBinary, tempDir, logFile)
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
			fmt.Printf("      Error: %v\n", err)
			failed++
		}

		// Display execution output
		if output != "" {
			fmt.Printf("      Output: %s\n", strings.TrimSpace(output))
		}
	}

	// 7. Summary
	fmt.Println(strings.Repeat("=", 60))
	summary := fmt.Sprintf("Passed: %d | Failed: %d | Total Time: %.2fs\n",
		passed, failed, time.Since(startTotal).Seconds())
	fmt.Print(summary)
	writeLog(logFile, "\n"+strings.Repeat("=", 60)+"\n")
	writeLog(logFile, summary)

	if failed > 0 {
		os.Exit(1)
	}
}

func runSingleTest(tc TestCase, arcBinary, tempDir string, log *os.File) (string, error) {
	// 1. Construct Source Code
	fullSource := fmt.Sprintf(`
extern c {
    func printf(*byte, ...) int32
    func malloc(usize) *void
    func free(*void)
    func memcpy(*void, *void, usize) *void
    func memset(*void, int32, usize) *void
    func strlen(*byte) usize
    func sleep(int32) int32
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
	exePath := filepath.Join(tempDir, "test_exe")

	// 2. Write Source File
	if err := os.WriteFile(srcPath, []byte(fullSource), 0644); err != nil {
		return "", fmt.Errorf("write source failed: %v", err)
	}
	writeLog(log, "Source Code:\n"+fullSource+"\n")

	// 3. Build (Compile + Link) using Arc
	ctx, cancel := context.WithTimeout(context.Background(), TEST_TIMEOUT)
	defer cancel()

	// Command: ./arc build test.arc -o test_exe -L /usr/lib/... -l c
	cmdBuild := exec.CommandContext(ctx, arcBinary, 
		"build", srcPath, 
		"-o", exePath, 
		"-L", LIB_PATH, 
		"-l", "c",
	)
	
	outBuild, err := cmdBuild.CombinedOutput()
	
	if ctx.Err() == context.DeadlineExceeded {
		writeLog(log, "Build Output: TIMEOUT\n")
		return "", fmt.Errorf("build timeout (exceeded %v)", TEST_TIMEOUT)
	}
	
	if err != nil {
		writeLog(log, "Build Output:\n"+string(outBuild))
		return "", fmt.Errorf("build failed: %s", strings.TrimSpace(string(outBuild)))
	}

	// 4. Run Executable with timeout
	ctx, cancel = context.WithTimeout(context.Background(), TEST_TIMEOUT)
	defer cancel()

	cmdRun := exec.CommandContext(ctx, exePath)
	outRun, err := cmdRun.CombinedOutput()
	output := string(outRun)
	writeLog(log, "Run Output:\n"+output)

	// Check for timeout
	if ctx.Err() == context.DeadlineExceeded {
		return output, fmt.Errorf("execution timeout (exceeded %v) - possible infinite loop or hang", TEST_TIMEOUT)
	}

	// Check for runtime errors (crashes, segfaults, etc.)
	if err != nil {
		// Try to determine if it was a crash vs normal error
		if exitErr, ok := err.(*exec.ExitError); ok {
			if exitErr.ExitCode() != 0 {
				return output, fmt.Errorf("runtime error (exit code %d)\nOutput: %s", exitErr.ExitCode(), output)
			}
		}
		return output, fmt.Errorf("runtime error: %v\nOutput: %s", err, output)
	}

	// 5. Assert Expectations
	if !strings.Contains(output, tc.Expected) {
		return output, fmt.Errorf("assertion failed.\nExpected substring: %q\nGot output: %q", tc.Expected, output)
	}

	return output, nil
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