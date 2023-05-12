package model

type SettingsJson struct {
	TaskLimit               int    `json:"taskLimit"`
	Sortable                bool   `json:"sortable"`
	TopicColor              bool   `json:"topicColor"`
	Mode                    string `json:"mode"`
	TopicListType           string `json:"topicListType"`
	NotificationIntervalSec int    `json:"notificationIntervalSec"`
	MyColor                 string `json:"myColor"`
	StopNotification        bool   `json:"stopNotification"`
}
