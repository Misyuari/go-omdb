package model

import "time"

type Log struct {
	ID             uint
	Module         string
	RequestPayload interface{} //json
	ServerType     string
	Timestamp      time.Time
}
