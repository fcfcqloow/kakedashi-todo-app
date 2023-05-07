package model

type TasksMoveRequest struct {
	From string `json:"from"`
	To   string `json:"to"`
	Idx  int    `json:"index"`
	Task Task   `json:"task"`
}
