package lmSensors

import (
	"encoding/json"
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
	fmt.Println("Registered module",m.Name," successfully!")
	
	return nil
}

func (m LmSensorsModule) ProvideData() modules.ModuleData{
	moduleData := modules.ModuleData{Module: m.Name}
	
	data := getData()
	for _, entry := range data{
		moduleData.Data = append(moduleData.Data, entry)
	}
	
	return moduleData
}

func checkDependencies() error{
	path, err := exec.LookPath("sensors")
	
	fmt.Println("Path of lm-sensors:",path)
	
	return err
}

func getData() []modules.KeyValue{
	data := []modules.KeyValue{}
	
	cmd := exec.Command("sensors", "-j", "-A")
	
	output, err := cmd.Output()
	if err != nil{
		data = append(data, modules.KeyValue{Key: "error", Value: err.Error()})
		return data
	}
	
	parsedData, err := parseData(output)
	if err != nil{
		data = append(data, modules.KeyValue{Key: "error", Value: err.Error()})
		return data
	}
	
	for _, entry := range parsedData{
		data = append(data, entry)
	}
	
	
	
	
	
	return data
}

func parseData(data []byte) ([]modules.KeyValue, error){
	var mappedData map[string]interface{}
	err := json.Unmarshal(data, &mappedData)
	if err != nil{
		return nil, fmt.Errorf("error while mapping raw JSON data: %w", err)
	}
	
	entries := parseKeyValuePairs(mappedData)
	
	return entries, nil
}

func parseKeyValuePairs(data map[string]interface{}) []modules.KeyValue {
	var entries []modules.KeyValue
	for k, v := range data {
		if nestedMap, ok := v.(map[string]interface{}); ok {
			entries = append(entries, modules.KeyValue{
				Key:   k,
				Value: parseKeyValuePairs(nestedMap),
			})
		} else {
			entries = append(entries, modules.KeyValue{
				Key:   k,
				Value: v,
			})
		}
	}
	return entries
}

func filterData (data []modules.KeyValue) []modules.KeyValue{
	filteredData := []modules.KeyValue{}
	
	for _, entry := range data{
		currentDevice := entry
		currentDeviceFiltered := modules.KeyValue{Key: currentDevice.Key}
		
		for _, value := range currentDevice{}
		
		filteredData = append(filteredData, )
	}
	
	return filteredData
}