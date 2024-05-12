package usecase

import (
	"11/model"
	"11/repository"
	"time"
)

type Usecase struct {
	Repository *repository.Repository
}

func NewUsecase(repository *repository.Repository) *Usecase {
	return &Usecase{
		Repository: repository,
	}
}

func (u *Usecase) CreateEvent(event *model.Event) error {
	storage := getStorageEvent(event)
	return u.Repository.AddEvent(event.UserID, storage)
}

func (u *Usecase) UpdateEvent(event *model.Event) error {
	storage := getStorageEvent(event)
	return u.Repository.UpdateEvent(event.UserID, storage)
}

func (u *Usecase) DeleteEvent(event *model.Event) error {
	return u.Repository.DeleteEvent(event.UserID)
}

func (u *Usecase) GetEventsForDay(userID int, date time.Time) ([]*model.Storage, error) {
	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	endOfDay := startOfDay.AddDate(0, 0, 1).Add(-time.Nanosecond)

	return u.Repository.GetEventsForPeriod(userID, startOfDay, endOfDay)
}

func (u *Usecase) GetEventsForWeek(userID int, date time.Time) ([]*model.Storage, error) {
	weekday := date.Weekday()
	offset := int(time.Monday - weekday)
	if offset > 0 {
		offset -= 7
	}
	startOfWeek := date.AddDate(0, 0, offset)
	endOfWeek := startOfWeek.AddDate(0, 0, 7).Add(-time.Nanosecond)
	return u.Repository.GetEventsForPeriod(userID, startOfWeek, endOfWeek)
}

func (u *Usecase) GetEventsForMonth(userID int, date time.Time) ([]*model.Storage, error) {
	startOfMonth := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())
	endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-time.Nanosecond)
	return u.Repository.GetEventsForPeriod(userID, startOfMonth, endOfMonth)
}
