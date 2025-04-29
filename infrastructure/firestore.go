package firestore

import (
	"context"
	"encoding/base64"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

func InitFirestoreClient(ctx context.Context, projectID string) (*firestore.Client, error) {
	b64 := os.Getenv("GOOGLE_CREDENTIALS_JSON_BASE64")
	if b64 == "" {
		log.Fatal("GOOGLE_CREDENTIALS_JSON_BASE64 is not set")
	}

	dec, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		log.Fatalf("failed to decode service key: %v", err)
	}
	creds := option.WithCredentialsJSON(dec)

	client, err := firestore.NewClient(ctx, projectID, creds)
	if err != nil {
		log.Fatalf("failed to create Firestore client: %v", err)
	}
	return client, nil
}
