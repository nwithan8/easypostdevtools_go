package easypostdevtools

import (
	"fmt"
	"github.com/EasyPost/easypost-go"
)

type Insurance struct {
	client *easypost.Client
	tools  *EasyPostDevTools
}

func (i *Insurance) GetMap(amount float64) map[string]interface{} {
	_map := map[string]interface{}{
		"amount": amount,
	}
	return _map
}

func (i *Insurance) Insure(shipment easypost.Shipment, amount float64) *easypost.Shipment {
	_map := i.GetMap(amount)
	insureShipment, err := i.client.InsureShipment(shipment.ID, fmt.Sprintf("%.2f", _map["amount"]))
	if err != nil {
		return nil
	}
	return insureShipment
}
