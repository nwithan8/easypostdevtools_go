package easypostdevtools

import "github.com/EasyPost/easypost-go"

type Labels struct {
	client *easypost.Client
	tools  *EasyPostDevTools
}

func (l *Labels) GetRandomLabelOptions() map[string]interface{} {
	return Mapper{}.GetMapFromJsonFile(LABEL_OPTIONS_JSON)
}
