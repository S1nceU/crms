package history

import "go.mod/src/crms/model"

type Repository interface {
	GetHistoryList() ([]*model.History, error)
	GetHistory(in *model.History) (*model.History, error)
	CreateHistory(in *model.History) (*model.History, error)
	UpdateHistory(in *model.History) (*model.History, error)
	DeleteHistory(in *model.History) error
}
