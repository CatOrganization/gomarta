package gomarta

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	responsePathPrefix     = "test-responses"
	trainsResponseFilename = responsePathPrefix + "/real_time_trains.json"
	busResponseFilename    = responsePathPrefix + "/real_time_bus.json"
)

func TestTrainAPIResponse(t *testing.T) {
	f, err := os.Open(trainsResponseFilename)
	require.NoError(t, err)

	buff, err := ioutil.ReadAll(f)
	require.NoError(t, err)

	var trainResponse TrainAPIResponse
	err = trainResponse.UnmarshalJSON(buff)
	require.NoError(t, err)

	expectedTrainsResponse := TrainAPIResponse{
		{
			Destination:    "Airport",
			Direction:      "S",
			EventTime:      "12/1/2019 8:11:32 PM",
			Line:           "GOLD",
			NextArrival:    "08:11:41 PM",
			Station:        "COLLEGE PARK STATION",
			TrainID:        "302506",
			WaitingSeconds: "-46",
			WaitingTime:    "Boarding",
		},
	}

	assert.Equal(t, expectedTrainsResponse, trainResponse)
}

func TestBusAPIResponse(t *testing.T) {
	f, err := os.Open(busResponseFilename)
	require.NoError(t, err)

	buff, err := ioutil.ReadAll(f)
	require.NoError(t, err)

	var busResponse BusAPIResponse
	err = busResponse.UnmarshalJSON(buff)
	require.NoError(t, err)

	expectedBusResponse := BusAPIResponse{
		{
			Adherence:         "0",
			BlockID:           "21",
			BlockAbbreviation: "110-1",
			Direction:         "Southbound",
			Latitude:          "33.7890561",
			Longitude:         "-84.386845",
			MessageTime:       "12/1/2019 8:53:01 PM",
			Route:             "110",
			StopID:            "901789",
			TimePoint:         "Arts Center Station",
			TripID:            "6861689",
			Vehicle:           "1468",
		},
	}

	assert.Equal(t, expectedBusResponse, busResponse)
}
