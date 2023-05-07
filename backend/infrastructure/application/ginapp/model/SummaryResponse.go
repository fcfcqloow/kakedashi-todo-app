package model

type Summary struct {
	Start int      `json:"start"`
	End   int      `json:"end"`
	Logs  []string `json:"logs"`
	Date  []string `json:"dateList"`
}
