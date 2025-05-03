package repository

import (
	"context"
	"gomibakokun_backend/domain"
)

type TrashcanRepository interface {
	CreateTrashcan(ctx context.Context, trashcan *domain.Trashcan) error
	GetAllTrashcan(ctx context.Context) ([]*domain.Trashcan, error)
	DeleteTrashcan(ctx context.Context, ID string) error
}
