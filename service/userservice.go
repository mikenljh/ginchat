package service

import (
	"ginchat/models"

	"github.com/gin-gonic/gin"
)

// GetIndex
// @Tags        首页
// @Summary     用户页面
// @Description 返回所有用户
// @Success      200 {string} json{"code","message"}
// @Router      /user [get]
func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()

	c.JSON(200, gin.H{
		"code":    0,
		"message": "success",
		"data":    data,
	})

}
