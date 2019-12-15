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

type RailEventAPIResponse struct {
	Events []RailEvent
}

// TODO: find better name
type RailEvent struct {
	Destination    string    `json:"destination"`
	Direction      direction `json:"direction"`
	EventTime      string    `json:"event_time"`
	Line           railLine  `json:"line"`
	NextArrival    string    `json:"next_arr"`
	Station        string    `json:"station"`
	TrainID        string    `json:"train_id"`
	WaitingSeconds string    `json:"waiting_seconds"`
	WaitingTime    string    `json:"waiting_time"`
}

type BusEventAPIResponse struct {
	Events []BusEvent
}

type BusEvent struct {
	Adherence string `json:"adherence"`
	BlockID   string `json:"blockid"`
	// TODO: find out what this actually is
	BlockAbbreviation string `json:"block_abbr"`
	Direction         string `json:"direction"`
	Latitude          string `json:"latitude"`
	Longitude         string `json:"longitude"`
	MessageTime       string `json:"msgtime"`
	Route             string `json:"route"`
	StopID            string `json:"stopid"`
	TimePoint         string `json:"timepoint"`
	TripID            string `json:"tripid"`
	Vehicle           string `json:"vehicle"`
}
