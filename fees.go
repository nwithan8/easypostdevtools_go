package easypostdevtools

import "github.com/EasyPost/easypost-go"

type Fees struct {
	client *easypost.Client
	tools  *EasyPostDevTools
}

func (f *Fees) Get(shipmentMap map[string]interface{}, shipment easypost.Shipment) []easypost.Fee {
	shipment = f.tools.Shipments.GetOrMakeShipment(shipmentMap, shipment)
	_fees := shipment.Fees
	fees := make([]easypost.Fee, len(_fees))
	for i, _fee := range _fees {
		fees[i] = *_fee
	}
	return fees
}
