package modules

type Module interface {
	Register() error
	ProvideData() ModuleData
}

type ModuleData struct{
	Module string
	Data []KeyValue
}

type KeyValue struct{
	Key any
	Value any
}