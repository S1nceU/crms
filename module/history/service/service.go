package service

import (
	"crms/model"
	"crms/module/history"
	"errors"
)

type HistoryService struct {
	repo history.Repository
}

func NewHistory(repo history.Repository) history.Service {
	return &HistoryService{
		repo: repo,
	}
}

func (u *HistoryService) GetHistoryList() ([]model.History, error) {
	var err error
	var point []*model.History
	var out []model.History
	point, err = u.repo.GetHistoryList()
	for i := 0; i < len(point); i++ {
		out = append(out, *point[i])
	}
	return out, err
}

func (u *HistoryService) GetHistory(in int) ([]model.History, error) {
	var err error
	var point []*model.History
	var out []model.History
	newHistory := &model.History{
		CustomerId: in,
	}
	if point, err = u.repo.GetHistory(newHistory); len(point) == 0 {
		return nil, errors.New("error CRMS : There is no this customer or no history")
	}
	for i := 0; i < len(point); i++ {
		out = append(out, *point[i])
	}
	return out, err
}

func (u *HistoryService) GetHistoryForDate(in string) ([]model.History, error) {
	var err error
	var point []*model.History
	var out []model.History
	newHistory := &model.History{
		Date: in,
	}
	if point, err = u.repo.GetHistoryForDate(newHistory); len(point) == 0 {
		return nil, errors.New("error CRMS : There was no customer in " + in)
	}
	for i := 0; i < len(point); i++ {
		out = append(out, *point[i])
	}
	return out, err
}

func (u *HistoryService) GetHistoryForHId(in int) (*model.History, error) {
	var err error
	newHistory := &model.History{
		HistoryId: in,
	}
	if newHistory, err = u.repo.GetHistoryForHId(newHistory); newHistory.CustomerId == 0 {
		return nil, errors.New("error CRMS : There is no this history")
	}
	return newHistory, err
}

//func (u *HistoryService) CreateHistory(in *model.History) (*model.History, error) {
//	return u.repo.CreateHistory(in)
//}
//func (u *HistoryService) UpdateHistory(in *model.History) (*model.History, error) {
//	return u.repo.UpdateHistory(in)
//}
//func (u *HistoryService) DeleteHistory(in *model.History) error {
//	return u.repo.DeleteHistory(in)
//}
