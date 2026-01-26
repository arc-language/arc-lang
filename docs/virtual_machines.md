# Arc Virtual Machine Compute Context (`vm`)

The `vm` keyword provides **Hardware Virtualization** as a first-class language construct.

Unlike `container` (which shares the host kernel via Namespaces), `vm` boots a separate Kernel using the host's native hypervisor APIs. This creates an **Arc Virtual Machine**, providing an isolated, clean slate for testing, building, and low-level system development.

## Quick Reference

```arc
// Define a fresh development environment
const dev_env = vm.config {
    vcpu: 4
    ram: "4GB"
    kernel: "linux-lts" 
    network: true
    drives: ["./src:/src:rw"] // Mount source code
}

func main() {
    // Spin up the VM, compile the project inside it, and return the binary
    let binary = vm(dev_env) func() []byte {
        // Runs inside the clean Linux VM
        system("make clean && make")
        return read_file("/src/bin/app")
    }()
    
    save_file("./app_release", binary)
}
```

---

## Native Architecture

Arc does not rely on heavy external dependencies. It interfaces directly with the Operating System's native hypervisor headers to create lightweight micro-VMs.

### 1. Linux: KVM (Kernel-based Virtual Machine)
On Linux, Arc interacts directly with the kernel via `ioctl`. It bypasses user-space emulators for maximum performance, acting as its own VMM (Virtual Machine Monitor) for simple workloads.

**The Low-Level C Interface:**
Arc abstracts the raw `ioctl` calls required to set up page tables and interrupt controllers.

```c
#include <linux/kvm.h>
#include <sys/ioctl.h>
#include <fcntl.h>

// Arc Runtime Internal Logic:
// 1. Open the KVM device
int kvm = open("/dev/kvm", O_RDWR | O_CLOEXEC);

// 2. Create the VM
int vm_fd = ioctl(kvm, KVM_CREATE_VM, 0);

// 3. Map memory and load Arc binary into Guest RAM...

// 4. Run the vCPU
// This enters "Guest Mode". The thread only returns when the VM exists (I/O or Halt).
ioctl(vcpu_fd, KVM_RUN, 0);
```

### 2. macOS: Virtualization.framework
On macOS, Arc links against the native Apple Silicon/Intel framework. This allows it to boot Linux with near-native performance without needing QEMU or kernel extensions.

**The Native Objective-C Interface:**
Arc uses the bridging header to access the `VZ` namespace.

```objective-c
#import <Virtualization/Virtualization.h>

// Arc Runtime Internal Logic:
VZVirtualMachineConfiguration *config = [[VZVirtualMachineConfiguration alloc] init];
config.bootLoader = [[VZEFIBootLoader alloc] init];
config.CPUCount = 4;

// Use VirtIO for high-performance storage/net
VZVirtioBlockDeviceConfiguration *disk = ...;

// Boot without GUI
VZVirtualMachine *vm = [[VZVirtualMachine alloc] initWithConfiguration:config];
[vm startWithCompletionHandler:^(NSError *err){ ... }];
```

### 3. Windows: Hyper-V (COM / HCS)
On Windows, Arc utilizes the Hyper-V Platform. It connects via the **Host Compute System (HCS)** API or the classic **WMI COM** interfaces to request partition resources.

**The C++/COM Interface:**
Arc handles the complex COM initialization and security descriptors.

```cpp
#include <windows.h>
#include <computecore.h> // Host Compute System API
#include <virtdisk.h>

// Arc Runtime Internal Logic:
// Initialize the Compute System (VM)
HCS_SYSTEM system;
HcsCreateComputeSystem(json_config, NULL, &system);

// Or via classic COM (WMI):
// CoCreateInstance(CLSID_Msvm_VirtualSystemManagementService, ...)
// Msvm_VirtualSystemSettingData...
```

---

## Developer Use Cases

The `vm` keyword is designed to solve "It works on my machine" problems and enable low-level development.

### 1. Clean Build Environments & Packaging
Ensure your software builds correctly on a pristine OS.
*   **Problem:** Your local machine has libraries installed that the CI server doesn't.
*   **Solution:** Use `vm` to boot a fresh Alpine or Debian image, mount your source code, build the package, and return the artifact. If it builds in the VM, it will build anywhere.

```arc
vm(alpine_config) func() {
    // This environment is guaranteed to be empty every time
    system("apk add gcc make")
    system("make build")
}()
```

### 2. Kernel Driver Debugging
Develop operating system drivers or kernel modules without crashing your computer.
*   **Problem:** If you write a buggy driver and load it, your host machine Blue Screens (BSOD) or Kernel Panics.
*   **Solution:** Load the driver inside an Arc VM. If it crashes, only the VM window closes. Your IDE and host OS remain perfectly stable. You can edit, re-compile, and re-boot the VM in milliseconds.

### 3. Network Stack Testing
Test complex networking code with isolation.
*   **Problem:** Testing firewall rules or custom TCP/IP stacks on your main interface is risky.
*   **Solution:** Boot a VM with a virtual network interface (`virtio-net`). You can flood the VM with traffic, poison ARP caches, or test custom protocols without disrupting your local Wi-Fi.

---

## Data Model & Shared Memory

### Value Capturing (Serialization)
Standard variables are copied by value across the boundary.

```arc
let config_data = load_json("config.json")

vm(cfg) func() {
    // config_data is serialized and injected into the VM's memory
    print(config_data["version"]) 
}()
```

### High-Performance Shared Memory
For heavy workloads (like compiling large C++ projects or processing video streams), Arc supports **Zero-Copy Shared Memory**.

```arc
// Allocate 2GB of physical RAM visible to both Host and Guest
let build_cache = vm.alloc_shared(2 * 1024 * 1024 * 1024) 

const box = vm.config {
    ram: "4GB",
    shared_memory: build_cache 
}

func main() {
    vm(box) func(cache: []byte) {
        // VM writes build artifacts directly to host RAM
        // No network transfer overhead.
        compiler.write_output(to: cache)
    }(build_cache)
}
```

---

## Licensing & Compliance

This architecture ensures your project remains easy to distribute.

1.  **No GPL Linking:** Arc interacts with hypervisors via system headers (`ioctl`, `Virtualization.framework`, `COM`). It does not statically link against GPL hypervisors like QEMU.
2.  **System Native:** Because it uses the OS's built-in virtualization (Hyper-V, KVM, HVF), users do not need to install VirtualBox or VMware to run your Arc programs.