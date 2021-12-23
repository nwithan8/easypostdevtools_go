package easypostdevtools

import "github.com/EasyPost/easypost-go"

type Smartrates struct {
	client *easypost.Client
	tools  *EasyPostDevTools
}

func (s *Smartrates) Get(shipmentMap map[string]interface{}, shipment easypost.Shipment) []easypost.Rate {
	shipment = s.tools.Shipments.GetOrMakeShipment(shipmentMap, shipment)
	_rates, err := s.client.GetShipmentSmartrates(shipment.ID)
	if err != nil {
		panic("Could not get shipment smartrates.")
	}
	rates := make([]easypost.Rate, len(_rates))
	for i, _rate := range _rates {
		rates[i] = *_rate
	}
	return rates
}
