package easypostdevtools

import "github.com/EasyPost/easypost-go"

type Carriers struct {
	client *easypost.Client
	tools  *EasyPostDevTools
}

func (c *Carriers) Get(amount int) []string {
	_carriers := JsonReader{}.GetRandomItemsFromJsonFile(CARRIERS_JSON, amount, false)
	carriers := make([]string, len(_carriers))
	for _, carrier := range _carriers {
		carriers = append(carriers, carrier.(string))
	}
	return carriers
}
