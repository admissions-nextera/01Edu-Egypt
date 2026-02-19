# System Monitor Project Guide

## 📋 Project Overview
Build a Desktop System Monitor application in C++ that displays real-time system information including CPU usage, memory, processes, network activity, thermal data, and fan status. You'll work with an existing codebase using the Dear ImGui library for the graphical interface, learning to read system information from Linux's `/proc` filesystem and present it in an interactive, visually appealing way.

**Key Difference**: Unlike previous projects, you're **modifying and extending existing code**, not building from scratch. This teaches you how to work with legacy codebases - a crucial real-world skill.

---

## 🎯 Learning Objectives

By completing this project, you will learn:
1. **C++ Programming**: Syntax, STL containers, file I/O, strings
2. **Immediate Mode GUI**: Dear ImGui paradigm and API
3. **Linux System Internals**: `/proc` filesystem, system monitoring
4. **Data Parsing**: Reading and parsing system files
5. **Real-time Updates**: Continuously updating displays
6. **Data Visualization**: Graphs, progress bars, tables
7. **UI/UX Design**: Creating intuitive interfaces
8. **Code Integration**: Working with existing codebases
9. **System Resources**: CPU, RAM, SWAP, Disk, Network concepts
10. **Performance Monitoring**: Tracking and displaying metrics

---

## 📚 Prerequisites - Topics You Must Know

### 1. **C++ Basics** (New Language!)
Since you're coming from Go, here are the C++ equivalents:

**Variables and Types**:
```cpp
// C++ (strongly typed like Go)
int x = 10;
float y = 3.14;
std::string name = "Hello";
bool flag = true;

// Go equivalent
var x int = 10
var y float64 = 3.14
var name string = "Hello"
var flag bool = true
```

**Functions**:
```cpp
// C++
int add(int a, int b) {
    return a + b;
}

// Go
func add(a int, b int) int {
    return a + b
}
```

**Containers** (like Go slices/maps):
```cpp
#include <vector>   // Dynamic array (like Go slice)
#include <map>      // Key-value pairs (like Go map)
#include <string>   // String type

std::vector<int> numbers = {1, 2, 3};
std::map<std::string, int> ages = {{"Alice", 30}, {"Bob", 25}};
```

**File I/O**:
```cpp
#include <fstream>
#include <string>

// Read file
std::ifstream file("/proc/cpuinfo");
std::string line;
while (std::getline(file, line)) {
    // Process line
}
file.close();
```

### 2. **C++ STL (Standard Template Library)**
- `std::vector<T>` - Dynamic arrays
- `std::string` - String operations
- `std::map<K,V>` - Associative containers
- `std::ifstream` - File input
- `std::stringstream` - String parsing

### 3. **Dear ImGui Basics**
- Immediate mode GUI concept
- Creating windows
- Rendering widgets (buttons, sliders, text)
- Handling user input
- Drawing graphs and tables

### 4. **Linux `/proc` Filesystem**
- What `/proc` is
- Key files: `/proc/cpuinfo`, `/proc/meminfo`, `/proc/stat`
- Process directories: `/proc/[pid]/`
- Reading and parsing proc files

### 5. **System Monitoring Concepts**
- CPU usage calculation
- Memory types: RAM, SWAP, Disk
- Process states
- Network statistics (RX/TX)
- Thermal zones and sensors

---

## 🖥️ Understanding Dear ImGui

### **What is Immediate Mode GUI?**

**Traditional (Retained Mode)**:
```
1. Create UI objects (buttons, labels)
2. Store them in memory
3. Update properties when needed
4. UI framework handles rendering
```

**Immediate Mode (ImGui)**:
```
Every frame:
1. Draw UI from scratch
2. Check for user input
3. Update application state
4. Repeat next frame
```

**ImGui Code Pattern**:
```cpp
// Every frame, you write:
ImGui::Begin("Window Title");
if (ImGui::Button("Click Me")) {
    // Button was clicked!
}
ImGui::Text("CPU: %.2f%%", cpuUsage);
ImGui::End();
```

**Key Concept**: You don't create and destroy widgets. You call the same drawing code every frame, and ImGui handles the rest.

---

## 🗂️ Understanding `/proc` Filesystem

### **What is `/proc`?**
A virtual filesystem created by the Linux kernel that provides interface to kernel data structures. It doesn't exist on disk - it's generated in real-time.

### **Important Files for This Project**

**1. System Information**:
```
/proc/version          - OS version
/proc/cpuinfo         - CPU details
/proc/stat            - CPU usage statistics
/proc/meminfo         - Memory information
/proc/diskstats       - Disk statistics
```

**2. Process Information**:
```
/proc/[pid]/stat      - Process statistics
/proc/[pid]/status    - Process status
/proc/[pid]/cmdline   - Command line
```

**3. Network Information**:
```
/proc/net/dev         - Network device statistics
```

**4. Thermal and Fan** (ThinkPad specific):
```
/proc/acpi/ibm/thermal - Temperature sensors
/proc/acpi/ibm/fan     - Fan status
```

### **How to Read Proc Files**

