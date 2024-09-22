package meminfo

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/sakul987/gObserver/modules"
	"github.com/sakul987/gObserver/pkg/constants"
)

type MeminfoModule struct{
	Name string
}

func (m MeminfoModule) Register() error{
	fmt.Println("Registered module",m.Name,"successfully!")
	
	return nil
}

func (m MeminfoModule) ProvideData() modules.ModuleData{
	moduleData := modules.ModuleData{Module: m.Name}
	
	data := getData()
	for _, entry := range data{
		moduleData.Data = append(moduleData.Data, entry)
	}
	
	return moduleData
}

func getData() []modules.KeyValue{
	data := []modules.KeyValue{}
	
	file, err := os.Open("/proc/meminfo")
	if err != nil{
		data = append(data, modules.KeyValue{Key: "error", Value: err.Error()})
		return data
	}
	defer file.Close()
	
	result, err := parseFile(file)
	if err != nil{
		data = append(data, modules.KeyValue{Key: "error", Value: err.Error()})
		return data
	}
	
	data = append(data, modules.KeyValue{Key:   constants.RAM_SIZE, Value: result[0]})
	data = append(data, modules.KeyValue{Key:   constants.RAM_USED, Value: result[1]})
	data = append(data, modules.KeyValue{Key:   constants.RAM_AVAILABLE, Value: result[2]})
	
	return data
}

func parseFile(file *os.File) ([]int, error){
	var size, available int
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := scanner.Text()
		if strings.HasPrefix(line, "MemTotal:"){
			fields := strings.Fields(line)
			size, err := strconv.Atoi(fields[1])
			if err != nil{
				return nil, err
			}
			size *= 1024
		} else if strings.HasPrefix(line, "MemAvailable:"){
			fields := strings.Fields(line)
			available, err := strconv.Atoi(fields[1])
			if err != nil{
				return nil, err
			}
			available *= 1024
		}
		
		if size != 0 && available != 0{
			break
		}
	}
	
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	
	used := size - available
	
	return []int{size, used, available}, nil
}
