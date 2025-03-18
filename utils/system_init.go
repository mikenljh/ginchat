package utils

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Config appL", viper.Get("app"))
}

func InitMySQL() {
	fmt.Println("Config mysql", viper.Get("mysql"))
	DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{})
	// user := models.UserBasic{}
	// DB.Find(&user)
	// fmt.Println(user)
}
