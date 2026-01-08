package main

func init() {
	// Phase 6: Structs
	AllTests = append(AllTests,
		TestCase{
			Name:    "struct_definition",
			Globals: `
    struct Point {
        x: int32
        y: int32
    }
    `,
			Body: `
    printf("struct_defined\n")
    `,
			Expected: "struct_defined",
			Phase: 6,
		},
		TestCase{
			Name:    "struct_init_named",
			Globals: `
    struct Point {
        x: int32
        y: int32
    }
    `,
			Body: `
    let p: Point = Point{x: 10, y: 20}
    printf("x=%d y=%d\n", p.x, p.y)
    `,
			Expected: "x=10 y=20",
			Phase: 6,
		},
		TestCase{
			Name:    "struct_init_default",
			Globals: `
    struct Point {
        x: int32
        y: int32
    }
    `,
			Body: `
    let p: Point = Point{}
    printf("x=%d y=%d\n", p.x, p.y)
    `,
			Expected: "x=0 y=0",
			Phase: 6,
		},
		TestCase{
			Name:    "struct_field_read",
			Globals: `
    struct Rectangle {
        width: int32
        height: int32
    }
    `,
			Body: `
    let rect = Rectangle{width: 100, height: 50}
    printf("width=%d\n", rect.width)
    `,
			Expected: "width=100",
			Phase: 6,
		},
		TestCase{
			Name:    "struct_field_write",
			Globals: `
    struct Point {
        x: int32
        y: int32
    }
    `,
			Body: `
    let p = Point{x: 5, y: 10}
    p.x = 99
    printf("x=%d\n", p.x)
    `,
			Expected: "x=99",
			Phase: 6,
		},
		TestCase{
			Name:    "struct_pointer_field",
			Globals: `
    struct Point {
        x: int32
        y: int32
    }
    `,
			Body: `
    let p = Point{x: 15, y: 25}
    let ptr: *Point = &p
    printf("x=%d\n", ptr.x)
    `,
			Expected: "x=15",
			Phase: 6,
		},
		TestCase{
			Name:    "struct_nested",
			Globals: `
    struct Point {
        x: int32
        y: int32
    }
    struct Line {
        start: Point
        end: Point
    }
    `,
			Body: `
    let line = Line{
        start: Point{x: 0, y: 0},
        end: Point{x: 10, y: 10}
    }
    printf("end_x=%d\n", line.end.x)
    `,
			Expected: "end_x=10",
			Phase: 6,
		},
		TestCase{
			Name:    "struct_method_inline",
			Globals: `
    struct Counter {
        value: int32
        
        func get(self c: Counter) int32 {
            return c.value
        }
    }
    `,
			Body: `
    let counter = Counter{value: 42}
    let val = counter.get()
    printf("value=%d\n", val)
    `,
			Expected: "value=42",
			Phase: 6,
		},
		TestCase{
			Name:    "struct_method_flat",
			Globals: `
    struct Counter {
        value: int32
    }
    
    func get(self c: Counter) int32 {
        return c.value
    }
    `,
			Body: `
    let counter = Counter{value: 77}
    let val = counter.get()
    printf("value=%d\n", val)
    `,
			Expected: "value=77",
			Phase: 6,
		},
		TestCase{
			Name:    "struct_mutating_method",
			Globals: `
    struct Counter {
        count: int32
        
        mutating increment(self c: *Counter) {
            c.count++
        }
    }
    `,
			Body: `
    let counter = Counter{count: 10}
    counter.increment()
    printf("count=%d\n", counter.count)
    `,
			Expected: "count=11",
			Phase: 6,
		},
		TestCase{
			Name:    "struct_multiple_fields",
			Globals: `
    struct Person {
        age: int32
        height: int32
        weight: int32
    }
    `,
			Body: `
    let person = Person{age: 30, height: 180, weight: 75}
    printf("age=%d height=%d\n", person.age, person.height)
    `,
			Expected: "age=30 height=180",
			Phase: 6,
		},
		TestCase{
			Name:    "struct_pass_by_value",
			Globals: `
    struct Point {
        x: int32
        y: int32
    }
    
    func print_point(p: Point) {
        printf("x=%d y=%d\n", p.x, p.y)
    }
    `,
			Body: `
    let pt = Point{x: 5, y: 10}
    print_point(pt)
    `,
			Expected: "x=5 y=10",
			Phase: 6,
		},
		TestCase{
			Name:    "struct_pass_by_pointer",
			Globals: `
    struct Point {
        x: int32
        y: int32
    }
    
    func move_point(p: *Point, dx: int32, dy: int32) {
        p.x += dx
        p.y += dy
    }
    `,
			Body: `
    let pt = Point{x: 0, y: 0}
    move_point(&pt, 5, 10)
    printf("x=%d y=%d\n", pt.x, pt.y)
    `,
			Expected: "x=5 y=10",
			Phase: 6,
		},
	)
}