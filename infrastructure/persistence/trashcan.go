package persistence

import (
	"context"
	"gomibakokun_backend/domain"
	"gomibakokun_backend/domain/repository"
	"log"

	"cloud.google.com/go/firestore"
)

type trashcanPersistence struct{}

func NewTrashcanPersistence() repository.TrashcanRepository {
	return &trashcanPersistence{}
}

func (tp trashcanPersistence) CreateTrashcan(ctx context.Context, client *firestore.Client, trashcan *domain.Trashcan) error {
	_, err := client.Collection("trashcans").Doc(trashcan.ID).Set(ctx, map[string]interface{}{
		"ID":        trashcan.ID,
		"latitude":  trashcan.Latitude,
		"longitude": trashcan.Longitude,
		"trashType": trashcan.TrashType,
	})

	if err != nil {
		log.Printf("An error has occurred to create trashcan: %s", err)
	}

	return err
}
