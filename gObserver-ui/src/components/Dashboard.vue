<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'


const ws = ref<WebSocket>()
const wsConnected = ref<boolean>(false)

const cpu_usage = ref(0);
const cpu_temp = ref(0);

const ssd_temp = ref(0);
const ssd_size = ref(0);
const ssd_used_perc = ref(0);
const ssd_available = ref(0);

const ram_size = ref(0);
const ram_used = ref(0);
const ram_used_perc = ref(0);
const ram_available = ref(0);

const uptime = ref(0);


onMounted(() => {
    connectWS();
})

//computed
const connectionState = computed(():string =>{
    return wsConnected.value ? "Connected" : "Not connected"
});
const connectionStateColor = computed(():string =>{
    return wsConnected.value ? "text-green-500" : "text-red-500"
});

// funcs
const connectWS = () => {    
    ws.value = new WebSocket("wss://localhost:3001/ws")
    
    ws.value.onopen = () => {
        wsConnected.value = true;
        console.log('Connected to WebSocket');
    };

    ws.value.onmessage = (event) => {
        handleMessage(event.data);
    };

    ws.value.onclose = () => {
        wsConnected.value = false;
        
        console.warn("Websocket connection closed. Retrying in 10s.");
        
        setTimeout(connectWS, 10000);
    };

    ws.value.onerror = () => {
        wsConnected.value = false;
    };
}

const handleMessage = (eventData: any) =>{
    //console.log(eventData);
    const data = JSON.parse(eventData);
    //console.log(data)
    
    // CPU
    cpu_usage.value = findValueByKey(data, "CPU Usage")?? -1;
    cpu_temp.value = Math.round(findValueByKey(data, "CPU Temperature")?? -1);
    
    // SSD
    ssd_temp.value = Math.round(findValueByKey(data, "SSD Temperature")?? -1);
    
    let ssd_size_bytes = findValueByKey(data, "Storage Size")?? -1;
    if (ssd_size_bytes != -1){
        ssd_size.value = parseFloat((ssd_size_bytes / 1024 / 1024 / 1024).toFixed(1));
    } else {
        ssd_size.value  = -1;
    }
    
    let ssd_used = findValueByKey(data, "Used Storage")?? -1;
    if (ssd_size_bytes != -1 && ssd_used != -1){
        ssd_used_perc.value = parseFloat(((ssd_used / ssd_size_bytes) * 100).toFixed(1));
    }
    
    let ssd_available_byte = findValueByKey(data, "Available Storage")?? -1;
    if (ssd_available_byte != -1){
        ssd_available.value = parseFloat((ssd_available_byte / 1024 / 1024 / 1024).toFixed(1));
    } else {
        ssd_available_byte.value  = -1;
    }

    
    // RAM
    let ram_size_bytes = findValueByKey(data, "RAM Size")?? -1;
    if (ram_size_bytes != -1){
        ram_size.value = parseFloat((ram_size_bytes / 1024 / 1024 / 1024).toFixed(1));
    } else {
        ram_size.value  = -1;
    }
    
    let ram_used_bytes = findValueByKey(data, "Used RAM")?? -1;
    if (ram_used_bytes != -1){
        ram_used.value = parseFloat((ram_used_bytes / 1024 / 1024 / 1024).toFixed(1));
    } else {
        ram_used.value  = -1;
    }
    
    if (ram_size_bytes != -1 && ram_used_bytes != -1){
        ram_used_perc.value = parseFloat(((ram_used_bytes / ram_size_bytes) * 100).toFixed(1));
    }
    
    let ram_available_bytes = findValueByKey(data, "Available RAM")?? -1;
    if (ram_available_bytes != -1){
        ram_available.value = parseFloat((ram_available_bytes / 1024 / 1024 / 1024).toFixed(1));
    } else {
        ram_available_bytes.value  = -1;
    }
    
    let uptime_s = Math.round(findValueByKey(data, "Uptime")?? -1);
    uptime.value = uptime_s;
    
    let data_interval_ms = findValueByKey(data, "Data interval")?? -1;
    console.log("Data interval: " + data_interval_ms);
}

const findValueByKey = (data: any, targetKey: string): any => {
    let result = null;

    if (Array.isArray(data)) {
        for (const item of data) {
            result = findValueByKey(item, targetKey);
            if (result !== null) {
                break;
            }
        }
    } else if (typeof data === 'object' && data !== null) {
        for (const key in data) {
            if (key === "Key" && data[key] === targetKey) {
                return data["Value"];
            }

            result = findValueByKey(data[key], targetKey);
            if (result !== null) {
                break;
            }
        }
    }

    return result;
};
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
        <div>
            <h1>Status</h1>
            <div class="grid grid-cols-2 gap-6">
                <div class="col-span-2" :class="connectionStateColor">{{connectionState}}</div>
                <div class="col-span-2">Server uptime:</div>
                <div class="col-span-2">{{uptime}}</div>
            </div>
        </div>
    </div>
</template>
