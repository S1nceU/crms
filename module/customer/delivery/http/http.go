package http

import (
	"crms/module/customer"
	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	ser customer.Service
}

func NewCustomerHandler(e *gin.Engine, ser customer.Service) {
	handler := &CustomerHandler{
		ser: ser,
	}
	e.GET("/api/customerList", handler.GetCustomerList)
}

// GetCustomerList @Summary GetCustomerList
// @Description Get all Customer
// @Tags Customer
// @Produce application/json
// @Success 200 {object} model.Customer
// @Failure 500 {string} string "{"Message": "Internal Error!"}"
// @Router /customerList [get]
func (u *CustomerHandler) GetCustomerList(c *gin.Context) {
	customerList, err := u.ser.GetCustomerList()
	if err != nil {
		c.JSON(500, gin.H{
			"Message": "Internal Error!",
		})
		return
	}
	c.JSON(200, customerList)
}
