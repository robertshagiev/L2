package usecase

import "11/model"

func getStorageEvent(event *model.Event) *model.Storage {
	return &model.Storage{
		Title: event.Title,
		Info:  event.Info,
		Date:  event.Date,
	}
}
