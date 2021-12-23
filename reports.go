package easypostdevtools

import "github.com/EasyPost/easypost-go"

type Reports struct {
	client *easypost.Client
	tools  *EasyPostDevTools
}

func (r *Reports) GetDatesMap() map[string]interface{} {
	dates := Dates{}.GetPastDates(2)
	return map[string]interface{}{
		"start_date": dates[1].String(),
		"end_date":   dates[0].String(),
	}
}

func (r *Reports) GetMap() map[string]interface{} {
	return map[string]interface{}{
		"shipment": r.GetDatesMap(),
	}
}

func (r *Reports) Get() easypost.Report {
	// make a report locally + POST to API
	_map := r.GetDatesMap()
	_report := easypost.Report{}
	Mapper{}.MapToObject(_map, &_report)
	report, err := r.client.CreateReport("shipment", &_report)
	if err != nil {
		panic("Could not create report")
	}
	return *report
}
