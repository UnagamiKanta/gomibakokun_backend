package persistence

import (
	"context"
	"gomibakokun_backend/domain"
	"gomibakokun_backend/domain/repository"
	"log"

	"cloud.google.com/go/firestore"
)

type trashcanPersistence struct {
	client *firestore.Client
}

func NewTrashcanPersistence(client *firestore.Client) repository.TrashcanRepository {
	return &trashcanPersistence{client: client}
}

func (tp trashcanPersistence) CreateTrashcan(ctx context.Context, trashcan *domain.Trashcan) error {
	_, err := tp.client.Collection("trashcans").Doc(trashcan.ID).Set(ctx, map[string]interface{}{
		"ID":              trashcan.ID,
		"latitude":        trashcan.Latitude,
		"longitude":       trashcan.Longitude,
		"trashType":       trashcan.TrashType,
		"nearestBuilding": trashcan.NearestBuilding,
	})

	if err != nil {
		log.Printf("An error has occurred to create trashcan: %s", err)
	}

	return err
}
