package helper

import (
	"TEST_TASK/dferr"
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"time"

	xj "github.com/basgys/goxml2json"
)

const timeParseLayout = "15:04:05"

type Param struct {
	ArrivalTimeString   string `json:"-ArrivalTimeString"`
	DepartureTimeString string `json:"-DepartureTimeString"`
	TrainID             string `json:"-TrainId"`
	DepartureStationID  string `json:"-DepartureStationId"`
	ArrivalStationID    string `json:"-ArrivalStationId"`
	Price               string `json:"-Price"`
}

type trLs struct {
	TrainLeg []Param `json:"TrainLeg"`
}

type information struct {
	TrainLegs trLs `json:"TrainLegs"`
}

// FetchTrainLegs - fetch data from file by passed path, parse it, return TrainLegs as objects
func FetchTrainLegs(path string) ([]Param, error) {
	data, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer data.Close()

	info, err := ioutil.ReadAll(data)
	if err != nil {
		return nil, err
	}

	xmlData := strings.NewReader(string(info))
	jsonData, err := xj.Convert(xmlData)
	if err != nil {
		return nil, err
	}

	var legs information
	if err := json.Unmarshal([]byte(jsonData.String()), &legs); err != nil {
		return nil, err
	}

	if len(legs.TrainLegs.TrainLeg) == 0 {
		return nil, dferr.ErrNoTrainLegs
	}
	return legs.TrainLegs.TrainLeg, nil
}

// TimeDiff - return time gap in Seconds between DepartureTimeString and ArrivalTimeString in TrainLeg.
func TimeDiff(el Param) (float64, error) {
	dep := el.DepartureTimeString
	ar := el.ArrivalTimeString
	d, err := time.Parse(timeParseLayout, dep)
	if err != nil {
		return 0.0, err
	}
	a, err := time.Parse(timeParseLayout, ar)
	if err != nil {
		return 0.0, err
	}
	return d.Sub(a).Seconds() * (-1), nil
}
