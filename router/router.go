package router

import (
	"github.com/gin-gonic/gin"
	"go-img/controller"
	"go-img/db"
)

func Router() {
	router := gin.Default()

	router.LoadHTMLGlob("public/templates/*.html")

	handler := controller.ImgHandler{
		Db: db.Get(),
	}

	router.GET("/", handler.GetAll)

	router.POST("/upload", handler.Upload)

	router.Run()

}
