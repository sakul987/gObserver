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

const hdd_size = ref(0);
const hdd_used = ref(0);
const hdd_used_perc = ref(0);
const hdd_available = ref(0);

const uptime = ref("0");

const data_interval_ms = ref(0); //TODO


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
    const data = JSON.parse(eventData);
    
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
        ssd_available.value  = -1;
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
    
    // HDD    
    let hdd_size_bytes = findValueByKey(data, "HDD Size")?? -1;
    if (hdd_size_bytes != -1){
        hdd_size.value = parseFloat((hdd_size_bytes / 1024 / 1024 / 1024).toFixed(1));
    } else {
        hdd_size.value  = -1;
    }
    
    let hdd_used_bytes = findValueByKey(data, "Used HDD")?? -1;
    if (hdd_size_bytes != -1 && hdd_used_bytes != -1){
        hdd_used.value = parseFloat((hdd_used_bytes / 1024 / 1024 / 1024).toFixed(1));
        hdd_used_perc.value = parseFloat(((hdd_used_bytes / hdd_size_bytes) * 100).toFixed(1));
    }
    
    let hdd_available_byte = findValueByKey(data, "Available HDD")?? -1;
    if (hdd_available_byte != -1){
        hdd_available.value = parseFloat((hdd_available_byte / 1024 / 1024 / 1024).toFixed(1));
    } else {
        hdd_available.value  = -1;
    }
    
    // Uptime
    let uptime_s = Math.round(findValueByKey(data, "Uptime")?? -1);
    uptime.value = formatSeconds(uptime_s);
    
    data_interval_ms.value = findValueByKey(data, "Data interval")?? -1;
    data_interval_ms.value = data_interval_ms.value; //TODO use var
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

const formatSeconds = (seconds: number): string =>{
    let years = Math.floor(seconds / (365 * 24 * 60 * 60));
    seconds %= (365 * 24 * 60 * 60);
    
    let months = Math.floor(seconds / (30 * 24 * 60 * 60));
    seconds %= (30 * 24 * 60 * 60);
    
    let weeks = Math.floor(seconds / (7 * 24 * 60 * 60));
    seconds %= (7 * 24 * 60 * 60);
    
    let days = Math.floor(seconds / (24 * 60 * 60));
    seconds %= (24 * 60 * 60);
    
    let hours = Math.floor(seconds / (60 * 60));
    seconds %= (60 * 60);
    
    let minutes = Math.floor(seconds / 60);
    seconds %= 60;
    
    let timeComponents = [
        { label: 'y', value: years },
        { label: 'm', value: months },
        { label: 'w', value: weeks },
        { label: 'd', value: days },
        { label: 'h', value: hours },
        { label: 'min', value: minutes },
        { label: 's', value: seconds },
      ];
      
    let formattedTime = timeComponents
        .filter(item => item.value > 0)
        .slice(0, 3)
        .map(item => `${item.value} ${item.label}`)
        .join(', ');
        
    return formattedTime || '-1 s';
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
            <h1>HDD</h1>
            <div class="grid grid-cols-2 gap-6">
                <div>
                    <h2>Usage</h2>
                    <div>{{hdd_used}} GiB</div>
                </div>
                <div>
                    <h2>Usage %</h2>
                    <div>{{hdd_used_perc}}%</div>
                </div>
                <div>
                    <h2>Size</h2>
                    <div>{{hdd_size}} GiB</div>
                </div>
                <div>
                    <h2>Available</h2>
                    <div>{{hdd_available}} GiB</div>
                </div>
            </div>
        </div>
        <div>
            <h1>Uptime</h1>
            <div class="col-span-2">{{uptime}}</div>
        </div>
        <div>
            <h1>Status</h1>
            <div :class="connectionStateColor">{{connectionState}}</div>
        </div>
    </div>
</template>
