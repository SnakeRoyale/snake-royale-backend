package model

type PostSnake struct {
	Token string `json:"token,omitempty"`
	Name  string `json:"name"`
}

type GetSnake struct {
	Status      string `json:"status"`
	StatusCode  int    `json:"statusCode"`
	TimeToStart int    `json:"timerInSeconds,omitempty"`
}
