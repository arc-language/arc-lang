# External Imports (`import c`, `import cpp`, `import objc`)

Arc unifies external dependency management under the `import` keyword. The language prefix (`c`, `cpp`, `objc`) specifies the ABI, the path specifies the source.

## Syntax Overview

```arc
import c "sqlite3"                     // System C library
import c "vcpkg.io/sqlite3@3.36"       // vcpkg with version
import cpp "vcpkg.io/boost-filesystem" // C++ library
import objc "AppKit"                   // macOS framework
import "github.com/user/arclib"        // Arc package (no prefix)
import c "nix.org/sqlite3@3.36"       // vcpkg with version
import c "brew.sh/sqlite3"              // Latest
import c "brew.sh/sqlite@3.36"          // Specific version

```

## How It Works

```
┌─────────────────────────────────────────────────────────────┐
│  Source                                                     │
├─────────────────────────────────────────────────────────────┤
│  import c "vcpkg.io/sqlite3@3.36"                           │
│  import c "pthread"                                         │
│  import objc "AppKit"                                       │
│  import "github.com/arc-lang/io"                            │
└─────────────────────────────────────────────────────────────┘
                            │
                            ▼
┌─────────────────────────────────────────────────────────────┐
│  Compiler Resolution                                        │
├─────────────────────────────────────────────────────────────┤
│  vcpkg.io/sqlite3@3.36  →  fetch vcpkg, install to ~/.arc/  │
│  pthread                →  system lib                       │
│  AppKit                 →  macOS framework                  │
│  github.com/arc-lang/io →  fetch Arc package                │
└─────────────────────────────────────────────────────────────┘
                            │
                            ▼
┌─────────────────────────────────────────────────────────────┐
│  Linker (deduplicated)                                      │
├─────────────────────────────────────────────────────────────┤
│  -L~/.arc/libs -lsqlite3 -lpthread -framework AppKit        │
└─────────────────────────────────────────────────────────────┘
```

---

## Import Types

### System Libraries

```arc
import c "sqlite3"      // -lsqlite3
import c "m"            // -lm (math)
import c "pthread"      // -lpthread
import cpp "stdc++"     // -lstdc++
```

### vcpkg Packages

```arc
import c "vcpkg.io/sqlite3"              // Latest
import c "vcpkg.io/sqlite3@3.36"         // Specific version
import c "vcpkg.io/openssl@1.1.1"
import cpp "vcpkg.io/boost-filesystem@1.82"
import cpp "vcpkg.io/fmt@10.1"
```

### macOS Frameworks

```arc
import objc "Foundation"
import objc "AppKit"
import objc "Metal"
import objc "CoreGraphics"
import objc "AVFoundation"
```

### Local Libraries

```arc
import c "./vendor/mylib"          // Relative path
import c "/opt/custom/libfoo"      // Absolute path
import cpp "../shared/engine"
```

### Arc Packages

```arc
import "github.com/arc-lang/io"
import "github.com/user/physics@2.0"
import "gitlab.com/company/internal"
```

No prefix = Arc package.

---

## Declarations

`import` handles linking. `extern` blocks provide declarations:

```arc
import c "vcpkg.io/sqlite3@3.36"

extern c {
    // Declarations only - no 'lib' directive needed
    opaque struct sqlite3 {}
    opaque struct sqlite3_stmt {}
    
    const SQLITE_OK: int32 = 0
    const SQLITE_ROW: int32 = 100
    
    func sqlite3_open(*byte, **sqlite3) int32
    func sqlite3_close(*sqlite3) int32
    func sqlite3_prepare_v2(*sqlite3, *byte, int32, **sqlite3_stmt, **byte) int32
    func sqlite3_step(*sqlite3_stmt) int32
    func sqlite3_finalize(*sqlite3_stmt) int32
}
```

```arc
import objc "AppKit"

extern objc {
    // Declarations only - no 'framework' directive needed
    class NSApplication {
        static func sharedApplication() *NSApplication
        func run(self *NSApplication) void
    }
    
    class NSWindow {
        new "initWithContentRect:styleMask:backing:defer:" (NSRect, uint64, uint64, bool) *NSWindow
        property title: *NSString
    }
}
```

---

## Versioning