**Example: Reading CPU Info**
```cpp
#include <fstream>
#include <string>

std::ifstream file("/proc/cpuinfo");
std::string line;

while (std::getline(file, line)) {
    if (line.find("model name") != std::string::npos) {
        // Parse CPU name from this line
        // Format: "model name    : Intel(R) Core(TM) i7-8550U"
    }
}
```

**Key Skills**:
- Opening files with `std::ifstream`
- Reading line by line with `std::getline`
- Parsing strings to extract data
- Converting strings to numbers

---

## 🛠️ Step-by-Step Implementation Guide

### **Phase 1: Understanding the Codebase** 📖

#### Step 1: Download and Explore
```bash
# Download the base code
wget https://assets.01-edu.org/system-monitor/system-monitor.zip
unzip system-monitor.zip
cd system-monitor
```

#### Step 2: Examine File Structure
```
system-monitor/
├── main.cpp          - Main loop, window creation
├── system.cpp        - System resources (CPU, thermal, fan)
├── mem.cpp           - Memory and process information  
├── network.cpp       - Network statistics
├── header.h          - Function declarations
├── Makefile          - Build configuration
└── imgui/            - ImGui library files
```

#### Step 3: Understand the Main Loop
Open `main.cpp` and find the render loop:

```cpp
while (!done) {
    // 1. Poll events (keyboard, mouse)
    // 2. Start new ImGui frame
    // 3. Create windows and widgets
    // 4. Render ImGui
    // 5. Update display
}
```

**Key Concept**: Everything you draw happens inside this loop, 60+ times per second.

#### Step 4: Build and Run
```bash
# Install SDL2
sudo apt install libsdl2-dev

# Build
make

# Run
./monitor
```

**What to Expect**: 
- A window should open
- You'll see empty or placeholder UI
- Your job: Fill in the functionality

---

### **Phase 2: System Information** 💻

#### Step 5: Get Operating System Name
In `system.cpp`, create function:

```cpp
std::string GetOSName() {
    // Read /proc/version
    // Parse for OS name
    // Return string
}
```

**Implementation Strategy**:
```cpp
std::string GetOSName() {
    std::ifstream file("/proc/version");
    std::string line;
    
    if (std::getline(file, line)) {
        // line contains: "Linux version 5.x.x ..."
        // Extract "Linux" or parse for more details
        if (line.find("Linux") != std::string::npos) {
            return "Linux";
        }
    }
    
    return "Unknown";
}
```

**Display in ImGui**:
```cpp
// In main.cpp render loop
ImGui::Begin("System Monitor");
std::string os = GetOSName();
ImGui::Text("OS: %s", os.c_str());
ImGui::End();
```

---

#### Step 6: Get Logged User
```cpp
std::string GetCurrentUser() {
    // Option 1: Use environment variable
    // Option 2: Parse /proc/self/loginuid
    // Return username
}
```

**Implementation**:
```cpp
#include <unistd.h>  // for getlogin()

std::string GetCurrentUser() {
    char* user = getlogin();
    if (user != nullptr) {
        return std::string(user);
    }
    return "unknown";
}
```

**Alternative** (reading from environment):
```cpp
std::string GetCurrentUser() {
    const char* user = std::getenv("USER");
    if (user != nullptr) {
        return std::string(user);
    }
    return "unknown";
}
```

---

#### Step 7: Get Hostname
```cpp
std::string GetHostname() {
    // Read /proc/sys/kernel/hostname
    // Or use gethostname() function
}
```

**Implementation**:
```cpp
#include <unistd.h>

std::string GetHostname() {
    char hostname[256];
    if (gethostname(hostname, sizeof(hostname)) == 0) {
        return std::string(hostname);
    }
    return "unknown";
}
```

**Or reading from file**:
```cpp
std::string GetHostname() {
    std::ifstream file("/proc/sys/kernel/hostname");
    std::string hostname;
    std::getline(file, hostname);
    return hostname;
}
```

---

#### Step 8: Get Process Statistics
Parse `/proc/stat` for task information:

```cpp
struct TaskStats {
    int running;
    int sleeping;
    int stopped;
    int zombie;
};

TaskStats GetTaskStats() {
    // Iterate through /proc/[pid]/ directories
    // Read /proc/[pid]/stat for each process
    // Count states: R (running), S (sleeping), T (stopped), Z (zombie)
}
```

**Implementation Strategy**:
```cpp
#include <filesystem>  // C++17

TaskStats GetTaskStats() {
    TaskStats stats = {0};
    
    // Iterate through /proc directories
    for (const auto& entry : std::filesystem::directory_iterator("/proc")) {
        // Check if directory name is a number (PID)
        std::string name = entry.path().filename();
        if (std::isdigit(name[0])) {
            // Read /proc/[pid]/stat
            std::string statFile = entry.path() / "stat";
            std::ifstream file(statFile);
            
            // Parse state character (3rd field)
            // Format: pid (name) state ...
            // Example: 1234 (bash) S ...
            
            // Count based on state
        }
    }
    
    return stats;
}
```

