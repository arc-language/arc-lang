# External Imports (`import c`, `import cpp`, `import objc`)

Arc unifies external dependency management under the `import` keyword. The language prefix (`c`, `cpp`, `objc`) specifies the ABI, the name specifies the library.

## Syntax Overview

```arc
import c "sqlite3"          // Auto-resolves to platform package
import c "curl"             // Auto-resolves based on ax.mod
import cpp "boost"          // C++ library, auto-resolved
import objc "AppKit"        // macOS framework
import "github.com/user/arclib"  // Arc package (no prefix)
```

**The compiler automatically resolves package sources based on your `ax.mod` configuration and detected OS.**

## How It Works

```
┌─────────────────────────────────────────────────────────────┐
│  Source                                                     │
├─────────────────────────────────────────────────────────────┤
│  import c "sqlite3"                                         │
│  import c "curl"                                            │
│  import objc "AppKit"                                       │
│  import "github.com/arc-lang/io"                            │
└─────────────────────────────────────────────────────────────┘
                            │
                            ▼
┌─────────────────────────────────────────────────────────────┐
│  ax.mod Configuration                                       │
├─────────────────────────────────────────────────────────────┤
│  require c (                                                │
│      sqlite3 v3.36 (                                        │
│          debian  "debian.org/libsqlite3-dev"                │
│          ubuntu  "ubuntu.org/libsqlite3-dev"                │
│          macos   "brew.sh/sqlite"                           │
│          windows "vcpkg.io/sqlite3"                         │
│          default "vcpkg.io/sqlite3"                         │
│      )                                                      │
│  )                                                          │
└─────────────────────────────────────────────────────────────┘
                            │
                            ▼
┌─────────────────────────────────────────────────────────────┐
│  Compiler Auto-Resolution                                   │
├─────────────────────────────────────────────────────────────┤
│  1. Parse imports: ["sqlite3", "curl"]                      │
│  2. Load ax.mod mappings                                    │
│  3. Detect OS: ubuntu 22.04                                 │
│  4. Resolve: sqlite3 → ubuntu.org/libsqlite3-dev            │
│  5. Build URL: https://ubuntu.org/packages/...              │
│  6. Download to: ~/.arc/cache/ubuntu.org/sqlite3/3.36/      │
│  7. Extract libs to: ~/.arc/libs/                           │
└─────────────────────────────────────────────────────────────┘
                            │
                            ▼
┌─────────────────────────────────────────────────────────────┐
│  Linker (auto-generated)                                    │
├─────────────────────────────────────────────────────────────┤
│  -L~/.arc/libs -lsqlite3 -lcurl -framework AppKit           │
└─────────────────────────────────────────────────────────────┘
```

---

## Import Types

### C/C++ Libraries (Auto-Resolved)

```arc
import c "sqlite3"      // Resolves via ax.mod
import c "curl"         // Resolves via ax.mod  
import c "openssl"      // Resolves via ax.mod
import cpp "boost"      // C++ library, resolves via ax.mod
import cpp "fmt"        // Modern C++ library
```

**The compiler automatically:**
1. Detects your OS (Ubuntu, macOS, Windows, etc.)
2. Looks up the mapping in `ax.mod`
3. Downloads from the appropriate CDN
4. Caches in `~/.arc/cache/`
5. Links automatically

### macOS Frameworks

```arc
import objc "Foundation"
import objc "AppKit"
import objc "Metal"
import objc "CoreGraphics"
import objc "AVFoundation"
```

Frameworks are auto-linked with `-framework` flags.

### Arc Packages

```arc
import "github.com/arc-lang/io"
import "github.com/user/physics@2.0"
import "gitlab.com/company/internal"
```

No prefix = Arc package from source control.

---

## Configuration: `ax.mod`

Your `ax.mod` file defines **where** to get packages for each platform:

```go
module github.com/myapp/example
arc 1.0

require c (
    sqlite3 v3.36 (
        debian   "debian.org/libsqlite3-dev"
        ubuntu   "ubuntu.org/libsqlite3-dev"
        macos    "brew.sh/sqlite"
        windows  "vcpkg.io/sqlite3"
        default  "vcpkg.io/sqlite3"      // Fedora, Arch, Alpine, etc.
    )
    
    curl v7.80 (
        debian   "debian.org/libcurl4-openssl-dev"
        ubuntu   "ubuntu.org/libcurl4-openssl-dev"
        macos    "brew.sh/curl"
        windows  "vcpkg.io/curl"
        default  "vcpkg.io/curl"
    )
)

require cpp (
    boost v1.82 (
        debian   "debian.org/libboost-all-dev"
        ubuntu   "ubuntu.org/libboost-all-dev"
        macos    "brew.sh/boost"
        windows  "vcpkg.io/boost"
        default  "vcpkg.io/boost"
    )
)
```

