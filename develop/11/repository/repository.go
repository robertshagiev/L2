package repository

import (
	"11/model"
	"errors"
	"sync"
	"time"
)

type Repository struct {
	mutex  sync.Mutex
	events map[int][]*model.Storage
}

func NewRepository() *Repository {
	return &Repository{
		events: make(map[int][]*model.Storage),
	}
}

func (r *Repository) AddEvent(userID int, storage *model.Storage) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.events[userID] = append(r.events[userID], storage)
	return nil
}

func (r *Repository) UpdateEvent(userID int, storage *model.Storage) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	events, exists := r.events[userID]
	if !exists {
		return errors.New("no events found for this user")
	}

	for i, existingEvent := range events {
		if existingEvent.Date == storage.Date {
			r.events[userID][i] = storage
			return nil
		}
	}

	return errors.New("event not found")
}

func (r *Repository) DeleteEvent(userID int) error {
	if _, exists := r.events[userID]; !exists {
		return errors.New("no events found for this user")
	}
	delete(r.events, userID)
	return nil

}

func (r *Repository) GetEventsForPeriod(userID int, start, end time.Time) ([]*model.Storage, error) {
	var result []*model.Storage
	events, exists := r.events[userID]
	if !exists {
		return nil, errors.New("no events found for this user")
	}
	for _, event := range events {
		if (event.Date.After(start) || event.Date.Equal(start)) && (event.Date.Before(end) || event.Date.Equal(end)) {
			result = append(result, event)
		}
	}
	return result, nil
}