**Parsing Process State**:
```cpp
// Read line from /proc/[pid]/stat
// Example: "1234 (bash) S 1233 1234 ..."
//                       ^ state character

std::string line;
std::getline(file, line);

// Find state (after second space, or between ')' and next space)
size_t pos = line.find(')');
if (pos != std::string::npos && pos + 2 < line.size()) {
    char state = line[pos + 2];
    
    switch(state) {
        case 'R': stats.running++; break;
        case 'S': stats.sleeping++; break;
        case 'T': stats.stopped++; break;
        case 'Z': stats.zombie++; break;
    }
}
```

---

#### Step 9: Get CPU Information
```cpp
std::string GetCPUModel() {
    // Read /proc/cpuinfo
    // Find line with "model name"
    // Extract CPU name
}
```

**Implementation**:
```cpp
std::string GetCPUModel() {
    std::ifstream file("/proc/cpuinfo");
    std::string line;
    
    while (std::getline(file, line)) {
        if (line.find("model name") != std::string::npos) {
            // Format: "model name	: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz"
            size_t colonPos = line.find(':');
            if (colonPos != std::string::npos) {
                std::string model = line.substr(colonPos + 2);
                return model;
            }
        }
    }
    
    return "Unknown CPU";
}
```

---

### **Phase 3: CPU Monitoring** 📊

#### Step 10: Calculate CPU Usage
CPU usage requires reading `/proc/stat` twice and calculating the difference:

```cpp
struct CPUStat {
    long user, nice, system, idle, iowait, irq, softirq;
};

CPUStat ReadCPUStat() {
    // Read /proc/stat
    // First line: "cpu  user nice system idle iowait irq softirq ..."
    // Parse numbers
}

float CalculateCPUUsage(CPUStat prev, CPUStat current) {
    // Calculate total time
    // Calculate idle time
    // Usage = (total - idle) / total * 100
}
```

**Implementation**:
```cpp
CPUStat ReadCPUStat() {
    CPUStat stat = {0};
    std::ifstream file("/proc/stat");
    std::string line;
    
    std::getline(file, line);
    // Line format: "cpu  74608 2520 24433 1117073 6176 4054 0 0 0 0"
    
    std::stringstream ss(line);
    std::string cpu;
    ss >> cpu;  // Skip "cpu"
    ss >> stat.user >> stat.nice >> stat.system >> stat.idle 
       >> stat.iowait >> stat.irq >> stat.softirq;
    
    return stat;
}

float CalculateCPUUsage(CPUStat prev, CPUStat current) {
    long prevIdle = prev.idle + prev.iowait;
    long currIdle = current.idle + current.iowait;
    
    long prevTotal = prev.user + prev.nice + prev.system + prevIdle 
                   + prev.irq + prev.softirq;
    long currTotal = current.user + current.nice + current.system + currIdle 
                   + current.irq + current.softirq;
    
    long totald = currTotal - prevTotal;
    long idled = currIdle - prevIdle;
    
    if (totald == 0) return 0.0f;
    
    return (float)(totald - idled) / totald * 100.0f;
}
```

**Usage Pattern**:
```cpp
// Store previous reading
static CPUStat prevStat = ReadCPUStat();

// In render loop:
CPUStat currentStat = ReadCPUStat();
float usage = CalculateCPUUsage(prevStat, currentStat);
prevStat = currentStat;
```

---

#### Step 11: Create Performance Graph
ImGui provides `ImGui::PlotLines()` for graphs:

```cpp
// Store history of values
static std::vector<float> cpuHistory(90, 0.0f);  // 90 data points

// Add new value, remove oldest
cpuHistory.erase(cpuHistory.begin());
cpuHistory.push_back(currentCPUUsage);

// Draw graph
ImGui::PlotLines("CPU", cpuHistory.data(), cpuHistory.size(), 
                 0, nullptr, 0.0f, 100.0f, ImVec2(0, 80));
```

**With Overlay Text**:
```cpp
// Draw graph
ImGui::PlotLines("##CPU", cpuHistory.data(), cpuHistory.size(), 
                 0, nullptr, 0.0f, 100.0f, ImVec2(400, 100));

// Draw overlay text
char overlay[32];
sprintf(overlay, "%.1f%%", currentCPUUsage);
ImGui::SameLine();
ImGui::Text("%s", overlay);
```

---

#### Step 12: Add Graph Controls
```cpp
// Variables to control graph
static bool animationPaused = false;
static float fps = 60.0f;
static float yScale = 100.0f;

// Checkbox to pause
ImGui::Checkbox("Pause Animation", &animationPaused);

// Slider for FPS
ImGui::SliderFloat("FPS", &fps, 1.0f, 120.0f);

// Slider for Y scale
ImGui::SliderFloat("Y Scale", &yScale, 50.0f, 200.0f);

// Only update if not paused
if (!animationPaused) {
    cpuHistory.erase(cpuHistory.begin());
    cpuHistory.push_back(currentCPUUsage);
}

// Use yScale in graph
ImGui::PlotLines("CPU", cpuHistory.data(), cpuHistory.size(), 
                 0, nullptr, 0.0f, yScale, ImVec2(0, 80));
```

