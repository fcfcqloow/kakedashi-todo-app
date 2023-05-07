package model

type WebsocketJsonValue struct {
	Type    string      `json:"valueType"`
	Created int         `json:"created"`
	Json    interface{} `json:"json"`
}

type WebsocketJson struct {
	Type  string             `json:"messageType"`
	Value WebsocketJsonValue `json:"value"`
}
