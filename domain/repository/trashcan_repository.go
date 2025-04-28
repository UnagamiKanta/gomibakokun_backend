package repository

import (
	"context"
	"gomibakokun_backend/domain"
)

type TrashcanRepository interface {
	CreateTrashcan(ctx context.Context, trashcan *domain.Trashcan) error
}
