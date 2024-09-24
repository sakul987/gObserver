package uptime

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/sakul987/gObserver/modules"
	"github.com/sakul987/gObserver/pkg/config"
	"github.com/sakul987/gObserver/pkg/constants"
)

type UptimeModule struct{
	Name string
}

func (m UptimeModule) Register() error{
	fmt.Println("Registered module",m.Name,"successfully!")
	
	return nil
}

func (m UptimeModule) ProvideData() modules.ModuleData{
	moduleData := modules.ModuleData{Module: m.Name}
	
	moduleData.Data = append(moduleData.Data, getData())
	
	return moduleData
}

func getData() modules.KeyValue{	
	file, err := os.ReadFile(config.MODULES_UPTIME_SOURCE)
	if err != nil{
		return modules.KeyValue{Key: "error", Value: err.Error()}
	}
	
	uptimeStr := strings.Fields(string(file))[0]
	
	uptime, err := strconv.ParseFloat(uptimeStr, 64)
	if err != nil{
		return modules.KeyValue{Key: "error", Value: err.Error()}
	}
	
	return modules.KeyValue{Key: constants.UPTIME, Value: uptime}
}