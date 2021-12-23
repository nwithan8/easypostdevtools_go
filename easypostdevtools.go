package easypostdevtools

import (
	"github.com/EasyPost/easypost-go"
	"github.com/joho/godotenv"
)

type KeyType int

const (
	TEST       KeyType = 1
	PRODUCTION KeyType = 2
)

type EasyPostDevTools struct {
	Addresses      Addresses
	Parcels        Parcels
	Insurance      Insurance
	Shipments      Shipments
	Options        Options
	Rates          Rates
	Smartrates     Smartrates
	TaxIdentifiers TaxIdentifiers
	Trackers       Trackers
	Batch          Batch
	CustomsItems   CustomsItems
	CustomsInfos   CustomsInfos
	Events         Events
	Fees           Fees
	Orders         Orders
	Pickups        Pickups
	Reports        Reports
	ScanForms      ScanForms
	Webhooks       Webhooks
	Users          Users
	Carriers       Carriers
	Labels         Labels
	PostageLabels  PostageLabels
}

func (e *EasyPostDevTools) SetupKey(key string, envDir string, keyType KeyType) {
	if key == "" && (envDir == "" || (keyType != TEST && keyType != PRODUCTION)) {
		panic("Must provide either key, or envDir and keyType")
	}
	_client := easypost.Client{}
	if &key != nil {
		_client = *easypost.New(key)
	} else {
		path := envDir + "/.env"
		config, _ := godotenv.Read(path)
		switch keyType {
		case TEST:
			key = config["EASYPOST_TEST_KEY"]
		case PRODUCTION:
			key = config["EASYPOST_PROD_KEY"]
		}
		_client = *easypost.New(key)
	}
	client := &_client
	e.Addresses = Addresses{client, e}
	e.Parcels = Parcels{client, e}
	e.Insurance = Insurance{client, e}
	e.Shipments = Shipments{client, e}
	e.Options = Options{client, e}
	e.Rates = Rates{client, e}
	e.Smartrates = Smartrates{client, e}
	e.TaxIdentifiers = TaxIdentifiers{client, e}
	e.Trackers = Trackers{client, e}
	e.Batch = Batch{client, e}
	e.CustomsItems = CustomsItems{client, e}
	e.CustomsInfos = CustomsInfos{client, e}
	e.Events = Events{client, e}
	e.Fees = Fees{client, e}
	e.Orders = Orders{client, e}
	e.Pickups = Pickups{client, e}
	e.Reports = Reports{client, e}
	e.ScanForms = ScanForms{client, e}
	e.Webhooks = Webhooks{client, e}
	e.Users = Users{client, e}
	e.Carriers = Carriers{client, e}
	e.Labels = Labels{client, e}
	e.PostageLabels = PostageLabels{client, e}
}
