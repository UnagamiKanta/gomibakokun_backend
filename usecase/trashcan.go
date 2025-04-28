package usecase

import (
	"context"
	"gomibakokun_backend/domain"
	"gomibakokun_backend/domain/repository"

	"github.com/google/uuid"
)

type TrashcanUseCase interface {
	CreateTrashcan(ctx context.Context, latitude float64, longitude float64, image string, trashType []string) error
}

type trashcanUseCase struct {
	trashcanRepository repository.TrashcanRepository
}

func NewTrashcanUseCase(tr repository.TrashcanRepository) TrashcanUseCase {
	return &trashcanUseCase{
		trashcanRepository: tr,
	}
}

func (tu *trashcanUseCase) CreateTrashcan(ctx context.Context, latitude float64, longitude float64, image string, trashType []string) error {
	ID, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	trashcan := domain.Trashcan{
		ID:        ID.String(),
		Latitude:  latitude,
		Longitude: longitude,
		Image:     image,
		TrashType: trashType,
	}

	err = tu.trashcanRepository.CreateTrashcan(ctx, &trashcan)

	return err
}
