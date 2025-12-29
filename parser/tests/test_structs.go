package main

func init() {
	RegisterTest("Struct Declaration", func() string {
		return `struct Point {
    x: int32
    y: int32
}`
	})

	RegisterTest("Struct Literal", func() string {
		return `let p = Point{x: 10, y: 20}`
	})

	RegisterTest("Struct Literal Type Inference", func() string {
		return `let p2 = Point{x: 5, y: 15}`
	})

	RegisterTest("Struct Literal Empty", func() string {
		return `let p3: Point = Point{}`
	})

	RegisterTest("Struct Field Access", func() string {
		return `let x = p.x`
	})

	RegisterTest("Struct Field Assignment", func() string {
		return `func test() {
    p.y = 30
}`
	})

	RegisterTest("Struct Inline Method", func() string {
		return `struct Point {
    x: int32
    y: int32
    
    func distance(self p: Point) float64 {
        return cast<float64>(p.x * p.x + p.y * p.y)
    }
    
    func move(self p: *Point, dx: int32, dy: int32) {
        p.x = p.x + dx
        p.y = p.y + dy
    }
}`
	})

	RegisterTest("Struct Mutating Method Inline", func() string {
		return `struct Counter {
    count: int32
    
    mutating increment(self c: *Counter) {
        c.count = c.count + 1
    }
    
    mutating add(self c: *Counter, value: int32) {
        c.count = c.count + value
    }
    
    func get_count(self c: Counter) int32 {
        return c.count
    }
}`
	})

	RegisterTest("Struct Mutating Method Usage", func() string {
		return `func test() {
    let counter = Counter{count: 0}
    counter.increment()
    counter.add(5)
    let value = counter.get_count()
}`
	})

	RegisterTest("Struct Flat Method", func() string {
		return `struct Point {
    x: int32
    y: int32
}

func distance(self p: Point) float64 {
    return cast<float64>(p.x * p.x + p.y * p.y)
}

func move(self p: *Point, dx: int32, dy: int32) {
    p.x = p.x + dx
    p.y = p.y + dy
}`
	})

	RegisterTest("Struct Mutating Method Flat", func() string {
		return `struct Point {
    x: int32
    y: int32
}

mutating reset(self p: *Point) {
    p.x = 0
    p.y = 0
}`
	})
}