package history

import "crms/model"

type Repository interface {
	GetHistoryList() ([]*model.History, error)
	GetHistory(in *model.History) ([]*model.History, error)
	GetHistoryForDate(in *model.History) ([]*model.History, error)
	GetHistoryForHistoryId(in int) (*model.History, error)
	CreateHistory(in *model.History) (*model.History, error)
	UpdateHistory(in *model.History) (*model.History, error)
	DeleteHistory(in int) error
}
