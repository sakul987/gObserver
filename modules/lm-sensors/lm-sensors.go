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
	fmt.Println("Registered module",m.Name,"successfully!")
	
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
	_, err := exec.LookPath("sensors")
	
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
	
	var mappedData map[string]interface{}
	err = json.Unmarshal(output, &mappedData)
	if err != nil {
		data = append(data, modules.KeyValue{Key: "error", Value: err.Error()})
		return data
	}
	
	var entries []modules.KeyValue
	parseKeyValuePairs(mappedData, "", &entries)

	result := mapToValues(entries)
	fmt.Println(result)
	
	return data
}

func parseKeyValuePairs(data map[string]interface{}, parentKey string, entries *[]modules.KeyValue) {
	for k, v := range data {
		fullKey := k
		if parentKey != "" {
			fullKey = parentKey + "." + k
		}

		switch value := v.(type) {
		case map[string]interface{}:
			// If the value is another map, recurse into it
			parseKeyValuePairs(value, fullKey, entries)
		default:
			// Otherwise, just add the key-value pair to the entries
			*entries = append(*entries, modules.KeyValue{
				Key:   fullKey,
				Value: v,
			})
		}
	}
}
