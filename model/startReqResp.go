package model

type StartRequestModel struct {
	Difficulty string `json:"difficulty"`
	Side       string `json:"side"`
}
type StartResponseModel struct {
	SessionId string `json:"sessionId"`
}
