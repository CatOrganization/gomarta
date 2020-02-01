package gomarta

type direction string

const (
	directionNorth = "N"
	directionSouth = "S"
	directionEast  = "E"
	directionWest  = "W"
)

type railLine string

const (
	railLineRed   = "Red"
	railLineGold  = "Gold"
	railLineBlue  = "Blue"
	railLineGreen = "Green"
)

// TODO find all the possible waiting times
type waitingTime string

// easyjson:json
type TrainAPIResponse []Train

// easyjson:json
type Train struct {
	Destination    string    `json:"DESTINATION"`
	Direction      direction `json:"DIRECTION"`
	EventTime      string    `json:"EVENT_TIME"`
	Line           railLine  `json:"LINE"`
	NextArrival    string    `json:"NEXT_ARR"`
	Station        string    `json:"STATION"`
	TrainID        string    `json:"TRAIN_ID"`
	WaitingSeconds string    `json:"WAITING_SECONDS"`
	WaitingTime    string    `json:"WAITING_TIME"`
}

// easyjson:json
type BusAPIResponse []Bus

// easyjson:json
type Bus struct {
	Adherence         string `json:"ADHERENCE"`
	BlockID           string `json:"BLOCKID"`
	BlockAbbreviation string `json:"BLOCK_ABBR"` // TODO: find out what this actually is
	Direction         string `json:"DIRECTION"`
	Latitude          string `json:"LATITUDE"`
	Longitude         string `json:"LONGITUDE"`
	MessageTime       string `json:"MSGTIME"`
	Route             string `json:"ROUTE"`
	StopID            string `json:"STOPID"`
	TimePoint         string `json:"TIMEPOINT"`
	TripID            string `json:"TRIPID"`
	Vehicle           string `json:"VEHICLE"`
}
