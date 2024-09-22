package cpuUsage

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/sakul987/gObserver/modules"
	"github.com/sakul987/gObserver/pkg/constants"
)

type CpuUsageModule struct{
	Name string
}

var previousTotal = 0
var previousIdle = 0

func (m CpuUsageModule) Register() error{
	fmt.Println("Registered module",m.Name,"successfully!")
	
	return nil
}

func (m CpuUsageModule) ProvideData() modules.ModuleData{
	moduleData := modules.ModuleData{Module: m.Name}
	
	data := getData()
	for _, entry := range data{
		moduleData.Data = append(moduleData.Data, entry)
	}
	
	return moduleData
}

func getData() []modules.KeyValue{
	data := []modules.KeyValue{}
	
	fields, err := getFields()
	if err != nil{
		data = append(data, modules.KeyValue{Key: "error", Value: err.Error()})
		return data
	}
	
	currentIdle, user, nice, system, err := extractValues(fields)
	if err != nil{
		data = append(data, modules.KeyValue{Key: "error", Value: err.Error()})
		return data
	}
	
	currentTotal := user + nice + system + currentIdle
	
	idleDiff := float64(currentIdle - previousIdle)
	totalDiff := float64(currentTotal - previousTotal)
	usage := 100*((totalDiff - idleDiff) / totalDiff)
	roundedUsage := math.Round(usage * 10) / 10
	
	previousIdle = currentIdle
	previousTotal = currentTotal
	
	data = append(data, modules.KeyValue{Key:   constants.CPU_USAGE, Value: roundedUsage})
	
	return data
}

func getFields() ([]string, error){
	data, err := os.ReadFile("/proc/stat")
	if err != nil {
		return nil, err
	}
	
	lines := strings.Split(string(data), "\n")
	if len(lines) < 1{
		return nil, err
	}
	
	line := lines[0]
	
	fields := strings.Fields(line)
	if len(fields) < 5{ 
		return nil, err
	}
	
	return fields, nil
}

func extractValues(fields []string) (int, int, int, int, error){
	currentIdle, err := strconv.Atoi(fields[4])
	if err != nil {
		return 0,0,0,0, err
	}
	
	user, err := strconv.Atoi(fields[1])
	if err != nil {
		return 0,0,0,0, err
	}
	
	nice, err := strconv.Atoi(fields[2])
	if err != nil {
		return 0,0,0,0, err
	}
	
	system, err := strconv.Atoi(fields[3])
	if err != nil {
		return 0,0,0,0, err
	}
	
	return currentIdle, user, nice, system, nil
}