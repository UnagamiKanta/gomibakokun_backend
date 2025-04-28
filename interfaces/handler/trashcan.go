package interfaces

import (
	"gomibakokun_backend/interfaces/response"
	"gomibakokun_backend/usecase"

	"github.com/labstack/echo/v4"
)

type TrashcanHandler interface {
	HandleTrashcanCreate(c echo.Context) error
}

type trashcanHandler struct {
	trashcanUsecase usecase.TrashcanUseCase
}

func NewTrashcanHandler(tu usecase.TrashcanUseCase) TrashcanHandler {
	return &trashcanHandler{
		trashcanUsecase: tu,
	}
}

func (th trashcanHandler) HandleTrashcanCreate(c echo.Context) error {
	var req response.CreateTrashcanReq
	if err := c.Bind(&req); err != nil {
		return err
	}

	ctx := c.Request().Context()

	err := usecase.TrashcanUseCase{}.CreateTrashcan(ctx, req.Latitude, req.Longitude, req.Image, req.TrashType)
}
