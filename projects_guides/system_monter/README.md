# System Monitor Project Guide

> **Before you start:** Download and extract the provided file system. Run `make` and confirm the application compiles and opens a window before writing a single line. You are adding features to existing code — understand what is already there before touching it.

---

## Objectives

By completing this project you will learn:

1. **C++ Programming** — Reading and modifying an existing C++ codebase
2. **Immediate Mode GUI** — How Dear ImGui renders frames and how its widget API works
3. **The `/proc` Filesystem** — Reading CPU, memory, process, and network data from Linux's virtual filesystem
4. **System Resources** — What CPU usage, RAM, SWAP, disk, and network metrics actually mean
5. **Real-Time Data Display** — Updating graphs and tables on every frame
6. **File Parsing in C++** — Reading and parsing text files from `/proc`

---

## Prerequisites — Topics You Must Know Before Starting

### 1. C++ Basics
- How to read a file in C++ (`std::ifstream`, `getline`)
- Structs and classes
- Vectors, maps, strings
- Search: **"C++ read file line by line"**
- Search: **"C++ struct example"**

### 2. Dear ImGui Paradigm
- What "immediate mode" means — every widget is drawn every frame
- How the main loop works: `NewFrame → widgets → Render`
- Basic widgets: `ImGui::Text`, `ImGui::PlotLines`, `ImGui::SliderFloat`, `ImGui::Checkbox`, `ImGui::BeginTabBar`
- Read: https://github.com/ocornut/imgui/wiki#about-the-imgui-paradigm
- Search: **"Dear ImGui immediate mode explained"**

### 3. Linux `/proc` Filesystem
- Run `man proc` in your terminal — read the sections on `/proc/stat`, `/proc/meminfo`, `/proc/net/dev`
- Explore `/proc` manually: `cat /proc/stat`, `cat /proc/meminfo`, `cat /proc/net/dev`
- Search: **"linux /proc filesystem explained"**

### 4. Build System
- How to run `make` and read a Makefile
- Install SDL2: `apt install libsdl2-dev`

**Do these before writing any code:**
- Run `make` — confirm the window opens
- Run `cat /proc/stat` — read what the fields mean in `man proc`
- Run `cat /proc/meminfo` — identify which fields are RAM and SWAP

---

## Project Structure

```
system-monitor/
├── header.h          ← shared structs and declarations — add yours here
├── main.cpp          ← main loop, window setup, calls your render functions
├── system.cpp        ← CPU, fan, thermal data — your System tab goes here
├── mem.cpp           ← RAM, SWAP, disk, processes — your Memory tab goes here
├── network.cpp       ← network interfaces — your Network tab goes here
├── Makefile
└── imgui/lib/        ← Dear ImGui source — do not modify
```

---

## Milestone 1 — Understand the Existing Code

**Goal:** You can explain what every existing function does and where in `main.cpp` each tab would be rendered.

**Questions to answer before writing anything:**
- Where in `main.cpp` is the main render loop? What three ImGui calls surround every frame?
- What files currently exist for the three tabs? What functions are already declared in `header.h`?
- How does Dear ImGui create a window? What is `ImGui::Begin` / `ImGui::End`?
- What does `ImGui::PlotLines` need as input — what type and format of data does it expect?
- How do you create a tabbed section with `ImGui::BeginTabBar` and `ImGui::BeginTabItem`?

**Resources:**
- Open `imgui/lib/imgui_demo.cpp` — this is the full demo of every widget. Search for `PlotLines`, `TabBar`, `SliderFloat` to see usage examples.
- https://github.com/ocornut/imgui/blob/master/imgui.h — all function signatures

**Verify:** Read `main.cpp` top to bottom. Draw on paper where each of the three monitors will be rendered.

---

## Milestone 2 — System Information (system.cpp)

**Goal:** The System tab displays OS type, logged-in user, hostname, and process counts.

**Questions to answer:**
- Which `/proc` file contains the OS information? Which file in `/etc`?
- How do you get the currently logged-in user in C++ without a system call?
- Which file contains the hostname?
- Which `/proc` file lists all processes and their states? How do you count running vs sleeping vs zombie?

**Code Placeholder:**
```cpp
// system.cpp

struct SystemInfo {
    // OS name
    // Username
    // Hostname
    // Process counts: running, sleeping, stopped, zombie
    // CPU model name
};

SystemInfo getSystemInfo() {
    // Read OS info from /etc/os-release or /proc/version
    // Read username from environment or /etc/passwd
    // Read hostname from /proc/sys/kernel/hostname
    // Count processes by reading state from /proc/[pid]/stat for each PID in /proc
    // Read CPU model from /proc/cpuinfo
}

void renderSystemInfo(const SystemInfo& info) {
    // ImGui::Text for OS, user, hostname
    // ImGui::Text for each process state count
    // ImGui::Text for CPU model
}
```

