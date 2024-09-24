package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sakul987/gObserver/modules"
	cpuUsage "github.com/sakul987/gObserver/modules/cpu-usage"
	"github.com/sakul987/gObserver/modules/df"
	lmSensors "github.com/sakul987/gObserver/modules/lm-sensors"
	"github.com/sakul987/gObserver/modules/meminfo"
	"github.com/sakul987/gObserver/pkg/config"
	intWebsocket "github.com/sakul987/gObserver/pkg/websocket"
	"golang.org/x/net/websocket"
)

func RunServer() error{
	//register modules
	usedModules := setModules()
	registerModules(usedModules)
	
	go intWebsocket.SendData(usedModules)
	
	//serve api
	// handler only to accept self signed certificate by calling it once (https://localhost:3001)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){w.WriteHeader(http.StatusOK)})
	http.Handle("/ws", websocket.Handler(intWebsocket.WebsocketHandler))
	log.Println("Starting server on :"+ config.BACKEND_PORT + "\n")
	return http.ListenAndServeTLS(":"+ config.BACKEND_PORT + "", config.CERTIFICATE_CRT, config.CERTIFICATE_KEY, nil)
}

func setModules() []modules.Module{
	usedModules := []modules.Module{}
	
	usedModules = append(usedModules, lmSensors.LmSensorsModule{Name: "lm-sensors"})
	usedModules = append(usedModules, df.DfModule{Name: "df"})
	usedModules = append(usedModules, meminfo.MeminfoModule{Name: "meminfo"})
	usedModules = append(usedModules, cpuUsage.CpuUsageModule{Name: "cpu-usage"})
	
	return usedModules
}

func registerModules(usedModules []modules.Module){
	fmt.Printf("\n------------- Registering modules -------------\n\n")
	for _, module := range usedModules{
		err := module.Register()
		if err != nil{
			log.Fatalf("Error while registering modules: %v",err)
		}
	}
	fmt.Printf("\n-------- Finished registering modules --------\n\n")
}