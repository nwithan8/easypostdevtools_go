package easypostdevtools

import "github.com/EasyPost/easypost-go"

type Shipments struct {
	client *easypost.Client
	tools  *EasyPostDevTools
}

func (s *Shipments) Create(shipmentMap map[string]interface{}) easypost.Shipment {
	_shipment := easypost.Shipment{}
	Mapper{}.MapToObject(shipmentMap, &_shipment)
	shipment, err := s.client.CreateShipment(&_shipment)
	if err != nil {
		panic(err)
	}
	return *shipment
}

func (s *Shipments) GetMap(toAddressMap map[string]interface{}, fromAddressMap map[string]interface{}, parcelMap map[string]interface{}) map[string]interface{} {
	if toAddressMap == nil || fromAddressMap == nil {
		addressMaps := s.tools.Addresses.GetMapsDifferentStates(2)
		toAddressMap = addressMaps[0]
		fromAddressMap = addressMaps[1]
	}
	if parcelMap == nil {
		parcelMap = s.tools.Parcels.GetMap()
	}
	shipmentMap := map[string]interface{}{
		"to_address":   toAddressMap,
		"from_address": fromAddressMap,
		"parcel":       parcelMap,
	}
	return shipmentMap
}
func (s *Shipments) GetReturnMap(toAddressMap map[string]interface{}, fromAddressMap map[string]interface{}, parcelMap map[string]interface{}) map[string]interface{} {
	_shipmentMap := s.GetMap(toAddressMap, fromAddressMap, parcelMap)
	_shipmentMap["is_return"] = true
	return _shipmentMap
}

func (s *Shipments) Get(toAddressMap map[string]interface{}, fromAddressMap map[string]interface{}, parcelMap map[string]interface{}) easypost.Shipment {
	shipmentMap := s.GetMap(toAddressMap, fromAddressMap, parcelMap)
	return s.Create(shipmentMap)
}

func (s *Shipments) GetReturn(toAddressMap map[string]interface{}, fromAddressMap map[string]interface{}, parcelMap map[string]interface{}) easypost.Shipment {
	shipmentMap := s.GetReturnMap(toAddressMap, fromAddressMap, parcelMap)
	return s.Create(shipmentMap)
}

func (s *Shipments) AddInsurance(shipment easypost.Shipment, amount float64) easypost.Shipment {
	_shipment := s.tools.Insurance.Insure(shipment, amount)
	return *_shipment
}

func (s *Shipments) Refund(shipment easypost.Shipment) easypost.Shipment {
	_shipment, err := s.client.RefundShipment(shipment.ID)
	if err != nil {
		panic(err)
	}
	return *_shipment
}

func (s *Shipments) MarkForReturn(shipmentMap map[string]interface{}) map[string]interface{} {
	shipmentMap["is_return"] = true
	return shipmentMap
}
