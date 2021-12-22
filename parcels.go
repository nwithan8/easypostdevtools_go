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
