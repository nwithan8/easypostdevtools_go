package easypostdevtools

import "github.com/EasyPost/easypost-go"

type Fees struct {
	client *easypost.Client
	tools  *EasyPostDevTools
}

func (f *Fees) Get(shipmentMap map[string]interface{}, shipment easypost.Shipment) []easypost.Fee {
	if &shipment == nil {
		if &shipmentMap == nil {
			shipmentMap = f.tools.Shipments.GetMap(nil, nil, nil)
		}
		shipment = f.tools.Shipments.Create(shipmentMap)
	}
	_fees := shipment.Fees
	fees := make([]easypost.Fee, len(_fees))
	for i, _fee := range _fees {
		fees[i] = *_fee
	}
	return fees
}
