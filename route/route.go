package route

import (
	"MovieApi/controller"
	"MovieApi/service"

	"github.com/gin-gonic/gin"
)

func RegisterApi(r *gin.Engine, app service.ServiceInterface) {
	controllers := controller.NewHttpServer(app)

	movieRouter := r.Group("/Movies")
	{
		movieRouter.POST("/", controllers.CreateMovie)
		movieRouter.GET("/", controllers.GetAllMovie)
		movieRouter.GET("/:movieId", controllers.GetMovieByID)
		movieRouter.PUT("/:movieId", controllers.UpdateMovie)
		movieRouter.DELETE("/:movieId", controllers.DeleteMovie)
	}
}
