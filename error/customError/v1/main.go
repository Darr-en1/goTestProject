package main

import (
	"github.com/gin-gonic/gin"
	errors2 "goTestProject/error/customError/v1/errors"
	web2 "goTestProject/error/customError/v1/web"
	"log"
	"net/http"
)

func ErrWrapper(handler func(g *gin.Context) (interface{}, error)) func(*gin.Context) {
	return func(context *gin.Context) {
		data, err := handler(context)
		if err != nil {
			var status int
			errorType := errors2.GetType(err)
			switch errorType {
			case errors2.BadRequest:
				status = http.StatusBadRequest
			case errors2.NotFound:
				status = http.StatusNotFound
			default:
				log.Printf("%+v", err)
				status = http.StatusInternalServerError
			}
			web2.Result(status, struct{}{}, errorType.String(), web2.ERROR, context)
		} else {
			web2.OkWithData(data, context)
		}

	}
}

func initRouter() *gin.Engine {
	var router = gin.Default()
	router.GET("/blog", ErrWrapper(web2.GetBlog))
	return router
}

func main() {

	router := initRouter()

	err := router.Run()
	if err != nil {
		log.Print("service startup fails")
	}
}
