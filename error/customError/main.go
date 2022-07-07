package main

import (
	"github.com/gin-gonic/gin"
	"goTestProject/error/customError/errors"
	"goTestProject/error/customError/web"
	"log"
	"net/http"
)

func ErrWrapper(handler func(g *gin.Context) (interface{}, error)) func(*gin.Context) {
	return func(context *gin.Context) {
		data, err := handler(context)
		if err != nil {
			var status int
			errorType := errors.GetType(err)
			switch errorType {
			case errors.BadRequest:
				status = http.StatusBadRequest
			case errors.NotFound:
				status = http.StatusNotFound
			default:
				log.Printf("%+v", err)
				status = http.StatusInternalServerError
			}
			web.Result(status, struct{}{}, errorType.String(), web.ERROR, context)
		} else {
			web.OkWithData(data, context)
		}

	}
}

func initRouter() *gin.Engine {
	var router = gin.Default()
	router.GET("/blog", ErrWrapper(web.GetBlog))
	return router
}

func main() {

	router := initRouter()

	err := router.Run()
	if err != nil {
		log.Print("service startup fails")
	}
}
