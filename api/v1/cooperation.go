package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wejectchen/ginblog/model"
	"github.com/wejectchen/ginblog/utils/errmsg"
)

// AddCooperation 添加合作模式
func AddCooperation(c *gin.Context) {
	var data model.Cooperation
	_ = c.ShouldBindJSON(&data)

	code, err := model.CreateCooperation(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// GetCooperation 查看合作模式
func GetCooperation(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

	data, total := model.GetCooperation(page, size)
	code := errmsg.SUCCESS
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// EditCooperation 编辑合作模式
func EditCooperation(c *gin.Context) {
	var data model.Cooperation
	_ = c.ShouldBindJSON(&data)
	id := data.ID

	code, err := model.EditCooperation(id, &data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if code == errmsg.ERROR_CATENAME_USED {
		c.Abort()
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// DeleteCooperation 合作模式
func DeleteCooperation(c *gin.Context) {
	var data model.Cooperation
	_ = c.ShouldBindJSON(&data)
	id := data.ID

	code, err := model.DeleteCooperation(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}
