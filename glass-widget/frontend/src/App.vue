<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import { GetSystemInfo } from './wailsjs/go/main/App'

const stats = ref({
    hostname: 'Loading...',
    platform: '...',
    username: '...',
    userType: '',
    cpuUsage: 0,
    cpuCores: 0,
    ramUsage: 0,
    ramUsed: 0,
    ramTotal: 0,
    diskCUsage: 0,
    diskDUsage: 0,
    netSentSpeed: 0,
    netRecvSpeed: 0
})

const currentTime = ref('')
const currentDate = ref('')
const timer = ref(null)
const clockTimer = ref(null)

const updateTime = () => {
    const now = new Date()
    currentTime.value = now.toLocaleTimeString('en-US', { hour12: false, hour: '2-digit', minute: '2-digit', second: '2-digit' })
    currentDate.value = now.toLocaleDateString('zh-CN', { year: 'numeric', month: 'long', day: 'numeric' })
}

const formatBytes = (bytes, decimals = 1) => {
    if (!+bytes) return '0 B'
    const k = 1024
    const dm = decimals < 0 ? 0 : decimals
    const sizes = ['B', 'KB', 'MB', 'GB', 'TB']
    const i = Math.floor(Math.log(bytes) / Math.log(k))
    return `${parseFloat((bytes / Math.pow(k, i)).toFixed(dm))} ${sizes[i]}`
}

const formatSpeed = (bytesPerSec) => {
    return formatBytes(bytesPerSec) + '/s'
}

// Light Mode Colors: Darker shades for text visibility
const getStatusColor = (percent) => {
    if (percent < 50) return 'text-emerald-600' // Darker green
    if (percent < 80) return 'text-orange-600'  // Darker orange
    return 'text-red-600'                       // Darker red
}

const getBarColor = (percent) => {
    if (percent < 50) return 'bg-emerald-500'
    if (percent < 80) return 'bg-orange-500'
    return 'bg-red-500'
}

const fetchData = async () => {
    try {
        stats.value = await GetSystemInfo()
    } catch (e) {
        console.error(e)
    }
}

onMounted(async () => {
    await fetchData()
    updateTime()
    timer.value = setInterval(fetchData, 1500)
    clockTimer.value = setInterval(updateTime, 1000)
})

onUnmounted(() => {
    if (timer.value) clearInterval(timer.value)
    if (clockTimer.value) clearInterval(clockTimer.value)
})
</script>

