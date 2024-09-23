package modules

type Module interface {
	Register() error
	ProvideData() ModuleData
}

type ModuleData struct{
	Module string `json:"module"`
	Data []KeyValue `json:"data"`
}

type KeyValue struct{
	Key any
	Value any
}