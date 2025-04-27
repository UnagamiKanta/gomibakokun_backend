package infra

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

func InitFirestoreClient(ctx context.Context, projectID string) *firestore.Client {
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("failed to initialize Firestore client: %v", err)
	}
	return client
}