**You write this once**, then all imports auto-resolve.

---

## Source Code: Clean and Simple

```arc
namespace main

// Simple imports - no URLs, no package details
import c "sqlite3"
import c "curl"
import objc "AppKit"
import "github.com/arc-lang/io"

// Declarations (no linking directives needed)
extern c {
    opaque struct sqlite3 {}
    opaque struct CURL {}
    
    const SQLITE_OK: int32 = 0
    
    func sqlite3_open(*byte, **sqlite3) int32
    func sqlite3_close(*sqlite3) int32
    func sqlite3_errmsg(*sqlite3) *byte
    func curl_easy_init() *CURL
}

extern objc {
    struct NSRect { origin: NSPoint; size: NSSize }
    struct NSPoint { x: float64; y: float64 }
    struct NSSize { width: float64; height: float64 }
    
    class NSApplication {
        static func sharedApplication() *NSApplication
        func run(self *NSApplication) void
    }
}

func main() {
    // Use SQLite
    let db: *sqlite3 = null
    if sqlite3_open("app.db", &db) != SQLITE_OK {
        io.printf("Error: %s\n", sqlite3_errmsg(db))
        return
    }
    defer sqlite3_close(db)
    
    // Use cURL
    let curl = curl_easy_init()
    
    // Use AppKit
    let app = NSApplication.sharedApplication()
    app.run()
}
```

**Key Points:**
- ✅ Just `import c "sqlite3"` - no URLs or versions
- ✅ `extern` blocks contain only declarations
- ✅ No `lib` or `framework` directives
- ✅ Platform differences handled by `ax.mod`
- ✅ Code is clean and portable

---

## Auto-Resolution Examples

### On Ubuntu

```bash
$ arc build
Detected OS: ubuntu 22.04 (linux-x86_64)
Resolving dependencies...

[1/2] sqlite3 v3.36
  ✓ Auto-resolved: ubuntu → ubuntu.org/libsqlite3-dev
  ✓ URL: https://ubuntu.org/packages/libsqlite3-dev/3.36/linux-x86_64.tar.gz
  → Downloading... [====================] 2.3 MB
  ✓ Cached: ~/.arc/cache/ubuntu.org/libsqlite3-dev/3.36/
  ✓ Linker: -lsqlite3

[2/2] curl v7.80
  ✓ Auto-resolved: ubuntu → ubuntu.org/libcurl4-openssl-dev
  ✓ Cache hit: ~/.arc/cache/ubuntu.org/libcurl4-openssl-dev/7.80/
  ✓ Linker: -lcurl

Building project...
Build complete: ./build/myapp
```

### On macOS

```bash
$ arc build
Detected OS: macOS 14.2 (darwin-arm64)
Resolving dependencies...

[1/2] sqlite3 v3.36
  ✓ Auto-resolved: macos → brew.sh/sqlite
  ✓ URL: https://brew.sh/bottles/sqlite/3.43.2/darwin-arm64.tar.gz
  → Downloading... [====================] 1.8 MB
  ✓ Cached: ~/.arc/cache/brew.sh/sqlite/3.43.2/
  ✓ Linker: -lsqlite3

[2/2] AppKit
  ✓ macOS framework: -framework AppKit

Building project...
Build complete: ./build/myapp
```

### On Arch Linux

```bash
$ arc build
Detected OS: Arch Linux (linux-x86_64)
Resolving dependencies...

[1/2] sqlite3 v3.36
  ✓ No specific mapping, using default → vcpkg.io/sqlite3
  ✓ URL: https://vcpkg.io/packages/sqlite3/3.36.0/linux-x86_64.tar.gz
  → Downloading... [====================] 2.0 MB
  ✓ Cached: ~/.arc/cache/vcpkg.io/sqlite3/3.36.0/
  ✓ Linker: -lsqlite3

Building project...
Build complete: ./build/myapp
```

---

## Versioning

Specified in `ax.mod`, not in imports:

```go
require c (
    sqlite3 v3.36 (...)       // Exact version
    curl v^7.80 (...)         // Compatible (>=7.80 <8.0)
    openssl v>=1.1.1 (...)    // Minimum version
)
```

---

## Deduplication

Compiler automatically deduplicates across files:

```arc
// file1.arc
import c "sqlite3"
import c "curl"

// file2.arc  
import c "sqlite3"    // Duplicate - resolved once
import c "openssl"

// file3.arc
import c "curl"       // Duplicate - resolved once
import objc "AppKit"
```

Final link command:
```bash
-L~/.arc/libs -lsqlite3 -lcurl -lopenssl -framework AppKit
```

---

## Cache Location

All packages download to:

