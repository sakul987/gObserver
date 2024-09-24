package websocket

import (
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/sakul987/gObserver/modules"
	"github.com/sakul987/gObserver/pkg/config"
	"github.com/sakul987/gObserver/pkg/constants"
	"golang.org/x/net/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var mutex = &sync.RWMutex{}

func SendData(usedModules []modules.Module){
	for {
		data := collectData(usedModules)
		
		jsonData, err := json.Marshal(data)
		if err != nil{
			log.Println("Error unmarshaling modules data:", err.Error())
		}
		
		mutex.Lock()
		for client := range clients{
			if err := websocket.Message.Send(client, string(jsonData)); err != nil{
				delete(clients, client)
				client.Close()
			}
		}
		mutex.Unlock()
		
		time.Sleep(config.DATA_COLLECT_INTERVAL_MS * time.Millisecond)
	}
}

func collectData(usedModules []modules.Module) []modules.ModuleData{
	modulesData := []modules.ModuleData{}
	
	for _, module := range usedModules{
		modulesData = append(modulesData, module.ProvideData())
	}
	
	modulesData = append(modulesData, modules.ModuleData{Module: "gObserver", Data: []modules.KeyValue{{Key: constants.DATA_INTERVAL, Value: config.DATA_COLLECT_INTERVAL_MS }}})
	
	return modulesData
}

func WebsocketHandler(ws *websocket.Conn) {
	defer func(){
		mutex.Lock()
		delete(clients, ws)
		mutex.Unlock()
		ws.Close()
		log.Println("Removed a client")
	}()
	
	mutex.Lock()
	clients[ws] = true
	mutex.Unlock()
	log.Println("Added a client")
	
	waitForRemoval(ws)
}

func waitForRemoval(ws *websocket.Conn){
	for{
		time.Sleep(config.WS_CLEANUP_INTERVAL_S * time.Second)
		
		mutex.RLock()
		_, exists := clients[ws]
		mutex.RUnlock()
		
		if !exists {
			break
		}
	}
}