**FPS Control**:
```cpp
// Calculate delay based on FPS
float delay = 1000.0f / fps;  // milliseconds
static auto lastUpdate = std::chrono::steady_clock::now();

auto now = std::chrono::steady_clock::now();
auto elapsed = std::chrono::duration_cast<std::chrono::milliseconds>(
    now - lastUpdate).count();

if (elapsed >= delay) {
    // Update graph
    lastUpdate = now;
}
```

---

### **Phase 4: Thermal and Fan Monitoring** 🌡️

#### Step 13: Read Thermal Sensors
```cpp
float GetCPUTemperature() {
    // Read /proc/acpi/ibm/thermal
    // OR /sys/class/thermal/thermal_zone0/temp
    // Parse temperature value
}
```

**Implementation (ThinkPad)**:
```cpp
float GetCPUTemperature() {
    std::ifstream file("/proc/acpi/ibm/thermal");
    std::string line;
    
    if (std::getline(file, line)) {
        // Format: "temperatures:	42 42 36 38 36 0 0 0 0 0 0 0 0 0 0 0"
        std::stringstream ss(line);
        std::string label;
        int temp;
        
        ss >> label;  // Skip "temperatures:"
        ss >> temp;   // First temperature
        
        return (float)temp;
    }
    
    return 0.0f;
}
```

**Alternative (Generic Linux)**:
```cpp
float GetCPUTemperature() {
    std::ifstream file("/sys/class/thermal/thermal_zone0/temp");
    int temp;
    
    if (file >> temp) {
        // Value is in millidegrees
        return temp / 1000.0f;
    }
    
    return 0.0f;
}
```

---

#### Step 14: Read Fan Information
```cpp
struct FanInfo {
    bool enabled;
    int speed;      // RPM
    int level;      // 0-7
};

FanInfo GetFanInfo() {
    // Read /proc/acpi/ibm/fan
    // Parse status, speed, level
}
```

**Implementation**:
```cpp
FanInfo GetFanInfo() {
    FanInfo info = {false, 0, 0};
    std::ifstream file("/proc/acpi/ibm/fan");
    std::string line;
    
    while (std::getline(file, line)) {
        if (line.find("status:") != std::string::npos) {
            info.enabled = (line.find("enabled") != std::string::npos);
        }
        else if (line.find("speed:") != std::string::npos) {
            std::stringstream ss(line);
            std::string label;
            ss >> label >> info.speed;  // "speed: 3200"
        }
        else if (line.find("level:") != std::string::npos) {
            std::stringstream ss(line);
            std::string label;
            ss >> label >> info.level;  // "level: 4"
        }
    }
    
    return info;
}
```

---

#### Step 15: Create Tabbed Interface
```cpp
// Create tab bar
if (ImGui::BeginTabBar("SystemTabs")) {
    
    // CPU Tab
    if (ImGui::BeginTabItem("CPU")) {
        // CPU graph and controls
        ImGui::Text("CPU Usage: %.1f%%", cpuUsage);
        // ... graph code ...
        ImGui::EndTabItem();
    }
    
    // Fan Tab
    if (ImGui::BeginTabItem("Fan")) {
        FanInfo fan = GetFanInfo();
        ImGui::Text("Status: %s", fan.enabled ? "Enabled" : "Disabled");
        ImGui::Text("Speed: %d RPM", fan.speed);
        ImGui::Text("Level: %d", fan.level);
        // ... graph code ...
        ImGui::EndTabItem();
    }
    
    // Thermal Tab
    if (ImGui::BeginTabItem("Thermal")) {
        float temp = GetCPUTemperature();
        ImGui::Text("Temperature: %.1f°C", temp);
        // ... graph code ...
        ImGui::EndTabItem();
    }
    
    ImGui::EndTabBar();
}
```

---

### **Phase 5: Memory Monitoring** 💾

#### Step 16: Read Memory Information
```cpp
struct MemoryInfo {
    long totalRAM;      // KB
    long freeRAM;
    long totalSwap;
    long freeSwap;
    long totalDisk;
    long usedDisk;
};

MemoryInfo GetMemoryInfo() {
    // Read /proc/meminfo for RAM and SWAP
    // Use df command or read /proc/mounts for disk
}
```

**Implementation**:
```cpp
MemoryInfo GetMemoryInfo() {
    MemoryInfo info = {0};
    std::ifstream file("/proc/meminfo");
    std::string line;
    
    while (std::getline(file, line)) {
        std::stringstream ss(line);
        std::string key;
        long value;
        std::string unit;
        
        ss >> key >> value >> unit;
        
        if (key == "MemTotal:") {
            info.totalRAM = value;
        }
        else if (key == "MemFree:") {
            info.freeRAM = value;
        }
        else if (key == "SwapTotal:") {
            info.totalSwap = value;
        }
        else if (key == "SwapFree:") {
            info.freeSwap = value;
        }
    }
    
    return info;
}
```

**Disk Information** (requires parsing `/proc/mounts` or using `statfs`):
```cpp
#include <sys/statvfs.h>

void GetDiskInfo(MemoryInfo& info) {
    struct statvfs stat;
    
    if (statvfs("/", &stat) == 0) {
        info.totalDisk = stat.f_blocks * stat.f_frsize / 1024;  // KB
        info.usedDisk = (stat.f_blocks - stat.f_bfree) * stat.f_frsize / 1024;
    }
}
```

