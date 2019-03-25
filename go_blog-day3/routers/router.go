package routers

import (
	"github.com/gin-gonic/gin"

	"../pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiv1 := r.Gruop("/api/v1")
	{

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticle)

		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticles)

		//新建文章
		apiv1.POST("/articles", v1.AddArticle)

		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)

		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}
