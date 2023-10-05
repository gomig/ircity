package ircity

// States get all states list
func States() []StateType {
	res := make([]StateType, 0)
	for _, s := range states {
		res = append(res, s)
	}
	return res
}

// State find state by code or return nil
func State(code uint) *StateType {
	if s, ok := states[code]; ok {
		return &s
	}
	return nil
}

// Cities get list of state cities
func Cities(state uint) []CityType {
	res := make([]CityType, 0)
	for _, c := range cities {
		if c.StateCode == state {
			res = append(res, c)
		}
	}
	return res
}

// City find city by code
func City(code uint) *CityType {
	if c, ok := cities[code]; ok {
		c.State = State(c.StateCode)
		return &c
	}
	return nil
}