**Resources:**
- `man proc` → search for `/proc/[pid]/stat` for process state field
- Search: **"C++ read directory entries opendir readdir"**
- Search: **"C++ getenv USERNAME"**

**Verify:** The rendered window shows your actual OS, username, hostname, and process counts. Cross-check process count with `ps aux | wc -l`.

---

## Milestone 3 — CPU Usage Graph (system.cpp)

**Goal:** The CPU tab shows a real-time performance graph, an overlay with the current CPU percentage, an FPS slider, a Y-scale slider, and a pause checkbox.

**Questions to answer:**
- `/proc/stat` gives cumulative CPU ticks. How do you compute CPU usage percentage from two successive reads?
- What fields in `/proc/stat` line 1 are: user, nice, system, idle, iowait? Which ones count as "busy"?
- `ImGui::PlotLines` takes an array of floats. How do you maintain a rolling history of N samples?
- How do you draw text on top of a `PlotLines` graph (the overlay parameter)?
- How do the FPS and Y-scale sliders affect the graph — what exactly do they control?
- How does the pause checkbox stop the data from updating while the graph stays visible?

**Code Placeholder:**
```cpp
// system.cpp

struct CpuState {
    // Previous tick values (for delta calculation)
    // Rolling history of usage percentages (circular buffer)
    // Pause flag
    // FPS setting (float)
    // Y scale setting (float)
};

float getCpuUsage(CpuState& state) {
    // Read /proc/stat first line
    // Parse: user, nice, system, idle, iowait, irq, softirq
    // Compute: total = sum of all fields
    // Compute: idle_delta = idle - prev_idle
    // Compute: total_delta = total - prev_total
    // Compute: usage = (1.0 - idle_delta / total_delta) * 100.0
    // Store current values as previous for next call
    // Return usage percentage
}

void renderCpuTab(CpuState& state) {
    // If not paused: sample getCpuUsage, append to history
    // ImGui::Checkbox for pause
    // ImGui::SliderFloat for FPS
    // ImGui::SliderFloat for Y scale
    // ImGui::PlotLines with overlay text showing current %
}
```

**Resources:**
- `man proc` → search `/proc/stat`
- Search: **"calculate cpu usage from /proc/stat"**
- Search: **"ImGui PlotLines overlay text"**

**Verify:** The graph animates in real time. Pause stops the animation. The overlay shows the current percentage. Cross-check with `top` or `htop`.

---

## Milestone 4 — Fan and Thermal (system.cpp)

**Goal:** Fan tab shows status, speed, and level with a graph. Thermal tab shows temperature with a graph.

**Questions to answer:**
- Where on a Linux system is fan information stored? (Hint: not `/proc` — search `/sys/class/hwmon`)
- Which file under `/sys/class/hwmon/hwmon*/` gives fan speed in RPM?
- Where is CPU temperature stored on Linux?
- What does "fan enabled" vs "fan active" mean — are they different files?
- How do you reuse the same graph component across CPU, Fan, and Thermal tabs?

**Code Placeholder:**
```cpp
// system.cpp

struct FanInfo {
    // Speed in RPM
    // Level
    // Is enabled (bool)
    // Is active (bool)
    // Rolling history
};

struct ThermalInfo {
    // Current temperature (millidegrees → degrees conversion needed)
    // Rolling history
};

FanInfo getFanInfo() {
    // Read /sys/class/hwmon/hwmon*/fan1_input for speed
    // Read enable and active files
    // Read pwm level if available
}

float getCpuTemp() {
    // Find the correct hwmon device for CPU temperature
    // Read temp1_input (value is in millidegrees Celsius)
    // Convert to Celsius
}

void renderFanTab(FanInfo& fan) {
    // Display status, speed, level as text
    // Show graph with same controls as CPU tab
}

void renderThermalTab(ThermalInfo& thermal) {
    // Show graph with temperature overlay text
}
```

**Resources:**
- Search: **"linux /sys/class/hwmon fan speed temperature"**
- Search: **"linux read cpu temperature from sysfs"**
- `ls /sys/class/hwmon/` then `cat` the files to identify which has fan/temp data

**Verify:** Fan speed matches what `sensors` command reports. Temperature matches `sensors` output.

---

## Milestone 5 — Memory, SWAP, and Disk (mem.cpp)

