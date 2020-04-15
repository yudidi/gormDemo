package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	local_db = "root:123456@tcp(localhost:3306)/test?charset=utf8"
)

type MyUser struct {
	ID   int `gorm:"TYPE:int(11);NOT NULL;PRIMARY_KEY;INDEX"`
	Name string
}

type MyProfile struct {
	ID        int `gorm:"TYPE:int(11);NOT NULL;PRIMARY_KEY;INDEX"`
	Name      string
	User      MyUser `gorm:"foreignkey:UserRefer"`
	UserRefer uint
}

func main() {
	db, err := gorm.Open("mysql", local_db)
	db.LogMode(true)

	if err != nil {
		panic(err)
	}
	defer db.Close()
	related(db)
	fmt.Println("=====")
	association(db)
}

func setData(db *gorm.DB) {
	db.Debug().AutoMigrate(&MyUser{})
	db.Debug().AutoMigrate(&MyProfile{})
	db.Debug().Create(&MyUser{ID: 1, Name: "uname1"})
	db.Debug().Create(&MyUser{ID: 2, Name: "uname2"})
	db.Debug().Create(&MyProfile{ID: 11, Name: "pname2", UserRefer: 1})
	db.Debug().Create(&MyProfile{ID: 22, Name: "pname2", UserRefer: 2})
}

// 查询1个结构体，同时查询关联的子结构体。
func related(db *gorm.DB) {
	var profile MyProfile
	db.Debug().First(&profile)
	fmt.Println(fmt.Sprintf("%+v", profile))
	db.Model(&profile).Related(&profile.User, "UserRefer")
	fmt.Println(fmt.Sprintf("%+v", profile))
}

// 查询1个结构体，同时查询关联的子结构体。
func association(db *gorm.DB) {
	var profile MyProfile
	db.Debug().First(&profile)
	fmt.Println(fmt.Sprintf("%+v", profile))
	// 查询profile相关的User, 并且把值赋给&profile.User
	db.Model(&profile).Association("User").Find(&profile.User)
	fmt.Println(fmt.Sprintf("%+v", profile))
}
