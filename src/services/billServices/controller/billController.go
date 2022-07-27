package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BillRoute struct {
}

// GetBillsYearAllDataREPost
// @Tags Bill
// @Summary 获取最近一年的bills
// @Description 描述信息
// @Security Bearer
// @Produce  json
// @Success 200 {object} []billModels.BillDetail
// @Router /v1/Bill/GetBillsYearAllData [post]
func (BillRoute) GetBillsYearAllDataREPost(c *gin.Context) {

}

func (BillRoute) GetBillsYearAllDataREGet(c *gin.Context) {
	c.JSONP(http.StatusBadGateway, gin.H{"ex": "error"})
}
