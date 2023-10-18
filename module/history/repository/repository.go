package repository

import (
	"crms/model"
	"crms/module/history"
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

func (u *HistoryRepository) GetHistoryList() ([]*model.History, error) {
	var err error
	var in []*model.History
	err = u.orm.Find(&in).Error
	return in, err
}

func (u *HistoryRepository) GetHistory(in *model.History) ([]*model.History, error) {
	var err error
	var out []*model.History
	err = u.orm.Where("customer_id = ?", in.Customer_id).Find(&out).Error
	return out, err
}

func (u *HistoryRepository) GetHistoryForDate(in *model.History) ([]*model.History, error) {
	var err error
	var out []*model.History
	err = u.orm.Where("date = ?", in.Date).Find(&out).Error
	return out, err
}

func (u *HistoryRepository) GetHistoryForHistoryId(in int) (*model.History, error) {
	var err error
	var out *model.History
	err = u.orm.Where("history_id = ?", in).First(&out).Error
	return out, err
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

func (u *HistoryRepository) DeleteHistory(in int) error {
	var out *model.History
	var err error
	err = u.orm.Where("history_id = ?", in).Delete(&out).Error
	return err
}
