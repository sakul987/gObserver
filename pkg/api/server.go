package api

import (
	"fmt"
	"log"
	"time"

	"github.com/sakul987/gObserver/modules"
	cpuUsage "github.com/sakul987/gObserver/modules/cpu-usage"
	"github.com/sakul987/gObserver/modules/df"
	lmSensors "github.com/sakul987/gObserver/modules/lm-sensors"
	"github.com/sakul987/gObserver/modules/meminfo"
)

func RunServer() error{
	//register modules
	usedModules := setModules()
	registerModules(usedModules)
	
	//run modules
	for {
		data := collectData(usedModules)
		fmt.Printf("Data: %v\n\n", data)
		time.Sleep(1 * time.Second)
	}
	
	//serve api
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

func collectData(usedModules []modules.Module) []modules.ModuleData{
	modulesData := []modules.ModuleData{}
	
	for _, module := range usedModules{
		modulesData = append(modulesData, module.ProvideData())
	}
	
	return modulesData
}