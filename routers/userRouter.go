package routers


import (
	"github.com/gbubemi/mysql2/controllers"
	"github.com/gin-gonic/gin"
)


func UserRoutes(router *gin.Engine, userRepo *controllers.UserRepo) {
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/", userRepo.CreateUser)
		userRoutes.GET("/", userRepo.GetUsers)
		userRoutes.GET("/:id", userRepo.GetUser)
		userRoutes.PUT("/:id", userRepo.UpdateUser)
		userRoutes.DELETE("/:id", userRepo.DeleteUser)
	}
}