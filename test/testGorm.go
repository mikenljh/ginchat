package main

import (
	"ginchat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// // type Product struct {
// // 	gorm.Model
// // 	Code  string
// // 	Price uint
// // }

func main() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic("failed to open")
	}
	// 迁移 schema
	db.AutoMigrate(&models.Message{})
	db.AutoMigrate(&models.Contact{})
	db.AutoMigrate(&models.GroupBasic{})
}

// 	// Create
// 	// user := &models.UserBasic{
// 	// 	Name:          "John Doe",
// 	// 	PassWord:      "1234",
// 	// 	Phone:         "1234567890",
// 	// 	Email:         "johndoe@example.com",
// 	// 	Identity:      "1234567890",
// 	// 	ClientIp:      "127.0.0.1",
// 	// 	ClientPort:    "9998",
// 	// 	LoginTime:     time.Now(),
// 	// 	HeartbeatTime: time.Now(),
// 	// 	LogoutTime:    time.Now(),
// 	// 	IsLogout:      false,
// 	// }

// 	// db.Create(user)

// 	// // Read
// 	// fmt.Println(db.First(user, 1))

// 	// // Update - 将 product 的 price 更新为 200
// 	// db.Model(user).Update("PassWord", "1234")

// }
