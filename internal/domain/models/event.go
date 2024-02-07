package models

type Event struct {
	ID         string
	UserId     int
	EventType  string
	SystemName string
	Message    string
	Severity   string
	Metadata   map[string]string	
}
