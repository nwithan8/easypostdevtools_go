package easypostdevtools

import "github.com/EasyPost/easypost-go"

type Addresses struct {
	client *easypost.Client
	tools  *EasyPostDevTools
}

type ADDRESS_RELATIONSHIP int

const (
	SAME_STATE          ADDRESS_RELATIONSHIP = 1
	DIFFERENT_STATES    ADDRESS_RELATIONSHIP = 2
	SAME_COUNTRY        ADDRESS_RELATIONSHIP = 3
	DIFFERENT_COUNTRIES ADDRESS_RELATIONSHIP = 4
)

func (a *Addresses) Create(addressMap map[string]interface{}) easypost.Address {
	// convert map into an easypost.Address locally + call POST to save Address on API
	_address := easypost.Address{}
	Mapper{}.MapToObject(addressMap, &_address)
	address, err := a.client.CreateAddress(&_address, nil)
	if err != nil {
		panic("Error creating address on API")
	}
	return *address
}

func (a *Addresses) GetMap(country *AddressesConstantsCountryEnum, state *AddressesConstantsStateEnum) map[string]interface{} {
	constants := Constants{}
	addressFile := constants.GetAddressFile(country, state)
	return Mapper{}.GetMapFromJsonFile(addressFile)
}

func (a *Addresses) Get(country *AddressesConstantsCountryEnum, state *AddressesConstantsStateEnum) easypost.Address {
	_map := a.GetMap(country, state)
	return a.Create(_map)
}

func (a *Addresses) GetMapsSameState(amount int) []map[string]interface{} {
	constants := Constants{}
	state := constants.GetRandomState()
	_stateEnum := constants.Addresses.ToStateAddressEnum(state)
	addressFile := constants.GetAddressFile(nil, &_stateEnum)
	return Mapper{}.GetMapsFromJsonFile(addressFile, amount, false)
}

func (a *Addresses) GetSameState(amount int) []easypost.Address {
	_maps := a.GetMapsSameState(amount)
	addresses := []easypost.Address{}
	for _, _map := range _maps {
		addresses = append(addresses, a.Create(_map))
	}
	return addresses
}

func (a *Addresses) GetMapsDifferentStates(amount int) []map[string]interface{} {
	constants := Constants{}
	if amount > len(constants.Addresses.stateEnums) {
		panic("Amount cannot be greater than " + string(rune(len(constants.Addresses.stateEnums))))
	}
	_maps := []map[string]interface{}{}
	stateEnumsIntefaces := constants.Addresses.GetStateEnumsInterfaces()
	states := Random{}.GetRandomItemsFromList(stateEnumsIntefaces, amount, false)
	for _, state := range states {
		_stateEnum := state.(*AddressesConstantsStateEnum)
		addressFile := constants.GetAddressFile(nil, _stateEnum)
		_maps = append(_maps, Mapper{}.GetMapFromJsonFile(addressFile))
	}
	return _maps
}

func (a *Addresses) GetDifferentStates(amount int) []easypost.Address {
	_maps := a.GetMapsDifferentStates(amount)
	addresses := []easypost.Address{}
	for _, _map := range _maps {
		addresses = append(addresses, a.Create(_map))
	}
	return addresses
}

func (a *Addresses) GetMapsSameCountry(amount int) []map[string]interface{} {
	constants := Constants{}
	country := constants.GetRandomCountry()
	_countryEnum := constants.Addresses.ToCountryAddressEnum(country)
	addressFile := constants.GetAddressFile(&_countryEnum, nil)
	return Mapper{}.GetMapsFromJsonFile(addressFile, amount, false)
}

func (a *Addresses) GetSameCountry(amount int) []easypost.Address {
	_maps := a.GetMapsSameCountry(amount)
	addresses := []easypost.Address{}
	for _, _map := range _maps {
		addresses = append(addresses, a.Create(_map))
	}
	return addresses
}

func (a *Addresses) GetMapsDifferentCountries(amount int) []map[string]interface{} {
	constants := Constants{}
	if amount > len(constants.Addresses.countryEnums) {
		panic("Amount cannot be greater than " + string(rune(len(constants.Addresses.countryEnums))))
	}
	_maps := []map[string]interface{}{}
	countryEnumsIntefaces := constants.Addresses.GetCountryEnumsInterfaces()
	countries := Random{}.GetRandomItemsFromList(countryEnumsIntefaces, amount, false)
	for _, country := range countries {
		_countryEnum := country.(*AddressesConstantsCountryEnum)
		addressFile := constants.GetAddressFile(_countryEnum, nil)
		_maps = append(_maps, Mapper{}.GetMapFromJsonFile(addressFile))
	}
	return _maps
}

func (a *Addresses) GetDifferentCountries(amount int) []easypost.Address {
	_maps := a.GetMapsDifferentCountries(amount)
	addresses := []easypost.Address{}
	for _, _map := range _maps {
		addresses = append(addresses, a.Create(_map))
	}
	return addresses
}

func (a *Addresses) GetMapsAmount(relationship ADDRESS_RELATIONSHIP, amount int) []map[string]interface{} {
	switch relationship {
	case SAME_STATE:
		return a.GetMapsSameState(amount)
	case DIFFERENT_STATES:
		return a.GetMapsDifferentStates(amount)
	case SAME_COUNTRY:
		return a.GetMapsSameCountry(amount)
	case DIFFERENT_COUNTRIES:
		return a.GetMapsDifferentCountries(amount)
	default:
		panic("Relationship not supported")
	}
}

func (a *Addresses) GetAmount(relationship ADDRESS_RELATIONSHIP, amount int) []easypost.Address {
	switch relationship {
	case SAME_STATE:
		return a.GetSameState(amount)
	case DIFFERENT_STATES:
		return a.GetDifferentStates(amount)
	case SAME_COUNTRY:
		return a.GetSameCountry(amount)
	case DIFFERENT_COUNTRIES:
		return a.GetDifferentCountries(amount)
	default:
		panic("Relationship not supported")
	}
}
