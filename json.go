package main

//JSON coming in
type OutagesJson struct {
	//	Messages interface{} `json:"messages"`
	Outages []outage `json:"pannes"`
}

type outage []interface{}

const (
	ClientsEffected = iota
	TimeStart       = iota
	TimeEndEstimate = iota
	EventCode       = iota
	LatLong         = iota
	OtherCode       = iota
	Other1          = iota
	Other2          = iota
	Other3          = iota
	Other4          = iota
)
