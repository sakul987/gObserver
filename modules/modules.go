package modules

type Module interface {
	Register() error
	ProvideData() ModuleData
}

type KeyValue struct{
	Key any
	Value any
}

type ModuleData struct{
	Module string
	Data []KeyValue
}
