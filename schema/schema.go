package schema

import "time"

type Request struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type PayloadDataHistory struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
	CreatedOn time.Time `json:"created_on"`
}

type Payload struct {
	OrderId string `json:"order_id"`
	History []PayloadDataHistory `json:"history"`
}

type History struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Response struct {
	OrderId string `json:"order_id"`
	History []History `json:"history"`
}
