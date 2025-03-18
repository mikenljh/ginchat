package service

import "github.com/gin-gonic/gin"

// GetIndex
// @Tags        首页
// @Summary     欢迎页面
// @Description 返回欢迎信息
// @Success      200 {string} string "Welcome"
// @Router      /index [get]
func GetIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "welcome",
	})
}
