package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/taingk/goxit/api/Config"
	"github.com/taingk/goxit/api/Models"
	"github.com/taingk/goxit/api/Routers"
)

var err error

func main() {

	Config.DB, err = gorm.Open("postgres", "host=localhost port=8083 user=postgres dbname=postgres password=root sslmode=disable")

	if err != nil {
		fmt.Println("statuse: ", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})
	Config.DB.AutoMigrate(&Models.Vote{})

	r := Routers.SetupRouter()
	// running
	r.Run()
}
