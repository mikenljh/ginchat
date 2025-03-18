package service

import (
	"fmt"
	"ginchat/models"
	"ginchat/utils"
	"math/rand"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// GetUserList
// @Tags        用户模块
// @Summary     查看所有用户
// @Description 返回所有用户
// @Success      200 {string} json{"code","message"}
// @Router      /user/getUserList [get]
func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()

	c.JSON(200, gin.H{
		"code":    0,
		"message": "success",
		"data":    data,
	})

}

// CreateUser
// @Tags        用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @param repassword query string false "确认密码"
// @Summary     添加用户
// @Description 新增用户
// @Success      200 {string} json{"code","message"}
// @Router      /user/createUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	password := c.Query("password")
	repassword := c.Query("repassword")
	salt := fmt.Sprintf("%06d", rand.Int31())
	data := models.FindUserByName(user.Name)
	if data.Name != "" {
		c.JSON(-1, gin.H{
			"message": "该用户名已注册",
		})
		return
	}
	if password != repassword {
		c.JSON(200, gin.H{
			"message": "两次输入的密码不一致",
		})
		return
	}
	user.PassWord = utils.MakePassword(password, salt)
	user.Salt = salt
	models.CreateUser(user)

	c.JSON(200, gin.H{
		"message": "新增用户成功",
	})

}

// DeleteUser
// @Tags        用户模块
// @param id query string false "id"
// @Summary     删除用户
// @Description 删除用户
// @Success      200 {string} json{"code","message"}
// @Router      /user/deleteUser [get]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(user)
	c.JSON(200, gin.H{
		"message": "删除用户成功",
	})

}

// UpdateUser
// @Tags        用户模块
// @param id formData string false "id"
// @param name formData string false "name"
// @param password formData string false "password"
// @param phone formData string false "phone"
// @param email formData string false "email"
// @Summary     修改用户
// @Description 修改用户
// @Success      200 {string} json{"code","message"}
// @Router      /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.PassWord = c.PostForm("password")
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")
	fmt.Println("update:", user)
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"message": "修改参数不匹配",
		})
	}
	models.UpdateUser(user)
	c.JSON(200, gin.H{
		"message": "修改用户成功",
	})

}

// FindUserByNameAndPwd
// @Tags        用户模块
// @Summary     登录
// @Description 通过用户名和密码登录
// @param name query string false "用户名"
// @param password query string false "密码"
// @Success      200 {string} json{"code","message"}
// @Router      /user/findUserByNameAndPwd [post]
func FindUserByNameAndPwd(c *gin.Context) {
	data := models.UserBasic{}
	name := c.Query("name")
	password := c.Query("password")
	user := models.FindUserByName(name)
	if user.Name == "" {
		c.JSON(200, gin.H{
			"message": "该用户不存在",
		})
		return
	}

	flag := utils.ValidPassword(password, user.Salt, user.PassWord)

	if !flag {
		c.JSON(200, gin.H{
			"message": "密码不正确",
		})
		return
	}
	pwd := utils.MakePassword(password, user.Salt)
	data = models.FindUserByNameAndPwd(name, pwd)

	c.JSON(200, gin.H{
		"code":    0,
		"message": "success",
		"data":    data,
	})

}
