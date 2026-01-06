import { createApp } from 'vue'
import './style.css'
import App from './App.vue'

// Mock Wails runtime for browser development
if (!window['go']) {
    window['go'] = {
        main: {
            App: {
                GetSystemInfo: () => Promise.resolve({
                    hostname: "Dev-Host",
                    platform: "Windows 11 Pro",
                    username: "DevUser",
                    userType: "Admin",
                    cpuUsage: 45.5,
                    cpuCores: 8,
                    ramUsage: 60.2,
                    ramUsed: 1024 * 1024 * 1024 * 4.5, // 4.5 GB
                    ramTotal: 1024 * 1024 * 1024 * 16, // 16GB
                    diskCUsage: 75,
                    diskDUsage: 30,
                    netSentSpeed: 1024 * 50, // 50KB/s
                    netRecvSpeed: 1024 * 1024 * 2.5 // 2.5MB/s
                })
            }
        }
    }
}

createApp(App).mount('#app')
