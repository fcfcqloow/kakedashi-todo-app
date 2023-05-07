package model

type Topic struct {
	Id          string `json:"id"`
	Value       string `json:"value"`
	Color       string `json:"color"`
	CreatedDate int    `json:"created"`
	DoneDate    int    `json:"done"`
}
