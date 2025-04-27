package usecase

import (
	"context"
	"gomibakokun_backend/domain"
	"gomibakokun_backend/domain/repository"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
)

type TrashcanUseCase interface {
	CreateTrashcan(ctx context.Context, client *firestore.Client, latitude float64, longitude float64, nearestBuilding string, trashType []string) error
}

type trashcanUseCase struct {
	trashcanRepository repository.TrashcanRepository
}

func NewTrashcanUseCase(tr repository.TrashcanRepository) TrashcanUseCase {
	return &trashcanUseCase{
		trashcanRepository: tr,
	}
}

func (tu *trashcanUseCase) CreateTrashcan(ctx context.Context, client *firestore.Client, latitude float64, longitude float64, nearestBuilding string, trashType []string) error {
	ID, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	trashcan := domain.Trashcan{
		ID:              ID.String(),
		Latitude:        latitude,
		Longitude:       longitude,
		NearestBuilding: nearestBuilding,
		TrashType:       trashType,
	}

	err = tu.trashcanRepository.CreateTrashcan(ctx, client, &trashcan)

	return err
}
