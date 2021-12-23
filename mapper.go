package easypostdevtools

import (
	"encoding/json"
	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
)

type Mapper struct {
}

func (m Mapper) JsonToMap(data []byte) map[string]interface{} {
	var _map map[string]interface{}
	err := json.Unmarshal(data, &_map)
	if err != nil {
		panic("Error unmarshalling json")
	}
	return _map
}

func (m Mapper) MapToJson(data map[string]interface{}) []byte {
	_json, err := json.Marshal(data)
	if err != nil {
		panic("Error marshalling json")
	}
	return _json
}

func (m Mapper) ObjectToMap(obj interface{}) map[string]interface{} {
	return structs.Map(obj)
}

func (m Mapper) MapToObject(data map[string]interface{}, obj interface{}) {
	err := mapstructure.Decode(data, obj)
	if err != nil {
		panic("Error decoding map")
	}
}

func (m Mapper) ObjectToJson(obj interface{}) []byte {
	_json, b := json.Marshal(obj)
	if b != nil {
		panic("Error marshalling object to json")
	}
	return _json
}

func (m Mapper) JsonToObject(data []byte, obj interface{}) {
	err := json.Unmarshal(data, &obj)
	if err != nil {
		panic("Error unmarshalling json to object")
	}
}

func (m Mapper) GetMapsFromJsonFile(filePath string, count int, allowDuplicates bool) []map[string]interface{} {
	return JsonReader{}.GetRandomMapsFromJsonFile(filePath, count, allowDuplicates)
}

func (m Mapper) GetMapFromJsonFile(filePath string) map[string]interface{} {
	_maps := m.GetMapsFromJsonFile(filePath, 1, false)
	return _maps[0]
}