---

#### Step 17: Display Memory with Progress Bars
```cpp
MemoryInfo mem = GetMemoryInfo();

// Calculate usage percentages
float ramUsage = (float)(mem.totalRAM - mem.freeRAM) / mem.totalRAM;
float swapUsage = mem.totalSwap > 0 ? 
    (float)(mem.totalSwap - mem.freeSwap) / mem.totalSwap : 0.0f;
float diskUsage = (float)mem.usedDisk / mem.totalDisk;

// Display with progress bars
ImGui::Text("RAM");
ImGui::ProgressBar(ramUsage, ImVec2(-1, 0), 
    "%.1f GB / %.1f GB", 
    (mem.totalRAM - mem.freeRAM) / 1024.0 / 1024.0,
    mem.totalRAM / 1024.0 / 1024.0);

ImGui::Text("SWAP");
ImGui::ProgressBar(swapUsage, ImVec2(-1, 0), 
    "%.1f GB / %.1f GB",
    (mem.totalSwap - mem.freeSwap) / 1024.0 / 1024.0,
    mem.totalSwap / 1024.0 / 1024.0);

ImGui::Text("Disk");
ImGui::ProgressBar(diskUsage, ImVec2(-1, 0),
    "%.1f GB / %.1f GB",
    mem.usedDisk / 1024.0 / 1024.0,
    mem.totalDisk / 1024.0 / 1024.0);
```

---

### **Phase 6: Process Table** 📋

#### Step 18: Read Process Information
```cpp
struct ProcessInfo {
    int pid;
    std::string name;
    char state;
    float cpuUsage;
    float memUsage;
};

std::vector<ProcessInfo> GetProcessList() {
    // Iterate /proc/[pid]/ directories
    // For each process:
    //   - Read PID from directory name
    //   - Read name from /proc/[pid]/comm
    //   - Read state from /proc/[pid]/stat
    //   - Calculate CPU usage
    //   - Calculate memory usage
}
```

**Implementation**:
```cpp
std::vector<ProcessInfo> GetProcessList() {
    std::vector<ProcessInfo> processes;
    
    for (const auto& entry : std::filesystem::directory_iterator("/proc")) {
        std::string name = entry.path().filename();
        
        // Check if directory name is a number
        if (std::isdigit(name[0])) {
            ProcessInfo proc;
            proc.pid = std::stoi(name);
            
            // Read process name
            std::ifstream commFile(entry.path() / "comm");
            std::getline(commFile, proc.name);
            
            // Read process state and stats
            std::ifstream statFile(entry.path() / "stat");
            std::string line;
            std::getline(statFile, line);
            
            // Parse stat line
            // Format: pid (name) state utime stime ...
            size_t pos = line.find(')');
            if (pos != std::string::npos) {
                proc.state = line[pos + 2];
                
                // Parse CPU and memory usage
                // (More complex - see next step)
            }
            
            processes.push_back(proc);
        }
    }
    
    return processes;
}
```

---

#### Step 19: Calculate Process CPU and Memory Usage
```cpp
float CalculateProcessCPU(int pid) {
    // Read /proc/[pid]/stat
    // Get utime and stime
    // Calculate percentage based on total CPU time
    // This requires tracking previous values
}

float CalculateProcessMemory(int pid) {
    // Read /proc/[pid]/status
    // Find VmRSS (Resident Set Size)
    // Calculate percentage of total RAM
}
```

**Memory Usage**:
```cpp
float CalculateProcessMemory(int pid, long totalRAM) {
    std::string path = "/proc/" + std::to_string(pid) + "/status";
    std::ifstream file(path);
    std::string line;
    
    while (std::getline(file, line)) {
        if (line.find("VmRSS:") != std::string::npos) {
            std::stringstream ss(line);
            std::string label;
            long rss;
            
            ss >> label >> rss;  // Value in KB
            
            return (float)rss / totalRAM * 100.0f;
        }
    }
    
    return 0.0f;
}
```

---

#### Step 20: Display Process Table
```cpp
// Create table
if (ImGui::BeginTable("Processes", 5, 
    ImGuiTableFlags_Borders | ImGuiTableFlags_RowBg)) {
    
    // Setup columns
    ImGui::TableSetupColumn("PID");
    ImGui::TableSetupColumn("Name");
    ImGui::TableSetupColumn("State");
    ImGui::TableSetupColumn("CPU %");
    ImGui::TableSetupColumn("Memory %");
    ImGui::TableHeadersRow();
    
    // Display each process
    std::vector<ProcessInfo> processes = GetProcessList();
    for (const auto& proc : processes) {
        ImGui::TableNextRow();
        
        ImGui::TableNextColumn();
        ImGui::Text("%d", proc.pid);
        
        ImGui::TableNextColumn();
        ImGui::Text("%s", proc.name.c_str());
        
        ImGui::TableNextColumn();
        ImGui::Text("%c", proc.state);
        
        ImGui::TableNextColumn();
        ImGui::Text("%.2f", proc.cpuUsage);
        
        ImGui::TableNextColumn();
        ImGui::Text("%.2f", proc.memUsage);
    }
    
    ImGui::EndTable();
}
```

