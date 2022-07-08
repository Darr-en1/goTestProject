package web

import "github.com/gin-gonic/gin"

type RequestParams struct {
	Id int `form:"id"`
}

func GetBlog(c *gin.Context) (interface{}, error) {
	req := new(RequestParams)
	_ = c.ShouldBindQuery(req)

	return getBlogService(req.Id)

}
