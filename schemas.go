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
	Destination    string   `json:"destination"`
	Direction      string   `json:"direction"`
	EventTime      string   `json:"event_time"`
	Line           railLine `json:"line"`
	NextArrival    string   `json:"next_arr"`
	Station        string   `json:"station"`
	TrainID        string   `json:"train_id"`
	WaitingSeconds string   `json:"waiting_seconds"`
	WaitingTime    string   `json:"waiting_time"`
}
