package repository

import (
	"gorm.io/gorm"
	"notifications/entities"
)

type EventRepository interface {
	Save(event *entities.Event) error
	GetAll() ([]*entities.Event, error)
	Remove(event *entities.Event) error
}
type GORMEventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *GORMEventRepository {
	return &GORMEventRepository{
		db: db,
	}
}

func (repo *GORMEventRepository) Save(event *entities.Event) error {
	return repo.db.Create(event).Error
}

func (repo *GORMEventRepository) GetAll() ([]*entities.Event, error) {
	var events []*entities.Event
	err := repo.db.Find(&events).Error
	return events, err
}

func (repo *GORMEventRepository) Remove(event *entities.Event) error {
	return repo.db.Delete(event).Error
}