**Goal:** The Memory tab shows RAM, SWAP, and disk usage each with a visual progress bar display.

**Questions to answer:**
- Which fields in `/proc/meminfo` give you: total RAM, available RAM, total SWAP, free SWAP?
- How do you calculate "used" from "total" and "available"?
- Where does Linux expose disk usage information? (`/proc/mounts` + `statvfs` syscall)
- What does `ImGui::ProgressBar` expect as input — what range of values?
- How do you display the usage as both a bar and a text label like `3.2 GB / 8.0 GB`?

**Code Placeholder:**
```cpp
// mem.cpp

struct MemoryInfo {
    // Total RAM (bytes)
    // Available RAM (bytes)
    // Total SWAP (bytes)
    // Free SWAP (bytes)
    // Disk total (bytes)
    // Disk used (bytes)
};

MemoryInfo getMemoryInfo() {
    // Read /proc/meminfo line by line
    // Extract MemTotal, MemAvailable, SwapTotal, SwapFree
    // Use statvfs() on "/" for disk info
}

void renderMemoryTab(const MemoryInfo& info) {
    // For RAM: ImGui::ProgressBar(used/total) + text showing used/total in GB
    // For SWAP: same pattern
    // For Disk: same pattern
}
```

**Resources:**
- `man statvfs` — read the struct fields
- Search: **"C++ statvfs disk usage example"**
- Search: **"ImGui ProgressBar usage"**

**Verify:** RAM used matches `free -h`. Disk usage matches `df -h /`.

---

## Milestone 6 — Process Table (mem.cpp)

**Goal:** A table displays all running processes with PID, name, state, CPU%, and memory%. A text filter narrows the list. Multiple rows can be selected.

**Questions to answer:**
- How do you enumerate all processes? (Read numeric directories under `/proc/`)
- Which file under `/proc/[pid]/` gives the process name? Which gives state and CPU ticks?
- How do you calculate CPU usage per process? (Same delta approach as total CPU, but per-process.)
- How do you calculate memory usage per process as a percentage of total RAM?
- How does `ImGui::InputText` filter a table — what happens on each frame when the filter changes?
- How does multi-row selection work with `ImGui::Selectable` inside a table?

**Code Placeholder:**
```cpp
// mem.cpp

struct Process {
    // PID
    // Name
    // State (char)
    // CPU usage (%)
    // Memory usage (%)
};

std::vector<Process> getProcessList() {
    // Open /proc and read all numeric directory entries
    // For each PID:
    //   Read /proc/[pid]/stat for name, state, utime, stime
    //   Read /proc/[pid]/statm for memory pages used
    //   Compute CPU% using delta from previous read
    //   Compute memory% = (pages * page_size) / total_ram * 100
    // Return sorted list
}

void renderProcessTable(std::vector<Process>& processes) {
    // ImGui::InputText for filter string
    // ImGui::BeginTable with columns: PID, Name, State, CPU%, Mem%
    // For each process:
    //   If name or PID contains filter string:
    //     ImGui::TableNextRow
    //     ImGui::Selectable for multi-select support
    //     ImGui::TableSetColumnIndex for each column
}
```

**Resources:**
- `man proc` → search `/proc/[pid]/stat` and `/proc/[pid]/statm`
- Search: **"Dear ImGui table selectable multi-select"**
- Search: **"C++ list directory entries /proc"**
- Open `imgui_demo.cpp` and search for `Tables` — read the table examples

**Verify:** Table shows all processes. Typing in the filter narrows results. Clicking multiple rows selects them.

---

## Milestone 7 — Network Information (network.cpp)

**Goal:** The Network tab shows IPv4 for each interface and two tables (RX and TX) with all required columns. Progress bars show per-interface usage with smart byte unit conversion.

**Questions to answer:**
- Which file contains all network interface statistics? What is its format?
- How do you extract the IPv4 address for each interface in C++?
- `/proc/net/dev` has two header lines before the data — how do you skip them?
- What is the rule for choosing GB, MB, or KB for display? (The spec gives an example: 452MB is right, 0.42GB is too small, 442144KB is too big.)
- How do you display a progress bar that goes from 0 to 2GB regardless of the unit used for the label?
- What are the 8 RX columns and 8 TX columns in `/proc/net/dev`?

