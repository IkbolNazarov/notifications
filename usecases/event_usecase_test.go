package usecases

import (
	"notifications/entities"
	"notifications/repository"
	"testing"
	"time"
)

func TestEventUsecase(t *testing.T) {
	repo := repository.NewInMemoryEventRepository()
	usecases := NewEventRepository(repo)

	event := &entities.Event{
		OrderType:  "Purchase",
		SessionID:  "session123",
		Card:       "1234**5678",
		EventDate:  time.Now(),
		WebsiteURL: "https://example.com",
	}

	if err := usecases.AddEvent(event); err != nil {
		t.Fatalf("Failed to add event: %v", err)
	}

	events, err := usecases.GetPendingEvents()
	if err != nil {
		t.Fatalf("Failed to get pending events: %v", err)
	}
	if len(events) != 1 {
		t.Fatalf("Expected 1 event, got %d", len(events))
	}

	if err := usecases.RemoveEvent(event); err != nil {
		t.Fatalf("Failed to remove event: %v", err)
	}
	events, err = usecases.GetPendingEvents()
	if err != nil {
		t.Fatalf("Failed to get pending events: %v", err)
	}
	if len(events) != 0 {
		t.Fatalf("Expected 0 events, got %d", len(events))
	}
}
