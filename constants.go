package easypostdevtools

type JsonFile struct {
	Name         string
	ParentFolder string
}

func (j JsonFile) GetJsonPath() string {
	return j.ParentFolder + "/" + j.Name + ".min.json"
}

type JsonAddressFile struct {
	JsonFile JsonFile
}

func (j JsonAddressFile) GetAddressFile() string {
	return "json/address/" + j.JsonFile.GetJsonPath()
}

type AddressesConstantsEnum struct {
	name            string
	JsonAddressFile JsonAddressFile
}

type AddressesConstants struct {
	countryEnums map[string]AddressesConstantsEnum
	stateEnums   map[string]AddressesConstantsEnum
}

func (c *AddressesConstants) Init() {
	c.countryEnums = make(map[string]AddressesConstantsEnum)
	c.countryEnums["UNITED_STATES"] = AddressesConstantsEnum{name: "UNITED_STATES", JsonAddressFile: JsonAddressFile{JsonFile: JsonFile{Name: "US", ParentFolder: "united-states"}}}
	c.countryEnums["CANADA"] = AddressesConstantsEnum{name: "CANADA", JsonAddressFile: JsonAddressFile{JsonFile: JsonFile{Name: "BC", ParentFolder: "canada"}}}
	c.countryEnums["CHINA"] = AddressesConstantsEnum{name: "CHINA", JsonAddressFile: JsonAddressFile{JsonFile: JsonFile{Name: "BJ", ParentFolder: "china"}}}
	c.countryEnums["HONG_KONG"] = AddressesConstantsEnum{name: "HONG_KONG", JsonAddressFile: JsonAddressFile{JsonFile: JsonFile{Name: "HK", ParentFolder: "china"}}}
	c.countryEnums["UNITED_KINGDOM"] = AddressesConstantsEnum{name: "UNITED_KINGDOM", JsonAddressFile: JsonAddressFile{JsonFile: JsonFile{Name: "UK", ParentFolder: "europe"}}}
	c.countryEnums["GERMANY"] = AddressesConstantsEnum{name: "GERMANY", JsonAddressFile: JsonAddressFile{JsonFile: JsonFile{Name: "DE", ParentFolder: "europe"}}}
	c.countryEnums["SPAIN"] = AddressesConstantsEnum{name: "SPAIN", JsonAddressFile: JsonAddressFile{JsonFile: JsonFile{Name: "ES", ParentFolder: "europe"}}}
	c.countryEnums["MEXICO"] = AddressesConstantsEnum{name: "MEXICO", JsonAddressFile: JsonAddressFile{JsonFile: JsonFile{Name: "MX", ParentFolder: "mexico"}}}
	c.countryEnums["AUSTRALIA"] = AddressesConstantsEnum{name: "AUSTRALIA", JsonAddressFile: JsonAddressFile{JsonFile: JsonFile{Name: "VT", ParentFolder: "australia"}}}
	c.stateEnums = make(map[string]AddressesConstantsEnum)
	c.stateEnums["ARIZONA"] = AddressesConstantsEnum{name: "ARIZONA", JsonAddressFile: JsonAddressFile{JsonFile: JsonFile{Name: "AZ", ParentFolder: "united-states"}}}
	c.stateEnums["CALIFORNIA"] = AddressesConstantsEnum{name: "CALIFORNIA", JsonAddressFile: JsonAddressFile{JsonFile: JsonFile{Name: "CA", ParentFolder: "united-states"}}}
	c.stateEnums["IDAHO"] = AddressesConstantsEnum{name: "IDAHO", JsonAddressFile: JsonAddressFile{JsonFile: JsonFile{Name: "ID", ParentFolder: "united-states"}}}
	c.stateEnums["KANSAS"] = AddressesConstantsEnum{name: "KANSAS", JsonAddressFile: JsonAddressFile{JsonFile: JsonFile{Name: "KS", ParentFolder: "united-states"}}}
	c.stateEnums["NEVADA"] = AddressesConstantsEnum{name: "NEVADA", JsonAddressFile: JsonAddressFile{JsonFile: JsonFile{Name: "NV", ParentFolder: "united-states"}}}
	c.stateEnums["NEW_YORK"] = AddressesConstantsEnum{name: "NEW_YORK", JsonAddressFile: JsonAddressFile{JsonFile: JsonFile{Name: "NY", ParentFolder: "united-states"}}}
	c.stateEnums["TEXAS"] = AddressesConstantsEnum{name: "TEXAS", JsonAddressFile: JsonAddressFile{JsonFile: JsonFile{Name: "TX", ParentFolder: "united-states"}}}
	c.stateEnums["UTAH"] = AddressesConstantsEnum{name: "UTAH", JsonAddressFile: JsonAddressFile{JsonFile: JsonFile{Name: "UT", ParentFolder: "united-states"}}}
	c.stateEnums["WASHINGTON"] = AddressesConstantsEnum{name: "WASHINGTON", JsonAddressFile: JsonAddressFile{JsonFile: JsonFile{Name: "WA", ParentFolder: "united-states"}}}
}

