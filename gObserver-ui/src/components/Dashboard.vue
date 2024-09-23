<script setup lang="ts">
import { ref, onMounted } from 'vue'


const ws = ref<WebSocket>()
const wsConnected = ref<boolean>(false)

const cpu_usage = ref(0)
const cpu_temp = ref(0)
const ssd_temp = ref(0)
const ssd_size = ref(0)
const ssd_used = ref(0)
const ssd_used_perc = ref(0)
const ssd_available = ref(0)
const ram_size = ref(0)
const ram_used = ref(0)
const ram_used_perc = ref(0)
const ram_available = ref(0)

onMounted(() => {
    cpu_usage.value = 10;
    connectWS();
})

// funcs
const connectWS = () => {
    ws.value = new WebSocket("wss://localhost:3001/ws")
    
    ws.value.onopen = () => {
        wsConnected.value = true;
        console.log('Connected to WebSocket');
    };

    ws.value.onmessage = (event) => {
        console.log('Message received:', event.data);
    };

    ws.value.onclose = () => {
        wsConnected.value = false;
        console.log('WebSocket connection closed');
    };

    ws.value.onerror = (error) => {
        console.error('WebSocket error:', error);
    };
}
</script>

<template>
    <div class="grid grid-cols-1 gap-36 sm:grid-cols-2 md:grid-cols-3">
        <div>
            <h1>CPU</h1>
            <div class="grid grid-cols-2 gap-6">
                <div>
                    <h2>Usage</h2>
                    <div>{{cpu_usage}}%</div>
                </div>
                <div>
                    <h2>Temp</h2>
                    <div>{{cpu_temp}}°C</div>
                </div>
            </div>
        </div>
        <div>
            <h1>RAM</h1>
            <div class="grid grid-cols-2 gap-6">
                <div>
                    <h2>Usage</h2>
                    <div>{{ram_used}} GiB</div>
                </div>
                <div>
                    <h2>Usage %</h2>
                    <div>{{ram_used_perc}}%</div>
                </div>
                <div>
                    <h2>Size</h2>
                    <div>{{ram_size}} GiB</div>
                </div>
                <div>
                    <h2>Available</h2>
                    <div>{{ram_available}} GiB</div>
                </div>
            </div>
        </div>
        <div>
            <h1>SSD</h1>
            <div class="grid grid-cols-2 gap-6">
                <div>
                    <h2>Temp</h2>
                    <div>{{ssd_temp}}°C</div>
                </div>
                <div>
                    <h2>Usage %</h2>
                    <div>{{ssd_used_perc}}%</div>
                </div>
                <div>
                    <h2>Size</h2>
                    <div>{{ssd_size}} GiB</div>
                </div>
                <div>
                    <h2>Available</h2>
                    <div>{{ssd_available}} GiB</div>
                </div>
            </div>
        </div>
    </div>
</template>
