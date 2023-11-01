package http

import (
	"crms/model"
	"crms/module/customer"
	"github.com/gin-gonic/gin"
	"strconv"
)

type CustomerHandler struct {
	ser customer.Service
}

func NewCustomerHandler(e *gin.Engine, ser customer.Service) {
	handler := &CustomerHandler{
		ser: ser,
	}
	e.GET("/api/customerList", handler.GetCustomerList)
	e.GET("/api/customer", handler.GetCustomer)
	e.POST("/api/customer", handler.CreateCustomer)
	e.PUT("/api/customer", handler.ModifyCustomer)
	e.DELETE("/api/customer", handler.DeleteCustomer)
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
// @Param ID query string true "Customer ID" example(L123546789)
// @Success 200 {object} model.Customer
// @Failure 500 {string} string "{"Message": err.Error()}"
// @Router /customer [get]
func (u *CustomerHandler) GetCustomer(c *gin.Context) {
	ID := c.Query("ID")
	customerData, err := u.ser.GetCustomer(ID)
	if err != nil {
		if err.Error() == "error CRMS : There is no this customer" {
			c.JSON(200, gin.H{
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

// CreateCustomer @Summary CreateCustomer
// @Description Create Customer
// @Tags Customer
// @Accept json
// @Produce application/json
// @Param Customer body model.Customer true "Customer Information"
// @Success 200 {object} model.Customer
// @Failure 500 {string} string "{"Message": err.Error()}"
// @Router /customer [post]
func (u *CustomerHandler) CreateCustomer(c *gin.Context) {
	json := model.Customer{}
	if err := c.BindJSON(&json); err != nil {
		c.JSON(400, gin.H{
			"Message": err.Error(),
		})
		return
	}
	createCustomer, err := u.ser.CreateCustomer(&json)
	if err != nil {
		if err.Error() == "error CRMS : This customer is already existed" {
			c.JSON(200, gin.H{
				"Message": err.Error(),
			})
			return
		} else if err.Error() == "error CRMS : Customer Info is incomplete" {
			c.JSON(200, gin.H{
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

	c.JSON(200, createCustomer)
}

// ModifyCustomer @Summary ModifyCustomer
// @Description Modify Customer
// @Tags Customer
// @Accept json
// @Produce application/json
// @Param Customer body model.Customer true "Customer Information"
// @Success 200 {object} model.Customer
// @Failure 500 {string} string "{"Message": err.Error()}"
// @Router /customer [put]
func (u *CustomerHandler) ModifyCustomer(c *gin.Context) {
	json := model.Customer{}
	if err := c.BindJSON(&json); err != nil {
		c.JSON(400, gin.H{
			"Message": err.Error(),
		})
		return
	}
	modifyCustomer, err := u.ser.UpdateCustomer(&json)
	if err != nil {
		if err.Error() == "error CRMS : There is no this customer" {
			c.JSON(200, gin.H{
				"Message": err.Error(),
			})
			return
		} else if err.Error() == "error CRMS : Customer Info is incomplete" {
			c.JSON(200, gin.H{
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
	c.JSON(200, gin.H{
		"Customer info": modifyCustomer,
		"Message":       "Modify success",
	})
}

// DeleteCustomer @Summary DeleteCustomer
// @Description DeleteCustomer by  CustomerId
// @Tags Customer
// @Produce application/json
// @Param CustomerId query string true "Customer Id"
// @Success 200 {object} string "Message": "Delete success"
// @Failure 500 {string} string "{"Message": err.Error()}"
// @Router /customer [delete]
func (u *CustomerHandler) DeleteCustomer(c *gin.Context) {
	CustomerId, _ := strconv.Atoi(c.Query("CustomerId"))
	err := u.ser.DeleteCustomer(CustomerId)
	if err != nil {
		if err.Error() == "error CRMS : There is no this customer" {
			c.JSON(200, gin.H{
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
	c.JSON(200, gin.H{
		"Message": "Delete success",
	})
}