type JsonFilePath string
type AddressesConstantsCountryEnum string
type AddressesConstantsStateEnum string

const (
	CUSTOMS_ITEMS_JSON JsonFilePath = "json/customs_items.json"
	CUSTOMS_INFO_JSON               = "json/customs_info.json"
	CARRIERS_JSON                   = "json/carriers.json"
	LABEL_OPTIONS_JSON              = "json/label_options.json"
	TRACKERS_JSON                   = "json/trackers.json"
	OPTIONS_JSON                    = "json/options.json"
	PICKUPS_JSON                    = "json/pickups.json"
)

const (
	UNITED_STATES  AddressesConstantsCountryEnum = "UNITED_STATES"
	CANADA                                       = "CANADA"
	CHINA                                        = "CHINA"
	HONG_KONG                                    = "HONG_KONG"
	UNITED_KINGDOM                               = "UNITED_KINGDOM"
	GERMANY                                      = "GERMANY"
	SPAIN                                        = "SPAIN"
	MEXICO                                       = "MEXICO"
	AUSTRALIA                                    = "AUSTRALIA"
)

const (
	ARIZONA    AddressesConstantsStateEnum = "ARIZONA"
	CALIFORNIA                             = "CALIFORNIA"
	IDAHO                                  = "IDAHO"
	KANSAS                                 = "KANSAS"
	NEVADA                                 = "NEVADA"
	NEW_YORK                               = "NEW_YORK"
	TEXAS                                  = "TEXAS"
	UTAH                                   = "UTAH"
	WASHINGTON                             = "WASHINGTON"
)

type Constants struct {
	Addresses AddressesConstants
}

func (c *Constants) Init() {
	c.Addresses = AddressesConstants{}
	c.Addresses.Init()
}

func (c *Constants) GetStateAddressFile(state AddressesConstantsStateEnum) string {
	c.Init()
	return c.Addresses.stateEnums[string(state)].JsonAddressFile.GetAddressFile()
}

func (c *Constants) GetRandomStateAddressFile() string {
	c.Init()
	var keys []interface{}
	for k := range c.Addresses.stateEnums {
		keys = append(keys, k)
	}
	random := Random{}
	selectedKey := random.GetRandomItemFromList(keys)
	return c.Addresses.stateEnums[selectedKey.(string)].JsonAddressFile.GetAddressFile()
}

func (c *Constants) GetRandomCountryAddressFile() string {
	c.Init()
	var keys []interface{}
	for k := range c.Addresses.countryEnums {
		keys = append(keys, k)
	}
	random := Random{}
	selectedKey := random.GetRandomItemFromList(keys)
	return c.Addresses.countryEnums[selectedKey.(string)].JsonAddressFile.GetAddressFile()
}

func (c *Constants) GetCountryAddressFile(country AddressesConstantsCountryEnum) string {
	c.Init()
	return c.Addresses.stateEnums[string(country)].JsonAddressFile.GetAddressFile()
}

func (c *Constants) GetAddressFile(country *AddressesConstantsCountryEnum, state *AddressesConstantsStateEnum) string {
	if country == nil && state == nil {
		panic("Must specify either country or state")
	}
	// if country is UNITED_STATES, then we need to check if state is specified
	if *country == UNITED_STATES {
		return c.GetRandomStateAddressFile()
	} else {
		return c.GetCountryAddressFile(*country)
	}
}

func (c *Constants) GetRandomAddressFile(country *AddressesConstantsCountryEnum, state *AddressesConstantsStateEnum) string {
	c.Init()
	if country != nil {
		return c.GetCountryAddressFile(*country)
	} else if state != nil {
		return c.GetStateAddressFile(*state)
	} else {
		random := Random{}
		if random.GetRandomBoolean() {
			return c.GetRandomCountryAddressFile()
		} else {
			return c.GetRandomStateAddressFile()
		}
	}
}
