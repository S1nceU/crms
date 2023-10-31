package http

import (
	"crms/module/customer"
	"fmt"
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
	e.POST("/api/customer", handler.GetCustomer)
}

// GetCustomerList @Summary GetCustomerList
// @Description Get all Customer
// @Accept json
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

// GetCustomer @Summary GetCustomer
// @Description Get Customer by ID
// @Tags Customer
// @Produce application/json
// @Param ID formData string true "Customer ID"
// @Success 200 {object} model.Customer
// @Failure 500 {string} string "{"Message": err.Error()}"
// @Router /customer [post]
func (u *CustomerHandler) GetCustomer(c *gin.Context) {
	ID := c.PostForm("ID")
	fmt.Println(ID)
	customerData, err := u.ser.GetCustomer(ID)
	if err != nil {
		if err.Error() == "error CRMS : There is no this customer" {
			c.JSON(210, gin.H{
				"Message": err.Error(),
			})
			return
		} else {
			c.JSON(500, gin.H{
				"Message": err.Error(),
			})
			return
		}
	}
	c.JSON(200, customerData)
}
