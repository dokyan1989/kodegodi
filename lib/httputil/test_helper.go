package httputil

import "time"

type SampleResponse struct {
	Code    uint32             `json:"code"`
	Message string             `json:"message"`
	Data    SampleResponseData `json:"data"`
}

type SampleResponseData struct {
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	CreatedAt time.Time `json:"createdAt"`
}

func GetTimeNow() time.Time {
	return time.Now().Round(time.Second)
}
