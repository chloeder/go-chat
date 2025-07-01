package repositories

import (
	"context"
	"log"

	"github.com/kooroshh/fiber-boostrap/app/models"
	"github.com/kooroshh/fiber-boostrap/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
)

func InsertMessage(ctx context.Context, data models.MessagePayload) error {
	_, err := database.MongoDB.InsertOne(ctx, data)
	if err != nil {
		log.Println("Error inserting message:", err)
		return err
	}
	return nil
}

func GetMessages(ctx context.Context) ([]models.MessagePayload, error) {
	cursor, err := database.MongoDB.Find(ctx, bson.D{})
	if err != nil {
		log.Println("Error getting messages:", err)
		return nil, err
	}

	var messages []models.MessagePayload
	for cursor.Next(ctx) {
		var message models.MessagePayload
		if err := cursor.Decode(&message); err != nil {
			log.Println("Error decoding message:", err)
			return nil, err
		}
		messages = append(messages, message)
	}

	cursor.Close(ctx)
	return messages, nil
}
