package history

import "github.com/S1nceU/CRMS/model"

type Service interface {
	GetHistoryList() ([]model.History, error)
	GetHistory(in int) ([]model.History, error)
	GetHistoryForDate(in string) ([]model.History, error)
	GetHistoryForHId(in int) (*model.History, error)
	CreateHistory(in *model.History) (*model.History, error)
	UpdateHistory(in *model.History) (*model.History, error)
	DeleteHistory(in int) error
}
