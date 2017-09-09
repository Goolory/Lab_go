package main

import (
	//"flag"
	//"config"
	"model"
	"controller"
	"github.com/jinzhu/gorm"
	 _ "github.com/go-sql-driver/mysql"
	"github.com/gin-gonic/gin"
	"middleware"
)

func main() {
	//configPath := flag.String("conf", "../config/config.json", "Config file path")
	//flag.Parse()
	//
	//err := config.LoadConfig(*configPath)
	//if err != nil {
	//	print("Config Failed!!!!", err)
	//	return
	//}

	db, err := gorm.Open("mysql", "go:go123456@/labgo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		print("Open db Faild !!", err)
		return
	}

	db.LogMode(true)
	model.InitDBModel(db)

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	router := r.Group("/")
	router.Use(middleware.CorAllowHandler)
	router.GET("/",controller.IndexHandler)

	r.Run()

}
