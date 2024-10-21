package handler

import (
	"encoding/json"
	"net/http"
	"notifications/entities"
	"notifications/usecases"
)

type EventHandler struct {
	EventUsecase usecases.EventUsecase
}

func NewEventHandler(eventUsecase usecases.EventUsecase) *EventHandler {
	return &EventHandler{
		EventUsecase: eventUsecase,
	}
}

func (h *EventHandler) HandleEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var event entities.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "wrong request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if event.EventDate.IsZero() {
		http.Error(w, "wrong request", http.StatusBadRequest)
		return
	}

	if err := h.EventUsecase.AddEvent(&event); err != nil {
		http.Error(w, "error adding event", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
