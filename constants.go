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

func (c *AddressesConstants) GetCountryEnums() []AddressesConstantsEnum {
	enums := make([]AddressesConstantsEnum, 0)
	for _, v := range c.countryEnums {
		enums = append(enums, v)
	}
	return enums
}

func (c *AddressesConstants) GetCountryEnumsInterfaces() []interface{} {
	values := []interface{}{}
	for _, enum := range c.GetCountryEnums() {
		values = append(values, enum)
	}
	return values
}

func (c *AddressesConstants) GetStateEnums() []AddressesConstantsEnum {
	enums := make([]AddressesConstantsEnum, 0)
	for _, v := range c.stateEnums {
		enums = append(enums, v)
	}
	return enums
}

func (c *AddressesConstants) GetStateEnumsInterfaces() []interface{} {
	values := []interface{}{}
	for _, enum := range c.GetStateEnums() {
		values = append(values, enum)
	}
	return values
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

func (c *AddressesConstants) ToCountryAddressEnum(addressEnum AddressesConstantsEnum) AddressesConstantsCountryEnum {
	switch addressEnum.name {
	case "UNITED_STATES":
		return UNITED_STATES
	case "CANADA":
		return CANADA
	case "CHINA":
		return CHINA
	case "HONG_KONG":
		return HONG_KONG
	case "UNITED_KINGDOM":
		return UNITED_KINGDOM
	case "GERMANY":
		return GERMANY
	case "SPAIN":
		return SPAIN
	case "MEXICO":
		return MEXICO
	case "AUSTRALIA":
		return AUSTRALIA
	}
	panic("Unknown country")
}

func (c *AddressesConstants) ToStateAddressEnum(addressEnum AddressesConstantsEnum) AddressesConstantsStateEnum {
	switch addressEnum.name {
	case "ARIZONA":
		return ARIZONA
	case "CALIFORNIA":
		return CALIFORNIA
	case "IDAHO":
		return IDAHO
	case "KANSAS":
		return KANSAS
	case "NEVADA":
		return NEVADA
	case "NEW_YORK":
		return NEW_YORK
	case "TEXAS":
		return TEXAS
	case "UTAH":
		return UTAH
	case "WASHINGTON":
		return WASHINGTON
	}
	panic("Unknown state")
}

type AddressesConstantsCountryEnum string
type AddressesConstantsStateEnum string

const (
	CUSTOMS_ITEMS_JSON string = "json/customs_items.json"
	CUSTOMS_INFO_JSON         = "json/customs_info.json"
	CARRIERS_JSON             = "json/carriers.json"
	LABEL_OPTIONS_JSON        = "json/label_options.json"
	TRACKERS_JSON             = "json/trackers.json"
	OPTIONS_JSON              = "json/options.json"
	PICKUPS_JSON              = "json/pickups.json"
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

func (c *Constants) GetCountry(country AddressesConstantsCountryEnum) AddressesConstantsEnum {
	c.Init()
	return c.Addresses.countryEnums[string(country)]
}

func (c *Constants) GetRandomCountry() AddressesConstantsEnum {
	c.Init()
	var keys []interface{}
	for k := range c.Addresses.countryEnums {
		keys = append(keys, k)
	}
	selectedKey := Random{}.GetRandomItemFromList(keys)
	return c.Addresses.countryEnums[selectedKey.(string)]
}

func (c *Constants) GetCountryAddressFile(country AddressesConstantsCountryEnum) string {
	_country := c.GetCountry(country)
	return _country.JsonAddressFile.GetAddressFile()
}

func (c *Constants) GetRandomCountryAddressFile() string {
	country := c.GetRandomCountry()
	return country.JsonAddressFile.GetAddressFile()
}

func (c *Constants) GetState(state AddressesConstantsStateEnum) AddressesConstantsEnum {
	c.Init()
	return c.Addresses.stateEnums[string(state)]
}

func (c *Constants) GetRandomState() AddressesConstantsEnum {
	c.Init()
	var keys []interface{}
	for k := range c.Addresses.stateEnums {
		keys = append(keys, k)
	}
	selectedKey := Random{}.GetRandomItemFromList(keys)
	return c.Addresses.stateEnums[selectedKey.(string)]
}

func (c *Constants) GetStateAddressFile(state AddressesConstantsStateEnum) string {
	_state := c.GetState(state)
	return _state.JsonAddressFile.GetAddressFile()
}

func (c *Constants) GetRandomStateAddressFile() string {
	state := c.GetRandomState()
	return state.JsonAddressFile.GetAddressFile()
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
