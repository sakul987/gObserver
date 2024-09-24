package websocket

import (
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/sakul987/gObserver/modules"
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
		
		time.Sleep(1 * time.Second)
	}
}

func collectData(usedModules []modules.Module) []modules.ModuleData{
	modulesData := []modules.ModuleData{}
	
	for _, module := range usedModules{
		modulesData = append(modulesData, module.ProvideData())
	}
	
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
		time.Sleep(30 * time.Second)
		
		mutex.RLock()
		_, exists := clients[ws]
		mutex.RUnlock()
		
		if !exists {
			break
		}
	}
}