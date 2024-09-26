package df

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/sakul987/gObserver/modules"
	"github.com/sakul987/gObserver/pkg/constants"
)

type DfModule struct {
	Name string
}

func (m DfModule) Register() error {
	if err := checkDependencies(); err != nil {
		return err
	}
	fmt.Println("Registered module", m.Name, "successfully!")

	return nil
}

func (m DfModule) ProvideData() modules.ModuleData {
	moduleData := modules.ModuleData{Module: m.Name}

	data := getData()
	for _, entry := range data {
		moduleData.Data = append(moduleData.Data, entry)
	}

	return moduleData
}

func checkDependencies() error {
	_, err := exec.LookPath("df")

	return err
}

func getData() []modules.KeyValue {
	data := []modules.KeyValue{}

	cmd := exec.Command("df", "-t", "ext4")

	output, err := cmd.Output()
	if err != nil {
		data = append(data, modules.KeyValue{Key: "error", Value: err.Error()})
		return data
	}

	resultSSD, err := parseData(strings.Split(string(output), "\n")[1])
	if err != nil {
		data = append(data, modules.KeyValue{Key: "error", Value: err.Error()})
		return data
	}

	data = append(data, modules.KeyValue{Key: constants.SSD_SIZE, Value: resultSSD[0]})
	data = append(data, modules.KeyValue{Key: constants.SSD_USED, Value: resultSSD[1]})
	data = append(data, modules.KeyValue{Key: constants.SSD_AVAILABLE, Value: resultSSD[2]})
	
	resultHDD, err := parseData(strings.Split(string(output), "\n")[2])
	if err != nil {
		data = append(data, modules.KeyValue{Key: "error", Value: err.Error()})
		return data
	}
	
	data = append(data, modules.KeyValue{Key: constants.HDD_SIZE, Value: resultHDD[0]})
	data = append(data, modules.KeyValue{Key: constants.HDD_USED, Value: resultHDD[1]})
	data = append(data, modules.KeyValue{Key: constants.HDD_AVAILABLE, Value: resultHDD[2]})

	return data
}

func parseData(line string) ([]int, error) {
	fields := strings.Fields(line)
	if len(fields) < 5 {
		return nil, fmt.Errorf("df did not provide sufficient data")
	}

	size, err := strconv.Atoi(fields[1])
	if err != nil {
		return nil, fmt.Errorf("unable to convert storage size to int")
	}

	used, err := strconv.Atoi(fields[2])
	if err != nil {
		return nil, fmt.Errorf("unable to convert used storage to int")
	}

	available, err := strconv.Atoi(fields[3])
	if err != nil {
		return nil, fmt.Errorf("unable to convert available storage to int")
	}

	return []int{size * 1024, used * 1024, available * 1024}, nil
}
