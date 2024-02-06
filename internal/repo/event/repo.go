package repo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	
)

type CreateEventDto struct {
	UserId     int
	EventType  string
	SystemName string
	Message     string
	Severity    string
	Metadata    map[string]string
}

type Repo struct {
	client *mongo.Client
}

func (r *Repo) CreateEvent(ctx context.Context, createDto CreateEventDto) (string, error) {
	result, err := r.client.Database("eventsdb").Collection("events").InsertOne(ctx, createDto)
	if err != nil {
		return "", err
	}

	insertedID := result.InsertedID.(primitive.ObjectID)
	return insertedID.Hex(), nil
	
}

func (r *Repo) GetEvent() {
	return
}

func NewRepo(client *mongo.Client) Repo {
	return Repo{client: client}
}
