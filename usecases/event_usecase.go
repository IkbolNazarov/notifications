package usecases

import (
	"notifications/entities"
	"notifications/repository"
)

type EventUsecase interface {
	AddEvent(event *entities.Event) error
	GetPendingEvents() ([]*entities.Event, error)
	RemoveEvent(event *entities.Event) error
}

type eventUsecase struct {
	eventRepo repository.EventRepository
}

func NewEventUsecase(eventRepo repository.EventRepository) EventUsecase {
	return &eventUsecase{
		eventRepo: eventRepo,
	}
}

func (u *eventUsecase) AddEvent(event *entities.Event) error {
	return u.eventRepo.Save(event)
}

func (u *eventUsecase) GetPendingEvents() ([]*entities.Event, error) {
	return u.eventRepo.GetAll()
}

func (u *eventUsecase) RemoveEvent(event *entities.Event) error {
	return u.eventRepo.Remove(event)
}
