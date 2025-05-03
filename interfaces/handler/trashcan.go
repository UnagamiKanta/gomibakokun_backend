package handler

import (
	"gomibakokun_backend/interfaces/response"
	"gomibakokun_backend/usecase"
	"net/http"

	"strconv"

	"github.com/labstack/echo/v4"
)

type TrashcanHandler interface {
	HandleTrashcanCreate(c echo.Context) error
	HandleTrashcansInRange(c echo.Context) error
	HandleTrashcanDelete(c echo.Context) error
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

func (th trashcanHandler) HandleTrashcansInRange(c echo.Context) error {
	latitude := c.QueryParam("latitude")
	longitude := c.QueryParam("longitude")

	ctx := c.Request().Context()

	range_radius := 20000 //TODO:リリース時治す

	latitudeFloat, err := strconv.ParseFloat(latitude, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"success": false})
	}

	longitudeFloat, err := strconv.ParseFloat(longitude, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"success": false})
	}

	trashcans, err := th.trashcanUsecase.GetTrashcansInRange(ctx, latitudeFloat, longitudeFloat, float64(range_radius))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"success": false})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": true, "trashcans": trashcans})
}

func (th trashcanHandler) HandleTrashcanDelete(c echo.Context) error {
	ID := c.QueryParam("id")

	ctx := c.Request().Context()

	err := th.trashcanUsecase.DeleteTrashcan(ctx, ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"success": false})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": true})
}
