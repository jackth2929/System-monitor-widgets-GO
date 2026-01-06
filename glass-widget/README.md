# GlassWidget (Light Mode Edition)

A compact, elegant system monitoring widget for Windows, built with **Go (Wails)** and **Vue 3**.  
Designed with a "Milky White" Glassmorphism aesthetic, optimized for **Light Mode** desktop environments.

## ✨ Features

*   **Real-time Monitoring**:
    *   **CPU**: Usage % and Physical Core count.
    *   **RAM**: Usage %, Used/Total capacity (e.g., 4.5 GB / 16 GB).
    *   **Disk**: Usage bars for C: and D: drives.
    *   **Network**: Real-time Upload/Download speeds (Auto-scaling units).
*   **Compact Design**: Fixed 450x380 window size, perfect for desktop corners.
*   **Glassmorphism UI**: High-quality frosted glass effect with dark text for high readability on light backgrounds.
*   **Info Header**: Displays real-time Clock, Date, Hostname, User Account, and System Platform info.
*   **Draggable**: Drag anywhere on the header to move the widget.

## 🛠️ Tech Stack

*   **Backend**: Go (Golang) 1.21+
*   **Frontend**: Vue 3 + Vite
*   **Styling**: Tailwind CSS v4
*   **Framework**: Wails v2 (Go + Webview)
*   **System Lib**: `gopsutil` for hardware statistics.

## 🚀 Prerequisites

Before you begin, ensure you have installed:

1.  **Go** (v1.21 or later): [Download](https://go.dev/dl/)
2.  **Node.js** (v18 or later): [Download](https://nodejs.org/)
3.  **Wails CLI**:
    ```powershell
    go install github.com/wailsapp/wails/v2/cmd/wails@latest
    ```

## 📦 Installation & Setup

1.  **Clone the repository**:
    ```powershell
    git clone <repository-url>
    cd glass-widget
    ```

2.  **Install Frontend Dependencies**:
    ```powershell
    cd frontend
    npm install
    cd ..
    ```

## 💻 Development

To run the application in development mode (with hot reload):

```powershell
wails dev
```

*   **Frontend-only mode**: If you just want to edit the UI without Go:
    ```powershell
    cd frontend
    npm run dev
    ```
    *(Note: This uses mock data defined in `main.js`)*

## 🔨 Build for Production

To create a standalone `.exe` file:

```powershell
wails build
```

The output file will be located at:
`build/bin/glass-widget.exe`

## 📂 Project Structure

```text
glass-widget/
├── app.go              # Backend Logic (System Stats)
├── main.go             # Entry Point & Window Config
├── wails.json          # Project Configuration (Window size, etc.)
├── frontend/
│   ├── src/
│   │   ├── App.vue     # Main UI Component
│   │   ├── style.css   # Tailwind & Glassmorphism Styles
│   │   └── wailsjs/    # Auto-generated Go bindings
│   └── package.json
└── build/              # Build artifacts
```

## 📝 License

MIT