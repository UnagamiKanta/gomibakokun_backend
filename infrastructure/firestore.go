package firestore

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

func InitFirestoreClient(ctx context.Context, projectID string) (*firestore.Client, error) {
	credJson := []byte(os.Getenv("GOOGLE_CREDENTIALS_JSON"))
	if len(credJson) == 0 {
		log.Println("GOOGLE_CREDENTIALS_JSON is not set")
	}

	client, err := firestore.NewClient(ctx, projectID, option.WithCredentialsJSON(credJson))
	if err != nil {
		log.Fatalf("failed to initialize Firestore client: %v", err)
		return nil, err
	}
	return client, nil
}
