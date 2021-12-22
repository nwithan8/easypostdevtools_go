package easypostdevtools

type EasyPostDevTools struct {
	Addresses      Addresses      `json:"addresses"`
	Parcels        Parcels        `json:"parcels"`
	Insurance      Insurance      `json:"insurance"`
	Shipments      Shipments      `json:"shipments"`
	Options        Options        `json:"options"`
	Rates          Rates          `json:"rates"`
	Smartrates     Smartrates     `json:"smartrates"`
	TaxIdentifiers TaxIdentifiers `json:"tax_identifiers"`
	Trackers       Trackers       `json:"trackers"`
	Batch          Batch          `json:"batch"`
	CustomsItems   CustomsItems   `json:"customs_items"`
	CustomsInfos   CustomsInfos   `json:"customs_infos"`
	Events         Events         `json:"events"`
	Fees           Fees           `json:"fees"`
	Orders         Orders         `json:"orders"`
	Pickups        Pickups        `json:"pickups"`
	Reports        Reports        `json:"reports"`
	ScanForms      ScanForms      `json:"scan_forms"`
	Webhooks       Webhooks       `json:"webhooks"`
	Users          Users          `json:"users"`
	Carriers       Carriers       `json:"carriers"`
	Labels         Labels         `json:"labels"`
	PostageLabels  PostageLabels  `json:"postage_labels"`
}
