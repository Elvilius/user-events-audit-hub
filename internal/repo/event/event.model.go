package repo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`
	UserID     string             `bson:"user_id"`
	EventType  string             `bson:"event_type"`
	Metadata   map[string]string  `bson:"metadata"`
	Timestamp  time.Time          `bson:"timestamp, default:current_timestamp"`
	SystemName string             `bson:"system_name"`
	Message    string             `bson:"message"`
	Severity   string             `bson:"severity"`
}
