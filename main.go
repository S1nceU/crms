package main

import (
	"fmt"
	"go.mod/src/crms/model"
	cr "go.mod/src/crms/module/customer/repository"
	cs "go.mod/src/crms/module/customer/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	USERNAME = "root"
	PASSWORD = "xu.6j03cj86u;6au/65k6"
	NETWORK  = "tcp"
	SERVER   = "127.0.0.2"
	PORT     = 3306
	DATABASE = "crms_sql"
)

func main() {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	var (
		db    *gorm.DB
		dbErr error
	)
	db, dbErr = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		panic("使用 gorm 連線 DB 發生錯誤，原因為 " + dbErr.Error())
	} else {
		fmt.Println("連線成功")
		_ = db
		//_ = customerRepo
	}
	var (
		customerRepo = cr.NewCustomerRepository(db)
		//hisRepo      = hr.NewHistoryRepository(db)
		customerSer = cs.NewCustomer(customerRepo)
		//hisSer       = cs.NewCustomer(hisRepo)
	)
	//_ = hisSer

	newC := &model.Customer{
		//Customer_id :,
		Name:        "",
		Gender:      "",
		Birthday:    "",
		ID:          "",
		Address:     "",
		Phonenumber: "",
		Carnumber:   "",
		Citizenship: "Taipei",
		Note:        "",
	}

	newH := &model.History{
		//History_id  int    `json:"history_id"  gorm:"primary_key;auto_increase;not null"`
		Customer_id: 5,
		Date:        "8/25",
		Nofpeople:   1,
		Price:       200,
		//Note:        "",
	}
	_ = newH
	_ = newC

	if point, err := customerSer.GetCustomerForCID(5); err != nil {
		//panic("錯誤 :" + err.Error())
		fmt.Println("錯誤: " + err.Error())
	} else {
		fmt.Println("你為什麼會動\n", point)
	}
}
