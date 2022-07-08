package main

import (
	"github.com/gin-gonic/gin"
	"goTestProject/error/customError/v2/errors"
	"goTestProject/error/customError/v2/web"
	"log"
)

func initRouter() *gin.Engine {
	var router = gin.Default()
	router.GET("/blog", errors.ErrWrapper(web.GetBlog))
	return router
}

func main() {

	router := initRouter()

	err := router.Run()
	if err != nil {
		log.Panic("service startup fails")
	}
}
