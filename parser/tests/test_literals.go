package main

func init() {
	RegisterTest("Char Literal Escape Newline", func() string {
		return `let newline: char = '\n'`
	})

	RegisterTest("Char Literal Escape Tab", func() string {
		return `let tab: char = '\t'`
	})

	RegisterTest("Char Literal Escape Backslash", func() string {
		return `let backslash: char = '\\'`
	})

	RegisterTest("Char Literal Escape Quote", func() string {
		return `let quote: char = '\''`
	})

	RegisterTest("Char Literal Escape Null", func() string {
		return `let null_char: char = '\0'`
	})

	RegisterTest("String Literal Escapes", func() string {
		return `func test() {
    let msg: string = "Hello\nWorld"
    let path: string = "C:\\Users\\file"
    let quote: string = "He said \"hello\""
    let tab: string = "Column1\tColumn2"
}`
	})
}