package easypostdevtools

import "github.com/EasyPost/easypost-go"

type Rates struct {
	client *easypost.Client
	tools  *EasyPostDevTools
}

func (r *Rates) Get(shipmentMap map[string]interface{}, shipment *easypost.Shipment) []easypost.Rate {
	if shipment == nil {
		if shipmentMap == nil {
			shipmentMap = r.tools.Shipments.GetMap(nil, nil, nil)
		}
		*shipment = r.tools.Shipments.Create(shipmentMap)
	}
	_rates := shipment.Rates
	rates := make([]easypost.Rate, len(_rates))
	for i, _rate := range _rates {
		rates[i] = *_rate
	}
	return rates
}
