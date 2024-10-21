package main

import (
	"fmt"
	"net/http"
	"notifications/config"
	"notifications/db"
	"notifications/entities"
	"notifications/handler"
	"notifications/repository"
	"notifications/usecases"
	"notifications/worker"
	"time"
)

func main() {
	dbConfig := config.NewDatabaseConfig()

	db, err := db.ConnectDatabase(dbConfig)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	if err := db.AutoMigrate(&entities.Event{}); err != nil {
		panic(fmt.Sprintf("Failed to migrate database: %v", err))
	}

	eventRepo := repository.NewEventRepository(db)
	eventUsecase := usecases.NewEventUsecase(eventRepo)
	eventHandler := handler.NewEventHandler(eventUsecase)

	http.HandleFunc("/events", eventHandler.HandleEvent)

	eventWorker := worker.NewWorker(eventUsecase, 5*time.Second)
	eventWorker.Start()
	defer eventWorker.Stop()

	port := ":8080"
	fmt.Println("Server is running on port", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