---

#### Step 21: Add Process Filtering
```cpp
// Text input for filter
static char filter[64] = "";
ImGui::InputText("Filter", filter, sizeof(filter));

// Filter processes
std::vector<ProcessInfo> processes = GetProcessList();
std::vector<ProcessInfo> filtered;

for (const auto& proc : processes) {
    // Check if process name contains filter text
    if (strlen(filter) == 0 || 
        proc.name.find(filter) != std::string::npos) {
        filtered.push_back(proc);
    }
}

// Display filtered processes
// ... table code using 'filtered' instead of 'processes' ...
```

---

#### Step 22: Add Multi-Selection
```cpp
// Store selected PIDs
static std::set<int> selectedPIDs;

// In table, make rows selectable
for (const auto& proc : filtered) {
    ImGui::TableNextRow();
    
    // Check if row is selected
    bool isSelected = selectedPIDs.count(proc.pid) > 0;
    
    ImGui::TableNextColumn();
    // Make PID column selectable
    if (ImGui::Selectable(std::to_string(proc.pid).c_str(), 
                          isSelected, 
                          ImGuiSelectableFlags_SpanAllColumns)) {
        // Toggle selection
        if (isSelected) {
            selectedPIDs.erase(proc.pid);
        } else {
            selectedPIDs.insert(proc.pid);
        }
    }
    
    // Display other columns
    ImGui::TableNextColumn();
    ImGui::Text("%s", proc.name.c_str());
    // ... etc ...
}
```

---

### **Phase 7: Network Monitoring** 🌐

#### Step 23: Read Network Statistics
```cpp
struct NetworkStats {
    std::string interface;  // eth0, wlan0, etc.
    long rxBytes, rxPackets, rxErrs, rxDrop;
    long txBytes, txPackets, txErrs, txDrop;
};

std::vector<NetworkStats> GetNetworkStats() {
    // Read /proc/net/dev
    // Parse each line for interface statistics
}
```

**Implementation**:
```cpp
std::vector<NetworkStats> GetNetworkStats() {
    std::vector<NetworkStats> stats;
    std::ifstream file("/proc/net/dev");
    std::string line;
    
    // Skip first two header lines
    std::getline(file, line);
    std::getline(file, line);
    
    while (std::getline(file, line)) {
        NetworkStats net;
        std::stringstream ss(line);
        
        // Format: "  eth0: 1234567 8901 2 3 4 5 6 7 8901234 5678 9 0 1 2 3 4"
        std::string iface;
        ss >> iface;
        
        // Remove trailing colon
        net.interface = iface.substr(0, iface.length() - 1);
        
        // Read RX stats
        ss >> net.rxBytes >> net.rxPackets >> net.rxErrs >> net.rxDrop;
        
        // Skip 4 fields (fifo, frame, compressed, multicast)
        long skip;
        for (int i = 0; i < 4; i++) ss >> skip;
        
        // Read TX stats
        ss >> net.txBytes >> net.txPackets >> net.txErrs >> net.txDrop;
        
        stats.push_back(net);
    }
    
    return stats;
}
```

---

#### Step 24: Display Network IP Addresses
```cpp
std::string GetIPAddress(const std::string& interface) {
    // Use getifaddrs() system call
    // Or parse output of 'ip addr' command
    // Or read from /sys/class/net/[interface]/address
}
```

**Implementation using getifaddrs**:
```cpp
#include <ifaddrs.h>
#include <arpa/inet.h>

std::string GetIPAddress(const std::string& interface) {
    struct ifaddrs* ifaddr;
    
    if (getifaddrs(&ifaddr) == -1) {
        return "N/A";
    }
    
    for (struct ifaddrs* ifa = ifaddr; ifa != nullptr; ifa = ifa->ifa_next) {
        if (ifa->ifa_addr == nullptr) continue;
        
        // Check if this is the interface we want
        if (std::string(ifa->ifa_name) == interface &&
            ifa->ifa_addr->sa_family == AF_INET) {
            
            char ip[INET_ADDRSTRLEN];
            inet_ntop(AF_INET, 
                     &((struct sockaddr_in*)ifa->ifa_addr)->sin_addr,
                     ip, INET_ADDRSTRLEN);
            
            freeifaddrs(ifaddr);
            return std::string(ip);
        }
    }
    
    freeifaddrs(ifaddr);
    return "N/A";
}
```

---

