package main

func init() {
	// Import tests
	RegisterTest("Import Single", func() string {
		return `import "some/path/package"`
	})

	RegisterTest("Import Multiple", func() string {
		return `import (
    "std/io"
    "std/os"
    "github.com/user/physics"
    "github.com/user/graphics"
    "gitlab.com/company/lib"
)`
	})

	RegisterTest("Namespace", func() string {
		return `namespace main`
	})

	// Variable and constant tests
	RegisterTest("Variable Mutable Typed", func() string {
		return `let x: int32 = 42`
	})

	RegisterTest("Variable Mutable Inferred", func() string {
		return `let x = 42`
	})

	RegisterTest("Constant Typed", func() string {
		return `const x: int32 = 42`
	})

	RegisterTest("Constant Inferred", func() string {
		return `const x = 42`
	})
}