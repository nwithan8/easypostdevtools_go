package easypostdevtools

import "github.com/EasyPost/easypost-go"

type CustomsInfos struct {
	client *easypost.Client
	tools  *EasyPostDevTools
}

func (c *CustomsInfos) GetMap(itemsAmount int, allowDuplicateItems bool) map[string]interface{} {
	_map := Mapper{}.GetMapFromJsonFile(CUSTOMS_INFO_JSON)
	_map["items"] = c.tools.CustomsItems.Get(itemsAmount, allowDuplicateItems)
	return _map
}

func (c *CustomsInfos) Get(itemsAmount int, allowDuplicateItems bool) easypost.CustomsInfo {
	// Create customs info locally + POST to API
	_map := c.GetMap(itemsAmount, allowDuplicateItems)
	_info := easypost.CustomsInfo{}
	Mapper{}.MapToObject(_map, &_info)
	info, err := c.client.CreateCustomsInfo(&_info)
	if err != nil {
		panic("Could not create customs info")
	}
	return *info
}
