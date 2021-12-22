package easypostdevtools

import "github.com/EasyPost/easypost-go"

type CustomsItems struct {
	client *easypost.Client
	tools  *EasyPostDevTools
}

func (c *CustomsItems) GetRandomCustomsItemsMaps(amount int, allowDuplicates bool) []map[string]interface{} {
	return Mapper{}.GetMapsFromJsonFile(CUSTOMS_ITEMS_JSON, amount, allowDuplicates)
}

func (c *CustomsItems) Get(amount int, allowDuplicates bool) []easypost.CustomsItem {
	// create customs items locally + POST to API
	customsItems := []easypost.CustomsItem{}
	for _, customsItem := range c.GetRandomCustomsItemsMaps(amount, allowDuplicates) {
		_item := easypost.CustomsItem{}
		Mapper{}.MapToObject(customsItem, &_item)
		item, err := c.client.CreateCustomsItem(&_item)
		if err != nil {
			panic("Could not create customs item")
		}
		customsItems = append(customsItems, *item)
	}
	return customsItems
}
