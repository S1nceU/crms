package service

import (
	"go.mod/src/crms/model"
	"go.mod/src/crms/module/history"
)

type HistoryService struct {
	repo history.Repository
}

func NewHistory(repo history.Repository) history.Service {
	return &HistoryService{
		repo: repo,
	}
}

func (u *HistoryService) GetHistoryList(data map[string]interface{}) ([]*model.History, error) {
	return u.repo.GetHistoryList(data)
}
func (u *HistoryService) GetHistory(in *model.History) (*model.History, error) {
	return u.repo.GetHistory(in)
}
func (u *HistoryService) CreateHistory(in *model.History) (*model.History, error) {
	return u.repo.CreateHistory(in)
}
func (u *HistoryService) UpdateHistory(in *model.History) (*model.History, error) {
	return u.repo.UpdateHistory(in)
}
func (u *HistoryService) DeleteHistory(in *model.History) error {
	return u.repo.DeleteHistory(in)
}
