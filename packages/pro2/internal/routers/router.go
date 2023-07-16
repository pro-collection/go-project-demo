package routers

import (
	"github.com/gin-gonic/gin"
	"go-project-demo/packages/pro2/internal/routers/api"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	article := api.NewArticle()
	tag := api.NewTag()

	api_v1 := r.Group("/api/v1")

	{
		api_v1.POST("/tags", tag.Create)
		api_v1.DELETE("/tags/:id", tag.Delete)
		api_v1.PUT("/tags/id", tag.Update)
		api_v1.PATCH("/tags/:id/state", tag.Update)
		api_v1.GET("/tags", tag.List)

		api_v1.POST("/articles", article.Create)
		api_v1.DELETE("/articles/:id", article.Delete)
		api_v1.PUT("/articles", article.Update)
		api_v1.PATCH("/articles/:id/state", article.Update)
		api_v1.GET("/articles/:id", article.Get)
		api_v1.GET("/articles", article.List)
	}

	return r
}
