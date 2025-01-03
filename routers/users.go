package routers

import (
	"github.com/fulsep/docker-crud-backend/tree/main/controllers"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {
	r.GET("", controllers.ListAllUsers)
	r.GET("/:id", controllers.DetailUser)
	r.PATCH("/:id", controllers.UpdateUser)
	r.POST("", controllers.CreateUser)
	r.DELETE("/:id", controllers.DeleteUser)
}
