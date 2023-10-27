package main

import (
	"crms/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
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
	server.GET("/", hello)
	err := server.Run(":8080")
	if err != nil {
		return
	}
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Hello": "True",
	})
}
