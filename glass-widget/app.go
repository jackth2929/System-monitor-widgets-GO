package main

import (
	"context"
	"encoding/json"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

// App struct
type App struct {
	ctx          context.Context
	prevNetStat  []net.IOCountersStat
	lastNetCheck time.Time
	todoFile     string
}

// NewApp creates a new App application struct
func NewApp() *App {
	// Setup storage path in user's home directory
	home, _ := os.UserHomeDir()
	todoPath := filepath.Join(home, ".glass-widget-todos.json")
	return &App{
		todoFile: todoPath,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.lastNetCheck = time.Now()
	a.prevNetStat, _ = net.IOCounters(false)
}

type SystemStats struct {
	Hostname string `json:"hostname"`
	Platform string `json:"platform"`
	Username string `json:"username"`
	UserType string `json:"userType"`

	CPUUsage float64 `json:"cpuUsage"`
	CPUCores int     `json:"cpuCores"`

	RAMUsage float64 `json:"ramUsage"`
	RAMUsed  uint64  `json:"ramUsed"`  // in bytes
	RAMTotal uint64  `json:"ramTotal"` // in bytes

	DiskCUsage float64 `json:"diskCUsage"`
	DiskDUsage float64 `json:"diskDUsage"` // -1 if not exists

	NetSentSpeed float64 `json:"netSentSpeed"` // bytes per second
	NetRecvSpeed float64 `json:"netRecvSpeed"` // bytes per second
}

// GetSystemInfo returns system statistics
func (a *App) GetSystemInfo() SystemStats {
	stats := SystemStats{}

	// Host Info
	if h, err := host.Info(); err == nil {
		stats.Hostname = h.Hostname
		stats.Platform = h.Platform + " " + h.PlatformVersion
	}

	// User Info
	if u, err := user.Current(); err == nil {
		stats.Username = u.Username
		// Simple check for admin/root
		if u.Uid == "0" || runtime.GOOS == "windows" {
			// On Windows, checking admin usually requires more steps, 
			// for a widget we might just label it.
			stats.UserType = "Admin"
		} else {
			stats.UserType = "User"
		}
	}

	// CPU
	if percents, err := cpu.Percent(0, false); err == nil && len(percents) > 0 {
		stats.CPUUsage = percents[0]
	}
	stats.CPUCores, _ = cpu.Counts(true)

	// RAM
	if v, err := mem.VirtualMemory(); err == nil {
		stats.RAMUsage = v.UsedPercent
		stats.RAMUsed = v.Used
		stats.RAMTotal = v.Total
	}

	// Disk C:
	if d, err := disk.Usage("C:\\"); err == nil {
		stats.DiskCUsage = d.UsedPercent
	}
	// Disk D:
	if d, err := disk.Usage("D:\\"); err == nil {
		stats.DiskDUsage = d.UsedPercent
	} else {
		stats.DiskDUsage = -1 // Indicate not available
	}

	// Network Speed Calculation
	currentNetStat, err := net.IOCounters(false)
	if err == nil && len(currentNetStat) > 0 && len(a.prevNetStat) > 0 {
		now := time.Now()
		duration := now.Sub(a.lastNetCheck).Seconds()

		if duration > 0 {
			sentDiff := float64(currentNetStat[0].BytesSent - a.prevNetStat[0].BytesSent)
			recvDiff := float64(currentNetStat[0].BytesRecv - a.prevNetStat[0].BytesRecv)

			stats.NetSentSpeed = sentDiff / duration
			stats.NetRecvSpeed = recvDiff / duration
		}

		a.prevNetStat = currentNetStat
		a.lastNetCheck = now
	}

	return stats
}

type TodoItem struct {
	ID        int64  `json:"id"`
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
}

// GetTodos reads todos from local JSON file
func (a *App) GetTodos() []TodoItem {
	data, err := os.ReadFile(a.todoFile)
	if err != nil {
		// Return default list if file doesn't exist
		return []TodoItem{
			{ID: 1, Text: "Welcome to GlassWidget!", Completed: false},
		}
	}

	var todos []TodoItem
	err = json.Unmarshal(data, &todos)
	if err != nil {
		return []TodoItem{}
	}
	return todos
}

// SaveTodos saves todos to local JSON file
func (a *App) SaveTodos(todos []TodoItem) error {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(a.todoFile, data, 0644)
}
