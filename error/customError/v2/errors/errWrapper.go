package errors

import (
	errors3 "errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ErrWrapper(handler func(g *gin.Context) (interface{}, error)) func(*gin.Context) {
	return func(context *gin.Context) {
		data, err := handler(context)
		if err != nil {
			var customErr = new(CustomError)
			// 根据业务扩展错误类型，比如validate 错误(validator.ValidationErrors)， 可以单独处理，但我还是推荐用 CustomError 对这些 error 封装
			if errors3.As(err, &customErr) {
				ResultFail(http.StatusBadRequest, customErr.Code, context)
			} else {
				ResultFail(http.StatusInternalServerError, ServerError, context)
			}
			// 自行更改成特定的log
			log.Printf("%+v", err)

		} else {
			ResultOk(data, context)
		}

	}
}
