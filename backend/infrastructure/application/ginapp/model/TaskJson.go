package model

type Task struct {
	Id       string   `json:"id"`
	Value    string   `json:"value"`
	DeadLine int      `json:"deadLine"`
	Topics   []string `json:"topics"`
	Priority *int     `json:"priority"`
	DoneDate *int     `json:"done"`
	Created  int      `json:"created"`
}
