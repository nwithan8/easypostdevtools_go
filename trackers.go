package easypostdevtools

import "github.com/EasyPost/easypost-go"

type Trackers struct {
	client *easypost.Client
	tools  *EasyPostDevTools
}

func (t *Trackers) GetMap() map[string]interface{} {
	return Mapper{}.GetMapFromJsonFile(TRACKERS_JSON)
}

func (t *Trackers) GetMapAsTrackerOptions(trackerMap map[string]interface{}) easypost.CreateTrackerOptions {
	if trackerMap == nil {
		trackerMap = t.GetMap()
	}
	_options := easypost.CreateTrackerOptions{}
	Mapper{}.MapToObject(trackerMap, &_options)
	return _options
}

func (t *Trackers) Get() easypost.Tracker {
	// Create a new tracker locally
	_options := t.GetMapAsTrackerOptions(nil)
	parcel, err := t.client.CreateTracker(&_options)
	if err != nil {
		panic(err)
	}
	return *parcel
}
