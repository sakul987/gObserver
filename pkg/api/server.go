package api

import (
	"fmt"
	"log"
	"time"

	"github.com/sakul987/gObserver/modules"
	lmSensors "github.com/sakul987/gObserver/modules/lm-sensors"
)

func RunServer() error{
	//register modules
	usedModules := setModules()
	registerModules(usedModules)
	
	//run modules
	for {
		data := collectData(usedModules)
		fmt.Printf("Data: %+v\n", data)
		time.Sleep(1 * time.Second)
	}
	
	//serve api
	return nil
}

func setModules() []modules.Module{
	usedModules := []modules.Module{}
	
	usedModules = append(usedModules, lmSensors.LmSensorsModule{Name: "lm-sensors"})
	
	return usedModules
}

func registerModules(usedModules []modules.Module){
	for _, module := range usedModules{
		err := module.Register()
		if err != nil{
			log.Fatalf("Error while registering modules: %v",err)
		}
	}
}

func collectData(usedModules []modules.Module) []modules.ModuleData{
	modulesData := []modules.ModuleData{}
	
	for _, module := range usedModules{
		modulesData = append(modulesData, module.ProvideData())
	}
	
	return modulesData
}