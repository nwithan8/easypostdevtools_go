package easypostdevtools

import "github.com/EasyPost/easypost-go"

type Pickups struct {
	client *easypost.Client
	tools  *EasyPostDevTools
}

func (p *Pickups) GetMap() map[string]interface{} {
	_map := Mapper{}.GetMapFromJsonFile(PICKUPS_JSON)
	addressMaps := p.tools.Addresses.GetMapsSameCountry(2)
	_toAddress := addressMaps[0]
	_fromAddress := addressMaps[1]
	parcelMap := p.tools.Parcels.GetMap()
	shipmentMap := p.tools.Shipments.GetMap(_toAddress, _fromAddress, parcelMap)
	_map["shipment"] = shipmentMap
	dates := Dates{}.GetFutureDates(2)
	_map["min_datetime"] = dates[0].String()
	_map["max_datetime"] = dates[1].String()
	return _map
}
