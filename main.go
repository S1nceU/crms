package main

import (
	"crms/model"
	cr "crms/module/customer/repository"
	cs "crms/module/customer/service"
	hr "crms/module/history/repository"
	hs "crms/module/history/service"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	USERNAME = "root"
	PASSWORD = "password"
	NETWORK  = "tcp"
	SERVER   = "127.0.0.2"
	PORT     = 3306
	DATABASE = "crms"
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
		var err error
		if err = db.AutoMigrate(&model.Customer{}); err != nil {
			return
		}
		if err = db.AutoMigrate(&model.History{}); err != nil {
			return
		}

		// defer db.close 似乎已被 gorm 刪掉功能
	}
	var (
		customerRepo = cr.NewCustomerRepository(db)
		hisRepo      = hr.NewHistoryRepository(db)
		customerSer  = cs.NewCustomer(customerRepo)
		hisSer       = hs.NewHistory(hisRepo)
	)

	newC := &model.Customer{
		//Customer_id :,
		Name:        "Jason",
		Gender:      "Male",
		Birthday:    "2001/09/25",
		ID:          "L123456789",
		Address:     "Taichung",
		Phonenumber: "0987654321",
		Carnumber:   "",
		Citizenship: "Taiwan",
		Note:        "",
	}

	newH := &model.History{
		//History_id
		CustomerId: 5,
		Date:       "8/25",
		Nofpeople:  1,
		Price:      200,
		//Note:        "",
	}
	_ = newH
	_ = newC
	_ = customerRepo
	_ = customerSer
	_ = hisSer
	_ = hisRepo

	inputJson := []byte(`{"CustomerId":2,"Name":"John", "Gender":"male", "Birthday":"2001/09/26", "ID":"A123456700", "Citizenship":"Taiwan", "Address":"Taichung"}`)
	_ = inputJson

	if point, err := hisSer.GetHistoryForHId(2); err != nil {
		fmt.Println("錯誤: " + err.Error())
	} else {
		fmt.Println("你為什麼會動\n", point)
	}
	//if point, err := customerSer.GetCustomerForCID(2); err != nil {
	//	fmt.Println("錯誤: " + err.Error())
	//} else {
	//	fmt.Println("你為什麼會動\n", point)
	//}

}
