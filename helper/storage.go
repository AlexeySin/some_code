package helper

import (
	"fmt"
	"strconv"
	"time"
)

type Counter struct {
	MovesCount int64
	SumPrice   float64
	TrainLegs  []Param
}

func (c *Counter) Add(TrainLeg Param) error {
	if c.IsPassed(TrainLeg) {
		return nil
	}
	c.MovesCount++
	sum, err := strconv.ParseFloat(TrainLeg.Price, 10)
	if err != nil {
		return err
	}
	c.SumPrice += sum
	c.TrainLegs = append(c.TrainLegs, TrainLeg)
	return nil
}

func (c *Counter) FetchTime() (string, error) {
	var times float64
	for _, v := range c.TrainLegs {
		tm, err := TimeDiff(v)
		if err != nil {
			return "", err
		}
		times += tm
	}
	return formatSeconds(times), nil
}

func (c *Counter) IsPassed(trainLeg Param) bool {
	for _, v := range c.TrainLegs {
		if v == trainLeg {
			return true
		}
	}
	return false
}

func (c *Counter) InDepartures(ArrivalStationID string, depID string) bool {
	for _, v := range c.TrainLegs {
		if (v.DepartureStationID == ArrivalStationID) && (v.ArrivalStationID == depID) {
			return true
		}
	}
	return false
}

func (c *Counter) IsUniqueStations(arID string, depID string) bool {
	for _, v := range c.TrainLegs {
		if (v.DepartureStationID == depID) && (v.ArrivalStationID == arID) {
			return false
		}
	}
	return true
}

func formatSeconds(times float64) string {
	h, m, s := time.Unix(0, int64(times)*int64(time.Second)).Clock()
	return fmt.Sprintf("%v:%v:%v", h, m, s)
}