**Code Placeholder:**
```cpp
// network.cpp

struct NetworkInterface {
    // Interface name (lo, wlp5s0, eth0, ...)
    // IPv4 address string
    // RX: bytes, packets, errs, drop, fifo, frame, compressed, multicast
    // TX: bytes, packets, errs, drop, fifo, colls, carrier, compressed
};

std::vector<NetworkInterface> getNetworkInfo() {
    // Read /proc/net/dev line by line, skip 2 header lines
    // For each interface line: parse name and all 16 values
    // For each interface: get IPv4 from getifaddrs() or ioctl(SIOCGIFADDR)
}

std::string formatBytes(long long bytes) {
    // If bytes >= 1GB threshold: format as GB with 2 decimal places
    // If bytes >= 1MB threshold: format as MB with 2 decimal places
    // Otherwise: format as KB with 2 decimal places
    // Return the formatted string with unit
}

void renderNetworkTab(std::vector<NetworkInterface>& interfaces) {
    // Display IPv4 for each interface

    // Tab bar with RX and TX tabs
    // Each tab: ImGui::BeginTable with correct columns
    // For each interface: one table row

    // Second tab bar with RX and TX usage bars
    // For each interface:
    //   ImGui::ProgressBar(bytes / 2GB_in_bytes)
    //   Label showing formatBytes(bytes) result
}
```

**Resources:**
- `man getifaddrs` — for getting IPv4 addresses
- Search: **"C++ getifaddrs IPv4 address example"**
- Search: **"linux /proc/net/dev format fields"**
- `cat /proc/net/dev` — read what you have on your machine

**Verify:** Interface names and IPv4 addresses match `ip addr show`. RX/TX values match `ifconfig` or `ip -s link`. Unit conversion: test with a value you know (e.g. `cat /proc/net/dev` after a large download).

---

## Debugging Checklist

- Does `make` fail? Check that `libsdl2-dev` is installed with `dpkg -l libsdl2-dev`. Check the Makefile for the correct include paths.
- Does the window open but show nothing new? Your render functions are written but not called from `main.cpp`. Find where to add your calls.
- Is CPU usage always 0% or 100%? You are reading `/proc/stat` once, not twice with a delta. You need two reads separated by time.
- Are process CPU percentages all 0? Same issue — you need to store previous tick values between frames.
- Does fan/temperature show 0 or fail? Run `ls /sys/class/hwmon/hwmon*/` and `cat` each file to find which hwmon device has your fan and CPU temperature.
- Are network bytes not matching `ifconfig`? Check that you are reading the right columns — `/proc/net/dev` has 8 RX columns then 8 TX columns, and the column order matters.
- Does the app crash when a process disappears between reading the directory and reading its `/proc/[pid]/stat`? Handle `ENOENT` — processes can die mid-read.

---

## Key Libraries and Files

| Item | What You Use It For | Reference |
|---|---|---|
| `imgui.h` | All ImGui widget functions | `imgui/lib/imgui.h` |
| `imgui_demo.cpp` | Examples of every widget | `imgui/lib/imgui_demo.cpp` |
| `/proc/stat` | CPU tick counts | `man proc` |
| `/proc/meminfo` | RAM and SWAP totals/usage | `man proc` |
| `/proc/[pid]/stat` | Per-process state and CPU ticks | `man proc` |
| `/proc/[pid]/statm` | Per-process memory pages | `man proc` |
| `/proc/net/dev` | Network interface RX/TX stats | `man proc` |
| `/sys/class/hwmon/` | Fan speed and CPU temperature | `man sysfs` |
| `statvfs()` | Disk usage | `man statvfs` |
| `getifaddrs()` | IPv4 address per interface | `man getifaddrs` |

---

## Submission Checklist

- [ ] Application compiles with `make` without errors or warnings
- [ ] **System tab**: OS, user, hostname, process counts, CPU model all displayed
- [ ] **CPU tab**: real-time graph with overlay %, FPS slider, Y-scale slider, pause checkbox
- [ ] **Fan tab**: status (enabled/active), speed, level, and real-time graph
- [ ] **Thermal tab**: real-time temperature graph with overlay text
- [ ] **Memory tab**: RAM usage with progress bar and GB label
- [ ] **SWAP tab**: SWAP usage with progress bar and GB label
- [ ] **Disk tab**: disk usage with progress bar
- [ ] **Process table**: PID, Name, State, CPU%, Mem% columns
- [ ] **Process filter**: text input narrows table in real time
- [ ] **Multi-row selection**: multiple processes can be selected simultaneously
- [ ] **Network IPv4**: shown for lo, wlp5s0, and any other interfaces
- [ ] **Network RX table**: all 8 columns correct
- [ ] **Network TX table**: all 8 columns correct
- [ ] **Network usage bars**: 0–2GB scale, correct unit conversion (GB/MB/KB)
- [ ] Application never crashes during normal use
- [ ] Handles missing `/proc/[pid]` entries gracefully (processes that exit mid-read)