package repository

import (
	"notifications/config"
	"notifications/db"
	"notifications/entities"
	"testing"
	"time"
)

func TestGORMEventRepository(t *testing.T) {
	dbConfig := &config.DatabaseConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "testuser",
		Password: "testpassword",
		DBName:   "testdb",
		SSLMode:  "disable",
	}

	db, err := db.ConnectDatabase(dbConfig)
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&entities.Event{}); err != nil {
		t.Fatalf("Failed to migrate database: %v", err)
	}

	db.Exec("DELETE FROM events")

	repo := NewEventRepository(db)
	event := &entities.Event{
		OrderType:  "test",
		SessionID:  "test session",
		Card:       "11111111",
		EventDate:  time.Now(),
		WebsiteURL: "https://somon.tj",
	}

	if err := repo.Save(event); err != nil {
		t.Fatalf("Failed to save event: %v", err)
	}

	events, err := repo.GetAll()
	if err != nil {
		t.Fatalf("Failed to get events: %v", err)
	}
	if len(events) != 1 {
		t.Fatalf("Expected 1 event, got %d", len(events))
	}

	if err := repo.Remove(event); err != nil {
		t.Fatalf("Failed to remove event: %v", err)
	}
	events, err = repo.GetAll()
	if err != nil {
		t.Fatalf("Failed to get events: %v", err)
	}
	if len(events) != 0 {
		t.Fatalf("Expected 0 events, got %d", len(events))
	}
}