```arc
import c "vcpkg.io/sqlite3"           // Latest
import c "vcpkg.io/sqlite3@3.36"      // Exact version
import c "vcpkg.io/sqlite3@^3.36"     // Compatible (>=3.36 <4.0)
import c "vcpkg.io/sqlite3@>=3.30"    // Minimum version
```

---

## Deduplication

Compiler automatically deduplicates across files:

```arc
// file1.arc
import c "vcpkg.io/sqlite3@3.36"
import c "pthread"

// file2.arc
import c "vcpkg.io/sqlite3@3.36"    // Duplicate - ignored
import c "m"

// file3.arc
import c "pthread"                   // Duplicate - ignored
import objc "AppKit"
```

Compiler collects:
```
c libs: ["sqlite3", "pthread", "m"]
frameworks: ["AppKit"]
```

---

## Artifact Location

Standard location for predictable builds:

```
~/.arc/
├── libs/                    # Compiled/downloaded libraries
│   ├── libsqlite3.so
│   ├── libsqlite3.a
│   └── libboost_filesystem.a
│
├── include/                 # Headers (for C/C++ interop)
│   └── sqlite3.h
│
├── cache/                   # Downloaded packages
│   └── vcpkg.io/
│       └── sqlite3/
│           └── 3.36/
│
└── bin/                     # Installed tools
```

Override with environment variable:

```bash
export ARC_LIBS=/opt/custom/libs
arc build
```

Or in `build.arc`:

```arc
module myapp {
    lib_paths: ["/opt/vendor/libs"]
}
```

---

## Complete Example

```arc
namespace main

// External dependencies
import c "vcpkg.io/sqlite3@3.36"
import objc "AppKit"

// Arc packages
import "github.com/arc-lang/io"

// C declarations
extern c {
    opaque struct sqlite3 {}
    const SQLITE_OK: int32 = 0
    func sqlite3_open(*byte, **sqlite3) int32
    func sqlite3_close(*sqlite3) int32
    func sqlite3_errmsg(*sqlite3) *byte
}

// ObjC declarations
extern objc {
    struct NSRect { origin: NSPoint; size: NSSize }
    struct NSPoint { x: float64; y: float64 }
    struct NSSize { width: float64; height: float64 }
    func NSMakeRect(float64, float64, float64, float64) NSRect
    
    class NSApplication {
        static func sharedApplication() *NSApplication
        property delegate: *id
        func run(self *NSApplication) void
        func setActivationPolicy "setActivationPolicy:" (self *NSApplication, int64) bool
    }
    
    class NSWindow {
        new "initWithContentRect:styleMask:backing:defer:" (NSRect, uint64, uint64, bool) *NSWindow
        property title: *NSString
        func makeKeyAndOrderFront "makeKeyAndOrderFront:" (self *NSWindow, *id) void
        func center(self *NSWindow) void
    }
    
    class NSString {
        static func stringWithUTF8String "stringWithUTF8String:" (*byte) *NSString
    }
    
    opaque class id {}
}

func main() {
    // Use SQLite
    let db: *sqlite3 = null
    if sqlite3_open("app.db", &db) != SQLITE_OK {
        io.printf("Error: %s\n", sqlite3_errmsg(db))
        return
    }
    defer sqlite3_close(db)
    
    // Use AppKit
    let app = NSApplication.sharedApplication()
    app.setActivationPolicy(0)
    
    let window = NSWindow.new(NSMakeRect(0, 0, 800, 600), 3, 2, false)
    window.title = NSString.stringWithUTF8String("Arc App")
    window.center()
    window.makeKeyAndOrderFront(null)
    
    app.run()
}
```

---

## Quick Reference

| Import | Resolves To |
|--------|-------------|
| `import c "sqlite3"` | `-lsqlite3` (system) |
| `import c "vcpkg.io/sqlite3@3.36"` | Fetch vcpkg → `-lsqlite3` |
| `import cpp "vcpkg.io/boost"` | Fetch vcpkg → `-lboost_*` |
| `import objc "AppKit"` | `-framework AppKit` |
| `import "./vendor/lib"` | `-L./vendor -llib` |
| `import "github.com/x/y"` | Fetch Arc package |

---

## Summary

- `import <abi> "source"` — external libraries (C, C++, ObjC)
- `import "source"` — Arc packages
- `extern` blocks — declarations only, no linking directives
- Compiler deduplicates across all source files
- Standard artifact location: `~/.arc/libs/`