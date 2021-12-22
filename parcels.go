package easypostdevtools

import "github.com/EasyPost/easypost-go"

type Parcels struct {
	client *easypost.Client
	tools  *EasyPostDevTools
}

func (p *Parcels) GetMap() map[string]interface{} {
	random := Random{}
	return map[string]interface{}{
		"weight": random.GetRandomFloatInRange(0.0, 100.0),
		"height": random.GetRandomFloatInRange(0.0, 100.0),
		"width":  random.GetRandomFloatInRange(0.0, 100.0),
		"length": random.GetRandomFloatInRange(0.0, 100.0),
	}
}

func (p *Parcels) Get() easypost.Parcel {
	// Create a new parcel locally + POST to API
	_map := p.GetMap()
	_parcel := easypost.Parcel{}
	Mapper{}.MapToObject(_map, &_parcel)
	parcel, err := p.client.CreateParcel(&_parcel)
	if err != nil {
		panic(err)
	}
	return *parcel
}
