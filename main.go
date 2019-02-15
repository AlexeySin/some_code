package main

import (
	"TEST_TASK/helper"
	"fmt"
)

const XMLFilePath = "./data.xml"

func main() {

	trainLegs, err := helper.FetchTrainLegs(XMLFilePath)
	if err != nil {
		helper.Out(err)
		return
	}

	var globCount helper.Counter

	start, err := helper.PluckBest(trainLegs, &globCount)
	if err != nil {
		helper.Out(err)
		return
	}

	for range trainLegs {
		station, err := helper.SearchByArrival(trainLegs, start, &globCount)
		if err != nil {
			helper.Out(err)
			return
		}
		start = station
	}

	tm, err := globCount.FetchTime()
	if err != nil {
		helper.Out(err)
		return
	}

	helper.Out(fmt.Sprintf("\nBEST PATH\nPrice: %v $\nMoves: %v\nTime: %s\n\n", globCount.SumPrice, globCount.MovesCount, tm))
	for _, v := range globCount.TrainLegs {
		helper.Out(fmt.Sprintf("%v\n", v))
	}
}
