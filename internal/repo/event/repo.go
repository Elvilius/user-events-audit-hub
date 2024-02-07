package repo

import (
	"context"

	"github.com/Elvilius/user-events-audit-hub/internal/domain/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const DBname = "eventsdb"

func NewRepo(client *mongo.Client) Repo {
	return Repo{client: client}
}

type Repo struct {
	client *mongo.Client
}

type EventID string

func (r *Repo) CreateEvent(ctx context.Context, newEvent models.Event) (EventID, error) {
	result, err := r.client.Database(DBname).Collection("events").InsertOne(ctx, newEvent)
	if err != nil {
		return "", err
	}

	insertedID := result.InsertedID.(primitive.ObjectID)

	return EventID(insertedID.Hex()), nil

}

func (r *Repo) GetEvent() {
	return
}