<template>
  <!-- Main Container: Light Text Theme -->
  <div class="h-screen w-screen flex flex-col p-4 space-y-4 text-slate-800 select-none overflow-hidden">
    
    <!-- Header -->
    <header class="drag-region glass-panel p-3 px-4 rounded-xl flex justify-between items-center cursor-move">
        <div class="flex flex-col">
             <div class="text-2xl font-bold tracking-widest font-mono leading-none text-slate-900">{{ currentTime }}</div>
             <div class="text-sm text-slate-500 mt-1">{{ currentDate }}</div>
        </div>
        <div class="text-right">
            <div class="font-bold text-xs leading-tight text-slate-800">{{ stats.username }}</div>
            <div class="text-[9px] text-slate-500 truncate mb-1">{{ stats.hostname }}</div>
            <div class="flex items-center justify-end space-x-2">
                <span class="text-[8px] text-slate-500 uppercase">{{ stats.platform }}</span>
                <span class="text-[8px] px-1.5 py-0.5 rounded bg-slate-200/50 border border-slate-300/50 text-slate-600">{{ stats.userType }}</span>
            </div>
        </div>
    </header>

    <!-- Grid Layout -->
    <div class="grid grid-cols-2 grid-rows-2 gap-3 flex-1 min-h-0">
        
        <!-- CPU Card -->
        <div class="glass-panel p-4 rounded-xl flex flex-col justify-between relative overflow-hidden">
            <div class="flex justify-between items-start z-10">
                <span class="text-[10px] font-bold text-slate-500 uppercase">CPU</span>
                <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 text-slate-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 3v2m6-2v2M9 19v2m6-2v2M5 9H3m2 6H3m18-6h-2m2 6h-2M7 19h10a2 2 0 002-2V7a2 2 0 00-2-2H7a2 2 0 00-2 2v10a2 2 0 002 2zM9 9h6v6H9V9z" />
                </svg>
            </div>
            <div class="mt-1 z-10">
                <div class="text-4xl font-bold tracking-tighter" :class="getStatusColor(stats.cpuUsage)">
                    {{ stats.cpuUsage.toFixed(1) }}<span class="text-lg">%</span>
                </div>
                <div class="text-[11px] text-slate-600 font-semibold">{{ stats.cpuCores }} Physical Cores</div>
            </div>
            <!-- Progress Bar Track: Darker gray for visibility on light bg -->
            <div class="w-full bg-slate-200 rounded-full h-1.5 mt-2 z-10">
                <div class="h-1.5 rounded-full transition-all duration-500" :class="getBarColor(stats.cpuUsage)" :style="{ width: `${stats.cpuUsage}%` }"></div>
            </div>
        </div>

        <!-- RAM Card -->
        <div class="glass-panel p-4 rounded-xl flex flex-col justify-between relative overflow-hidden">
            <div class="flex justify-between items-start z-10">
                <span class="text-[10px] font-bold text-slate-500 uppercase">Memory</span>
                <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 text-slate-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19.428 15.428a2 2 0 00-1.022-.547l-2.387-.477a6 6 0 00-3.86.517l-.318.158a6 6 0 01-3.86.517L6.05 15.21a2 2 0 00-1.806.547M8 4h8l-1 1v5.172a2 2 0 00.586 1.414l5 5c1.26 1.26.367 3.414-1.415 3.414H4.828c-1.782 0-2.674-2.154-1.414-3.414l5-5A2 2 0 009 10.172V5L8 4z" />
                </svg>
            </div>
            <div class="mt-1 z-10">
                <div class="text-4xl font-bold tracking-tighter" :class="getStatusColor(stats.ramUsage)">
                    {{ stats.ramUsage.toFixed(1) }}<span class="text-lg">%</span>
                </div>
                <div class="text-[11px] text-slate-600 font-mono font-semibold">
                    {{ formatBytes(stats.ramUsed) }} <span class="text-slate-400">/</span> {{ formatBytes(stats.ramTotal) }}
                </div>
            </div>
            <div class="w-full bg-slate-200 rounded-full h-1.5 mt-2 z-10">
                <div class="h-1.5 rounded-full transition-all duration-500" :class="getBarColor(stats.ramUsage)" :style="{ width: `${stats.ramUsage}%` }"></div>
            </div>
        </div>

        <!-- Disk Card -->
        <div class="glass-panel p-4 rounded-xl flex flex-col justify-center space-y-3">
             <div class="flex items-center space-x-2">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-3 w-3 text-slate-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 7v10c0 2.21 3.582 4 8 4s8-1.79 8-4V7M4 7c0 2.21 3.582 4 8 4s8-1.79 8-4M4 7c0-2.21 3.582-4 8-4s8 1.79 8 4m0 5c0 2.21-3.582 4-8 4s-8-1.79-8-4" />
                </svg>
                <span class="text-[10px] font-bold text-slate-500 uppercase">Storage</span>
            </div>
            <div class="space-y-2">
                <div>
                    <div class="flex justify-between text-[11px] mb-1 font-semibold text-slate-700">
                        <span>Local (C:)</span>
                        <span :class="getStatusColor(stats.diskCUsage)">{{ stats.diskCUsage.toFixed(0) }}%</span>
                    </div>
                    <div class="w-full bg-slate-200 rounded-full h-1.5 border border-slate-200/50">
                        <div class="h-1.5 rounded-full transition-all duration-500" :class="getBarColor(stats.diskCUsage)" :style="{ width: `${stats.diskCUsage}%` }"></div>
                    </div>
                </div>
                <div v-if="stats.diskDUsage >= 0">
                    <div class="flex justify-between text-[11px] mb-1 font-semibold text-slate-700">
                        <span>Data (D:)</span>
                        <span :class="getStatusColor(stats.diskDUsage)">{{ stats.diskDUsage.toFixed(0) }}%</span>
                    </div>
                    <div class="w-full bg-slate-200 rounded-full h-1.5 border border-slate-200/50">
                        <div class="h-1.5 rounded-full transition-all duration-500" :class="getBarColor(stats.diskDUsage)" :style="{ width: `${stats.diskDUsage}%` }"></div>
                    </div>
                </div>
            </div>
        </div>

        <!-- Network Card -->
        <div class="glass-panel p-4 rounded-xl flex flex-col justify-between">
            <div class="flex items-center justify-between mb-2">
                <span class="text-[10px] font-bold text-slate-500 uppercase">Network</span>
                <svg xmlns="http://www.w3.org/2000/svg" class="h-3 w-3 text-slate-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.684 13.342C8.886 12.938 9 12.482 9 12c0-.482-.114-.938-.316-1.342m0 2.684a3 3 0 110-2.684m0 2.684l6.632 3.316m-6.632-6l6.632-3.316m0 0a3 3 0 105.367-2.684 3 3 0 00-5.367 2.684zm0 9.316a3 3 0 105.368 2.684 3 3 0 00-5.368-2.684z" />
                </svg>
            </div>
            <div class="grid grid-cols-1 gap-2">
                <div class="flex flex-col">
                    <div class="flex items-center space-x-1 mb-0.5">
                         <svg class="w-2 h-2 text-slate-400" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M5 10l7-7m0 0l7 7m-7-7v18" /></svg>
                         <span class="text-[8px] text-slate-500 font-bold tracking-wider">UPLOAD</span>
                    </div>
                    <span class="text-sm font-mono font-medium leading-none text-slate-700">{{ formatSpeed(stats.netSentSpeed) }}</span>
                </div>
                <div class="flex flex-col">
                    <div class="flex items-center space-x-1 mb-0.5">
                         <svg class="w-2 h-2 text-emerald-500" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M19 14l-7 7m0 0l-7-7m7 7V3" /></svg>
                         <span class="text-[8px] text-slate-500 font-bold tracking-wider uppercase">Download</span>
                    </div>
                    <span class="text-sm font-mono font-medium text-emerald-600 leading-none">{{ formatSpeed(stats.netRecvSpeed) }}</span>
                </div>
            </div>
        </div>

    </div>
  </div>
</template>