```
~/.arc/
├── cache/                   # Downloaded packages by source
│   ├── debian.org/
│   │   └── libsqlite3-dev/3.36/linux-x86_64/
│   ├── ubuntu.org/
│   │   └── libsqlite3-dev/3.36/linux-x86_64/
│   ├── brew.sh/
│   │   └── sqlite/3.43.2/darwin-arm64/
│   └── vcpkg.io/
│       └── sqlite3/3.36.0/
│           ├── linux-x86_64/      # Arch, Fedora, etc.
│           ├── windows-x86_64/
│           └── darwin-arm64/
│
├── libs/                    # Extracted libraries
│   ├── libsqlite3.a
│   ├── libsqlite3.so
│   └── libcurl.a
│
└── include/                 # Extracted headers
    ├── sqlite3.h
    └── curl/
```

Override with:
```bash
export ARC_LIBS=/custom/path
```

---

## Platform Coverage

The auto-resolution system supports:

**Direct Mappings:**
- **Debian/Ubuntu** → `debian.org`, `ubuntu.org`
- **macOS** → `brew.sh`
- **Windows** → `vcpkg.io`
- **NixOS** (optional) → `nixos.org`

**Universal Fallback:**
- **All other distros** → `vcpkg.io` (Fedora, Arch, Alpine, FreeBSD, etc.)

**Or use `vcpkg.io` everywhere for maximum consistency:**

```go
require c (
    sqlite3 v3.36 (
        default "vcpkg.io/sqlite3"    // Works on all platforms
    )
)
```

---

## Complete Example

### Source Code (`main.arc`)

```arc
namespace main

import c "sqlite3"
import objc "AppKit"
import "github.com/arc-lang/io"

extern c {
    opaque struct sqlite3 {}
    const SQLITE_OK: int32 = 0
    func sqlite3_open(*byte, **sqlite3) int32
    func sqlite3_close(*sqlite3) int32
}

extern objc {
    struct NSRect { origin: NSPoint; size: NSSize }
    struct NSPoint { x: float64; y: float64 }
    struct NSSize { width: float64; height: float64 }
    func NSMakeRect(float64, float64, float64, float64) NSRect
    
    class NSApplication {
        static func sharedApplication() *NSApplication
        func run(self *NSApplication) void
    }
    
    class NSWindow {
        new "initWithContentRect:styleMask:backing:defer:" (NSRect, uint64, uint64, bool) *NSWindow
        property title: *NSString
        func center(self *NSWindow) void
        func makeKeyAndOrderFront "makeKeyAndOrderFront:" (self *NSWindow, *id) void
    }
    
    class NSString {
        static func stringWithUTF8String "stringWithUTF8String:" (*byte) *NSString
    }
    
    opaque class id {}
}

func main() {
    // SQLite auto-resolved for your platform
    let db: *sqlite3 = null
    if sqlite3_open("app.db", &db) != SQLITE_OK {
        io.printf("Database error\n")
        return
    }
    defer sqlite3_close(db)
    
    // AppKit auto-linked on macOS
    let app = NSApplication.sharedApplication()
    let window = NSWindow.new(NSMakeRect(0, 0, 800, 600), 3, 2, false)
    window.title = NSString.stringWithUTF8String("Arc App")
    window.center()
    window.makeKeyAndOrderFront(null)
    app.run()
}
```

### Build Config (`ax.mod`)

```go
module github.com/myapp/example
arc 1.0

require (
    github.com/arc-lang/io v1.2
)

require c (
    sqlite3 v3.36 (
        debian   "debian.org/libsqlite3-dev"
        ubuntu   "ubuntu.org/libsqlite3-dev"
        macos    "brew.sh/sqlite"
        windows  "vcpkg.io/sqlite3"
        default  "vcpkg.io/sqlite3"
    )
)
```

That's it! Run `arc build` and everything auto-resolves.

---

## Quick Reference

| Import Statement | What It Does |
|-----------------|--------------|
| `import c "sqlite3"` | Auto-resolves via `ax.mod` to platform package |
| `import cpp "boost"` | Auto-resolves C++ library |
| `import objc "AppKit"` | Auto-links macOS framework |
| `import "github.com/x/y"` | Fetches Arc package |

**Configuration is in `ax.mod`, not in code.**

---

## Summary

✅ **Write once:** `import c "sqlite3"`  
✅ **Configure once:** Platform mappings in `ax.mod`  
✅ **Runs everywhere:** Auto-resolves based on detected OS  
✅ **No URLs in code:** Clean, portable source files  
✅ **No manual linking:** Compiler handles everything  
✅ **Cached locally:** `~/.arc/cache/` for fast rebuilds  
✅ **No sudo needed:** Everything in user space

The compiler knows where to download packages from based on your `ax.mod` configuration and automatically detects your platform.