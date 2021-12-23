package easypostdevtools

import "github.com/EasyPost/easypost-go"

type Rates struct {
	client *easypost.Client
	tools  *EasyPostDevTools
}

func (r *Rates) Get(shipmentMap map[string]interface{}, shipment easypost.Shipment) []easypost.Rate {
	shipment = r.tools.Shipments.GetOrMakeShipment(shipmentMap, shipment)
	_rates := shipment.Rates
	rates := make([]easypost.Rate, len(_rates))
	for i, _rate := range _rates {
		rates[i] = *_rate
	}
	return rates
}
