package repository

import (
	"go.mod/src/crms/model"
	"go.mod/src/crms/module/history"
	"gorm.io/gorm"
)

type HistoryRepository struct {
	orm *gorm.DB
}

func NewHistoryRepository(orm *gorm.DB) history.Repository {
	return &HistoryRepository{
		orm: orm,
	}
}

func (u *HistoryRepository) GetHistoryList(map[string]interface{}) ([]*model.History, error) {
	var err error
	var in = make([]*model.History, 0)
	err = u.orm.First(&in).Error
	return in, err
}

func (u *HistoryRepository) GetHistory(in *model.History) (*model.History, error) {
	var err error
	err = u.orm.First(&in).Error
	return in, err
}

func (u *HistoryRepository) CreateHistory(in *model.History) (*model.History, error) {
	var err error
	err = u.orm.Create(&in).Error
	return in, err
}

func (u *HistoryRepository) UpdateHistory(in *model.History) (*model.History, error) {
	var err error
	err = u.orm.Save(&in).Error
	return in, err
}

func (u *HistoryRepository) DeleteHistory(in *model.History) error {
	return u.orm.Delete(&in).Error
}
