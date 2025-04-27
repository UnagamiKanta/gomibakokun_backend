package repository

import (
	"context"
	"gomibakokun_backend/domain"

	"cloud.google.com/go/firestore"
)

type TrashcanRepository interface {
	CreateTrashcan(ctx context.Context, client *firestore.Client, trashcan *domain.Trashcan) error
}
