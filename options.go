package easypostdevtools

import "github.com/EasyPost/easypost-go"

type Options struct {
	client *easypost.Client
	tools  *EasyPostDevTools
}

func (o *Options) GetMap() map[string]interface{} {
	return Mapper{}.GetMapFromJsonFile(OPTIONS_JSON)
}
