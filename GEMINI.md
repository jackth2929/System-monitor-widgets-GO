
---

# 專案規格書：Go + Wails 桌面多功能小工具 (GlassWidget)

## 1. 專案概述 (Overview)

本專案旨在開發一款基於 **Go (Backend)** 與 **Wails (Frontend)** 的跨平台桌面小工具。核心設計理念為「輕量、美觀、實用」，採用 Glassmorphism (玻璃擬態) 設計風格，提供系統效能監控功能，旨在提升使用者的桌面美感與工作效率。

## 2. 技術堆疊 (Tech Stack)

| 領域 | 技術/工具 | 說明 |
| --- | --- | --- |
| **核心框架** | **Wails v2** | 提供 Go 與 Web 前端的橋接，生成原生執行檔。 |
| **後端 (Backend)** | **Go (Golang) 1.21+** | 處理業務邏輯、系統呼叫、資料持久化。 |
| **前端 (Frontend)** | **Vue 3 + Vite** | (建議) 響應式 UI 開發，效能優異。 |
| **樣式 (Styling)** | **Tailwind CSS** | 快速實現 Glassmorphism 與動態顏色樣式。 |
| **系統監控庫** | **gopsutil** | (`github.com/shirou/gopsutil`) 獲取 CPU/RAM/Disk/Net 資訊。 |
| **資料存儲** | **JSON / BoltDB** | 本地輕量化儲存 (待辦事項、設定)。 |
| **天氣 API** | **OpenWeatherMap** | (或其他免費 API) 提供天氣數據。 |

---

## 3. 系統架構與配置 (Architecture & Configuration)

### 3.1 Wails 視窗配置 (`wails.json`)

為了達到設計要求，應用程式需設定為無邊框且背景透明：

* **Frameless:** `true` (無邊框)
* **Transparent:** `true` (允許背景透明)
* **Resizable:** `false` (固定大小，保持版面精緻)
* **AlwaysOnTop:** `true` (可選：是否常駐最上層)
* **Draggable:** 使用 CSS `-webkit-app-region: drag` 實現視窗拖曳。

### 3.2 資料流向

1. **前端 (UI):** 發送事件 (如點擊「完成任務」) 或定期輪詢 (Polling)。
2. **Wails Bridge:** 將前端請求轉發給 Go Runtime。
3. **後端 (Go):** 執行 API 請求、讀取系統狀態或寫入文件，將結果回傳前端。

---

## 4. 功能需求詳細說明 (Functional Requirements)

### 📊 系統監控模組 (System Monitor Module)

* **更新頻率:** 每 1~2 秒 (Ticker)。
* **監控指標 & 邏輯:**
* **Device Info:**: 顯示主機名稱 (Hostname)、作業系統版本 (Platform)、使用者名稱 (Username) 及權限類型 (User Type)。
* **CPU:** 總體使用率 (%)及核心數。
* **RAM:** 顯示記憶體使用百分比及具體容量 (已用 GB / 總共 GB)。
* **Disk:** 鎖定監控 `C:`、`D:` 磁碟使用率。
* **Network:** 計算每秒的 Bytes Sent/Recv，轉換為 KB/s 或 MB/s。


* **狀態顏色邏輯 (CSS Class Binding):**
* 🟢 **正常 (Normal):** < 50% (綠色)
* 🟠 **警告 (Warning):** 50% - 80% (橙色)
* 🔴 **危險 (Critical):** > 80% (紅色)


---

## 5. UI/UX 設計規範 (Design Guidelines)

### 5.1 玻璃擬態 (Glassmorphism) 實作

需使用 CSS Backdrop Filter 實現毛玻璃效果：

```css
.widget-container {
  background: rgba(255, 255, 255, 0.15); /* 低透明度白色 */
  backdrop-filter: blur(12px);           /* 背景模糊 */
  -webkit-backdrop-filter: blur(12px);
  border: 1px solid rgba(255, 255, 255, 0.3); /* 半透明邊框 */
  box-shadow: 0 4px 30px rgba(0, 0, 0, 0.1);  /* 柔和陰影 */
  border-radius: 16px;
  color: white;
}

```

### 5.2 字體與圖示

* **字體:** Inter 或 Roboto (無襯線，易讀性高)。
* **圖示:** 使用 Phosphor Icons 或 Heroicons (線條風格，適合極簡設計)。

---

## 6. 開發階段規劃 (Development Phases)

* **Phase 1: 基礎建設 (Skeleton)**
* 初始化 Wails 專案。
* 設定視窗透明度與拖曳功能。
* 建立基本 UI Grid 佈局。


* **Phase 2: 後端邏輯 (Go Backend)**
* 實作 `SystemStats` 結構體與 `GetSystemInfo()` 方法 (gopsutil)。
* 實作 `WeatherAPI` 呼叫與 JSON 解析。
* 實作 `TodoService` 的 Load/Save JSON 功能。


* **Phase 3: 前端整合 (Frontend Integration)**
* 使用 Vue/React 綁定後端方法。
* 實作動態顏色邏輯 (CPU/RAM)。
* 完成待辦事項的互動 (新增/勾選)。


* **Phase 4: 優化與打包 (Polish & Build)**
* 調整 Glassmorphism 參數。
* 加入計時器結束的動畫/音效。
* 編譯 Windows `.exe` 檔案並測試資源佔用。

---