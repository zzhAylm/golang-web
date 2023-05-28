package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserInfo struct {
	gorm.Model
	Name  string
	Age   int
	Hobby string
}

func main() {

	db, err := gorm.Open(mysql.Open("root:123456@tcp(123.57.173.24:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		return
	}
	name := db.Name()
	fmt.Println(name)
	// 自动创建表
	err = db.AutoMigrate(&UserInfo{})

	userInfo := UserInfo{Name: "zzh1", Age: 18, Hobby: "游戏"}

	db.Create(&userInfo)

	u := UserInfo{}
	db.First(&u)
	fmt.Println(u)
	sqlDB, err := db.DB()

	err = sqlDB.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

}
