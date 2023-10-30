package delivery

import "github.com/gin-gonic/gin"

type HttpHandler interface {
	GetCustomerList(c *gin.Context)
	//GetUser(c *gin.Context) error
	//CreateUser(c *gin.Context) error
	//ModifyUser(c *gin.Context) error
	//DeleteUser(c *gin.Context) error
}
