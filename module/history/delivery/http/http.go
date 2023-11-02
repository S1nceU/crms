package http

import (
	"crms/module/history"
	"github.com/gin-gonic/gin"
	"strconv"
)

type HistoryHandler struct {
	ser history.Service
}

func NewHistoryHandler(e *gin.Engine, ser history.Service) {
	handler := &HistoryHandler{
		ser: ser,
	}
	e.GET("/api/historyList", handler.GetHistoryList)
	e.GET("/api/history", handler.GetHistory)
}

// GetHistoryList @Summary GetHistoryList
// @Description Get all History
// @Accept json
// @Tags History
// @Produce application/json
// @Success 200 {object} model.History
// @Failure 500 {string} string "{"Message": "Internal Error!"}"
// @Router /historyList [get]
func (u *HistoryHandler) GetHistoryList(c *gin.Context) {
	historyList, err := u.ser.GetHistoryList()
	if err != nil {
		c.JSON(500, gin.H{
			"Message": "Internal Error!",
		})
		return
	}
	c.JSON(200, historyList)
}

// GetHistory @Summary GetHistory
// @Description Get History by CustomerId
// @Tags History
// @Produce application/json
// @Param CustomerId query string true "Customer Id"
// @Success 200 {object} model.History
// @Failure 500 {string} string "{"Message": err.Error()}"
// @Router /history [get]
func (u *HistoryHandler) GetHistory(c *gin.Context) {
	customerId, _ := strconv.Atoi(c.Query("CustomerId"))
	historyData, err := u.ser.GetHistory(customerId)
	if err != nil {
		if err.Error() == "error CRMS : There is no this customer or no history" {
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
	c.JSON(200, historyData)
}
