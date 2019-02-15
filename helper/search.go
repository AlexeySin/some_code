package helper

func byPrice(trainLegs []Param) Param {
	minPrice := trainLegs[0]
	for _, v := range trainLegs {
		if v.Price < minPrice.Price {
			minPrice = v
		}
	}
	return minPrice
}

func PluckBest(trainLegs []Param, globCount *Counter) (Param, error) {
	result := byPrice(trainLegs)
	if err := globCount.Add(result); err != nil {
		return Param{}, err
	}
	return result, nil
}

func conditions(v, arrival Param, globCount *Counter) bool {
	return (v.DepartureStationID == arrival.ArrivalStationID) &&
		(globCount.IsPassed(v) == false) &&
		(globCount.InDepartures(v.ArrivalStationID, v.DepartureStationID) == false) &&
		(globCount.IsUniqueStations(v.ArrivalStationID, v.DepartureStationID) == true)
}

func SearchByArrival(trLs []Param, arrival Param, globCount *Counter) (Param, error) {
	var stations []Param
	for _, v := range trLs {
		if conditions(v, arrival, globCount) {
			stations = append(stations, v)
		}
	}
	if len(stations) == 0 {
		return Param{}, nil
	}
	best, err := PluckBest(stations, globCount)
	if err != nil {
		return Param{}, err
	}
	if err := globCount.Add(best); err != nil {
		return Param{}, err
	}
	return best, nil
}
