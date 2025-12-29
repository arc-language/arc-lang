package main

func init() {
	// Vector tests
	RegisterTest("Vector Type", func() string {
		return `let v: vector<int32> = {}`
	})

	RegisterTest("Vector Literal Empty", func() string {
		return `let empty: vector<int32> = {}`
	})

	RegisterTest("Vector Literal With Values", func() string {
		return `let nums: vector<int32> = {1, 2, 3, 4, 5}`
	})

	RegisterTest("Vector Literal Inferred", func() string {
		return `let items = {10, 20, 30}`
	})

	// Map tests
	RegisterTest("Map Type", func() string {
		return `let m: map<string, int32> = {}`
	})

	RegisterTest("Map Literal Empty", func() string {
		return `let empty: map<string, int32> = {}`
	})

	RegisterTest("Map Literal With Values", func() string {
		return `let scores: map<string, int32> = {"alice": 100, "bob": 95}`
	})

	RegisterTest("Map Literal Inferred", func() string {
		return `let config = {"host": "localhost", "port": "8080"}`
	})
}