#### Step 25: Create Network Tables
```cpp
// Tab bar for RX and TX
if (ImGui::BeginTabBar("NetworkTabs")) {
    
    // RX Tab
    if (ImGui::BeginTabItem("RX")) {
        if (ImGui::BeginTable("RX_Table", 8)) {
            ImGui::TableSetupColumn("Interface");
            ImGui::TableSetupColumn("Bytes");
            ImGui::TableSetupColumn("Packets");
            ImGui::TableSetupColumn("Errs");
            ImGui::TableSetupColumn("Drop");
            ImGui::TableSetupColumn("Fifo");
            ImGui::TableSetupColumn("Frame");
            ImGui::TableSetupColumn("Compressed");
            ImGui::TableHeadersRow();
            
            std::vector<NetworkStats> stats = GetNetworkStats();
            for (const auto& net : stats) {
                ImGui::TableNextRow();
                ImGui::TableNextColumn(); ImGui::Text("%s", net.interface.c_str());
                ImGui::TableNextColumn(); ImGui::Text("%ld", net.rxBytes);
                ImGui::TableNextColumn(); ImGui::Text("%ld", net.rxPackets);
                ImGui::TableNextColumn(); ImGui::Text("%ld", net.rxErrs);
                ImGui::TableNextColumn(); ImGui::Text("%ld", net.rxDrop);
                // ... more columns ...
            }
            
            ImGui::EndTable();
        }
        ImGui::EndTabItem();
    }
    
    // TX Tab (similar structure)
    if (ImGui::BeginTabItem("TX")) {
        // ... similar code for TX stats ...
        ImGui::EndTabItem();
    }
    
    ImGui::EndTabBar();
}
```

---

#### Step 26: Convert Bytes to Appropriate Units
```cpp
struct FormattedSize {
    float value;
    std::string unit;
};

FormattedSize FormatBytes(long bytes) {
    const long KB = 1024;
    const long MB = KB * 1024;
    const long GB = MB * 1024;
    
    if (bytes >= GB) {
        return {(float)bytes / GB, "GB"};
    } else if (bytes >= MB) {
        return {(float)bytes / MB, "MB"};
    } else if (bytes >= KB) {
        return {(float)bytes / KB, "KB"};
    } else {
        return {(float)bytes, "B"};
    }
}
```

**Rules from Subject**:
- Not too big: 442144.28 KB ❌
- Not too small: 0.42 GB ❌
- Just right: 431.78 MB ✓

**Implementation**:
```cpp
FormattedSize FormatBytes(long bytes) {
    const long KB = 1024;
    const long MB = KB * 1024;
    const long GB = MB * 1024;
    
    float value;
    std::string unit;
    
    if (bytes >= GB) {
        value = (float)bytes / GB;
        // Only use GB if value is >= 1.0
        if (value >= 1.0) {
            return {value, "GB"};
        }
    }
    
    if (bytes >= MB) {
        value = (float)bytes / MB;
        // Use MB if value is reasonable (1 - 1024)
        if (value >= 1.0) {
            return {value, "MB"};
        }
    }
    
    if (bytes >= KB) {
        value = (float)bytes / KB;
        return {value, "KB"};
    }
    
    return {(float)bytes, "B"};
}
```

---

#### Step 27: Display Network Usage with Progress Bars
```cpp
// Tab for RX/TX visual display
if (ImGui::BeginTabBar("NetworkVisual")) {
    
    if (ImGui::BeginTabItem("RX Visual")) {
        std::vector<NetworkStats> stats = GetNetworkStats();
        
        for (const auto& net : stats) {
            FormattedSize size = FormatBytes(net.rxBytes);
            
            // Progress bar from 0 to 2GB
            float progress = (float)net.rxBytes / (2.0 * 1024 * 1024 * 1024);
            if (progress > 1.0f) progress = 1.0f;
            
            ImGui::Text("%s:", net.interface.c_str());
            ImGui::ProgressBar(progress, ImVec2(-1, 0),
                "%.2f %s / 2.00 GB", size.value, size.unit.c_str());
        }
        
        ImGui::EndTabItem();
    }
    
    if (ImGui::BeginTabItem("TX Visual")) {
        // Similar for TX
        ImGui::EndTabItem();
    }
    
    ImGui::EndTabBar();
}
```

---

## 🐛 Common Issues and Solutions

### Issue 1: File Not Found
**Problem**: Cannot read /proc files
**Solution**: Check file path, ensure running on Linux, handle file errors
```cpp
std::ifstream file(path);
if (!file.is_open()) {
    // Handle error
    return default_value;
}
```

### Issue 2: Parsing Errors
**Problem**: String parsing fails
**Solution**: Use `std::stringstream`, check `std::string::npos`
```cpp
size_t pos = line.find(":");
if (pos == std::string::npos) {
    // Not found
    continue;
}
```

### Issue 3: Graph Not Updating
**Problem**: Graph appears frozen
**Solution**: Ensure you're updating data in render loop
```cpp
// Must be inside main loop!
while (!done) {
    // Update data here
    float cpu = GetCPUUsage();
    UpdateGraph(cpu);
    
    // Render UI
}
```

### Issue 4: Memory Leaks
**Problem**: Application uses too much memory
**Solution**: Limit history size, clean up old data
```cpp
// Limit vector size
if (cpuHistory.size() > 100) {
    cpuHistory.erase(cpuHistory.begin());
}
```

### Issue 5: Compilation Errors
**Problem**: C++ syntax errors
**Solution**: 
- Include necessary headers
- Use `std::` prefix for standard library
- End statements with semicolons
- Match braces and parentheses

---

## 📋 Testing Checklist

**System Information**:
- [ ] OS name displays correctly
- [ ] Username matches `who` command
- [ ] Hostname matches `hostname` command
- [ ] Task counts match `top` command
- [ ] CPU model matches `/proc/cpuinfo`

