package history

import "crms/model"

type Service interface {
	GetHistoryList() ([]model.History, error)
	GetHistory(in int) ([]model.History, error)
	GetHistoryForDate(in string) ([]model.History, error)
	GetHistoryForHId(in int) (*model.History, error)
	CreateHistory(in []byte) (*model.History, error)
	UpdateHistory(in []byte) (*model.History, error)
	DeleteHistory(in int) error
}
