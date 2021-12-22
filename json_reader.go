package easypostdevtools

import (
	"encoding/json"
	"io/ioutil"
)

type JsonReader struct {
}

func (jr *JsonReader) ReadFile(path string) []byte {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic("Could not read file")
	}
	return data
}

func (jr *JsonReader) ReadJsonFileJson(path string) map[string]interface{} {
	data := jr.ReadFile(path)
	var result map[string]interface{}
	err := json.Unmarshal(data, &result)
	if err != nil {
		panic("Could not parse json file")
	}
	return result
}

func (jr *JsonReader) ReadJsonFileArray(path string) []map[string]interface{} {
	data := jr.ReadFile(path)
	var result []map[string]interface{}
	err := json.Unmarshal(data, &result)
	if err != nil {
		panic("Could not parse json file")
	}
	return result
}

func (jr *JsonReader) GetRandomMapsFromJsonFile(path string, amount int, allow_duplicates bool) []map[string]interface{} {
	data := jr.ReadJsonFileArray(path)
	// convert data to list of interface
	var result []interface{}
	for _, v := range data {
		result = append(result, v)
	}
	random := Random{}
	items := random.GetRandomItemsFromList(result, amount, allow_duplicates)
	// convert items to list of map
	var resultMaps []map[string]interface{}
	for _, v := range items {
		resultMaps = append(resultMaps, v.(map[string]interface{}))
	}
	return resultMaps
}

func (jr *JsonReader) GetRandomItemsFromJsonFile(path string, amount int, allow_duplicates bool) []interface{} {
	data := jr.ReadJsonFileArray(path)
	// convert data to list of interface
	var result []interface{}
	for _, v := range data {
		result = append(result, v)
	}
	random := Random{}
	return random.GetRandomItemsFromList(result, amount, allow_duplicates)
}