**CPU Monitoring**:
- [ ] CPU usage percentage is accurate
- [ ] Graph updates in real-time
- [ ] Pause button works
- [ ] FPS slider controls update rate
- [ ] Y scale slider works

**Thermal/Fan**:
- [ ] Temperature matches system sensors
- [ ] Fan status/speed/level display correctly
- [ ] Graphs update properly
- [ ] Controls work

**Memory**:
- [ ] RAM usage matches `free -h`
- [ ] SWAP usage matches `free -h`
- [ ] Disk usage matches `df -h /`
- [ ] Progress bars are accurate

**Process Table**:
- [ ] All columns display correctly
- [ ] Values match `top` command
- [ ] Filter works
- [ ] Multi-selection works

**Network**:
- [ ] IP addresses match `ifconfig`
- [ ] RX/TX tables match `/proc/net/dev`
- [ ] Byte conversion is appropriate
- [ ] Progress bars show usage correctly

---

## ✅ Submission Checklist

**Code Quality**:
- [ ] Code compiles without warnings
- [ ] Proper error handling
- [ ] Memory management (no leaks)
- [ ] Comments explain complex logic
- [ ] Follows C++ best practices

**Functionality**:
- [ ] All required features implemented
- [ ] UI is responsive and intuitive
- [ ] Real-time updates work
- [ ] No crashes or freezes

**Performance**:
- [ ] Application runs smoothly (60 FPS)
- [ ] Low CPU usage when idle
- [ ] Efficient file reading
- [ ] No memory leaks

---

## 📖 C++ Quick Reference

### **File I/O**
```cpp
#include <fstream>
std::ifstream file("path");
std::string line;
while (std::getline(file, line)) {
    // Process line
}
file.close();
```

### **String Operations**
```cpp
#include <string>
std::string str = "hello";
str.find("ll");              // Find substring
str.substr(0, 5);           // Get substring
std::stoi(str);             // String to int
std::to_string(123);        // Int to string
```

### **String Parsing**
```cpp
#include <sstream>
std::stringstream ss("10 20 30");
int a, b, c;
ss >> a >> b >> c;
```

### **Vectors**
```cpp
#include <vector>
std::vector<int> v = {1, 2, 3};
v.push_back(4);            // Add element
v.erase(v.begin());        // Remove first
v.size();                  // Get size
```

### **File System**
```cpp
#include <filesystem>
for (const auto& entry : std::filesystem::directory_iterator("/proc")) {
    std::string name = entry.path().filename();
}
```

---

## 🚀 Pro Tips

1. **Start Simple**: Get one feature working before moving to next
2. **Test Incrementally**: Compile and test after each change
3. **Use Debugger**: Learn gdb for debugging C++
4. **Read Proc Files**: Use `cat /proc/...` to see file formats
5. **Compare with Tools**: Use `top`, `htop`, `free` to verify
6. **Handle Errors**: Always check file operations
7. **Static Variables**: Use for persistent data across frames
8. **Vector Reserve**: Pre-allocate space for better performance
9. **String Streams**: Best for parsing formatted text
10. **ImGui Demo**: Run ImGui demo for widget examples

---

## 💡 Extension Ideas

After completing requirements:

1. **More Graphs**: Add graphs for memory, network
2. **Alerts**: Warn when CPU/Memory too high
3. **Save Data**: Export statistics to file
4. **Process Control**: Kill processes from UI
5. **Custom Themes**: Add dark/light themes
6. **Zoom**: Zoom in/out on graphs
7. **History**: Show historical data
8. **Multiple CPUs**: Per-core CPU usage
9. **GPU Monitoring**: If available
10. **System Logs**: Display system logs

---

## 📚 Learning Resources

**C++ Basics**:
- [Learn C++](https://www.learncpp.com/)
- [C++ Reference](https://en.cppreference.com/)
- [C++ STL Tutorial](https://www.geeksforgeeks.org/cpp-stl-tutorial/)

**Dear ImGui**:
- [ImGui GitHub](https://github.com/ocornut/imgui)
- [ImGui Demo](https://github.com/ocornut/imgui/blob/master/imgui_demo.cpp)
- [ImGui Wiki](https://github.com/ocornut/imgui/wiki)

**Linux /proc**:
- `man proc` (in terminal)
- [Linux /proc Guide](https://www.kernel.org/doc/html/latest/filesystems/proc.html)
- [Understanding /proc](https://tldp.org/LDP/Linux-Filesystem-Hierarchy/html/proc.html)

**System Monitoring**:
- Study `top`, `htop` source code
- [Linux System Monitoring](https://www.brendangregg.com/linuxperf.html)

---

## 🎓 Learning Phases

**Week 1**: C++ basics + ImGui basics + Simple displays
**Week 2**: CPU, Thermal, Fan monitoring
**Week 3**: Memory and Process table
**Week 4**: Network monitoring + Polish + Testing

---

**Remember**: This project teaches you how real system tools work. Take time to understand the /proc filesystem and how the OS exposes information. The skills you learn here apply to system programming, DevOps, and performance monitoring! 🖥️📊