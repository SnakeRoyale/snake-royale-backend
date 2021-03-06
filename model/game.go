package model

type Connection struct {
	Name  string
	Token int64
}

type PostConnection struct {
	Token string `json:"token,omitempty"`
	Name  string `json:"name"`
}

type StatusResponse struct {
	Status      string `json:"status"`
	StatusCode  int    `json:"statusCode"`
	TimeToStart int    `json:"timerInSeconds,omitempty"`
}
