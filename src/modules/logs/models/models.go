package models

import (
	"time"
)

type Log struct {
	Created   time.Time `bson:"created"`
	IP        string    `bson:"ip"`
	UserAgent string    `bson:"user_agent"`
	Method    string    `bson:"method"`
	Path      string    `bson:"path"`
	Status    int       `bson:"status"`
	Message   string    `bson:"message"`
	AppID     string    `bson:"app_id"`
	UserID    string    `bson:"user_id"`
}
