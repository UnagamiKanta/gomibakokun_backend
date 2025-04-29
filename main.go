package main

import (
	firestore "gomibakokun_backend/infrastructure"
	"gomibakokun_backend/infrastructure/persistence"
	"gomibakokun_backend/interfaces/handler"
	"gomibakokun_backend/usecase"

	"context"
	"log"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	ctx := context.Background()
	client, err := firestore.InitFirestoreClient(ctx, "gomibakokun")
	if err != nil {
		log.Fatalf("failed to initialize Firestore client: %v", err)
	}

	trashcanPersistence := persistence.NewTrashcanPersistence(client)
	trashcanUseCase := usecase.NewTrashcanUseCase(trashcanPersistence)
	trashcanHandler := handler.NewTrashcanHandler(trashcanUseCase)

	e := echo.New()
	e.POST("/trashcan", trashcanHandler.HandleTrashcanCreate)
	e.Start(":" + port)
}
