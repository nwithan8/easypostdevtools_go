package easypostdevtools

import "github.com/EasyPost/easypost-go"

type PostageLabels struct {
	client *easypost.Client
	tools  *EasyPostDevTools
}

func (p *PostageLabels) Get(shipmentMap map[string]interface{}, shipment easypost.Shipment) easypost.PostageLabel {
	shipment = p.tools.Shipments.GetOrMakeShipment(shipmentMap, shipment)
	postageLabel := shipment.PostageLabel
	return *postageLabel
}
