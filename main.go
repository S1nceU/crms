package main

import (
	_ "crms/docs"
	"crms/model"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_customerHandlerHttpDelivery "crms/module/customer/delivery/http"
	_customerRepo "crms/module/customer/repository"
	_customerSer "crms/module/customer/service"

	_historyHandlerHttpDelivery "crms/module/history/delivery/http"
	_historyRepo "crms/module/history/repository"
	_historySer "crms/module/history/service"
)

const (
	USERNAME = "root"
	PASSWORD = "password"
	NETWORK  = "tcp"
	SERVER   = "127.0.0.2"
	PORT     = 3306
	DATABASE = "crms"
)

var swagHandler gin.HandlerFunc

// @title CRMS_Swagger
// @version 1.0
// @description CRMS_Swagger information
// @termsOfService http://www.google.com

// @contact.name Jason Yang
// @contact.url http://www.google.com
// @contact.email jjkk900925@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath /api

func init() {
	swagHandler = ginSwagger.WrapHandler(swaggerFiles.Handler)
}

func main() {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	var (
		db    *gorm.DB
		dbErr error
	)
	if db, dbErr = gorm.Open(mysql.Open(dsn), &gorm.Config{}); dbErr != nil {
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
	}

	server := gin.Default()

	customerRepo := _customerRepo.NewCustomerRepository(db)
	customerSer := _customerSer.NewCustomerService(customerRepo)
	_customerHandlerHttpDelivery.NewCustomerHandler(server, customerSer)

	historyRepo := _historyRepo.NewHistoryRepository(db)
	historySer := _historySer.NewHistoryService(historyRepo)
	_historyHandlerHttpDelivery.NewHistoryHandler(server, historySer)

	if swagHandler != nil {
		server.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	err := server.Run(":8080")
	if err != nil {
		return
	}
}
