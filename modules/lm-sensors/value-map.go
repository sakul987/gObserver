package lmSensors

import (
	"github.com/sakul987/gObserver/modules"
	"github.com/sakul987/gObserver/pkg/constants"
)

var mapping []modules.KeyValue = []modules.KeyValue{
	{
		Key: "coretemp-isa-0000.Package id 0.temp1_input",
		Value: constants.CPU_TEMP,
	},
	{
		Key: "nvme-pci-0100.Composite.temp1_input",
		Value: constants.SSD_TEMP,
	},
}

func mapToValues(rawData []modules.KeyValue) []modules.KeyValue{
	data := []modules.KeyValue{}
	
	for _, entry := range rawData{
		for _, mapEntry := range mapping{
			if entry.Key == mapEntry.Key {
				data = append(data, modules.KeyValue{
					Key:   mapEntry.Value,
					Value: entry.Value,
				})
			}
		}
	}
	
	
	return data
}