package test

import (
	"log"

	"fmt"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

type Reception struct {
	UID                  string  `gorm:"primaryKey" json:"uid"`
	Mail                 string  `gorm:"not null" json:"mail"`
	Name                 string  `gorm:"not null" json:"name"`
	AttendsFirstDay      bool    `gorm:"not null" json:"attends_first_day"`
	AttendsSecondDay     bool    `gorm:"not null" json:"attends_second_day"`
	TemperatureFirstDay  float32 `json:"temperature_first_day"`
	TemperatureSecondDay float32 `json:"temperature_second_day"`
}

var DB *gorm.DB

func NewSqlHandler() {
	dsn := "root:gorupass@tcp(mysql:3306)/reception"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf(err.Error())
	}
	DB = db
}

func Init() {
	DB.AutoMigrate(&Reception{})
}
func InitReception() {
	var users = []Reception{
		{UID: "33u@2", Mail: "goru.technology@gmail.com", Name: "伊藤優汰", AttendsFirstDay: true, AttendsSecondDay: false},
	}
	DB.Save(&users)
}
func GetReceptionUserInfo(uid string) (Reception, error) {
	var receptionUserInfo Reception
	err := DB.First(&receptionUserInfo, "uid = ?", uid).Error
	if err != nil {
		return receptionUserInfo, err
	}
	return receptionUserInfo, nil
}

func main() {
	NewSqlHandler()
	Init()
	InitReception()
	test, err := GetReceptionUserInfo("33u@2")

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(test)
}
