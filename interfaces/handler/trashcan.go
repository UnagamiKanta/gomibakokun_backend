package handler

import (
	"gomibakokun_backend/interfaces/response"
	"gomibakokun_backend/usecase"
	"net/http"

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
		return c.JSON(http.StatusInternalServerError, echo.Map{"success": false})
	}

	ctx := c.Request().Context()

	err := th.trashcanUsecase.CreateTrashcan(ctx, req.Latitude, req.Longitude, req.Image, req.TrashType, req.NearestBuilding)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"success": false})
	}

	return c.JSON(http.StatusCreated, echo.Map{"success": true})
}
