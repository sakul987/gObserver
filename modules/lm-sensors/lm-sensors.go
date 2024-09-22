package lmSensors

import (
	"fmt"
	"os/exec"

	"github.com/sakul987/gObserver/modules"
)

type LmSensorsModule struct{
	Name string
}

func (m LmSensorsModule) Register() error{
	if err := checkDependencies(); err != nil{
		return err
	}
	
	return nil
}

func (m LmSensorsModule) ProvideData() modules.ModuleData{
	data := modules.ModuleData{Module: m.Name}
	
	return data
}

func checkDependencies() error{
	path, err := exec.LookPath("sensors")
	
	fmt.Println("Path of lm-sensors:",path)
	
	return err
}