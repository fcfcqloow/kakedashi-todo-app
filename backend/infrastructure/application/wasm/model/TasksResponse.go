package model

type TasksResponse struct {
	Todo       []Task `json:"todo"`
	Done       []Task `json:"done"`
	Pending    []Task `json:"pending"`
	Processing []Task `json:"processing"`
}
