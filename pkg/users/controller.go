package users

import (
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/users")
	routes.POST("/", h.AddUser)
	routes.GET("/", h.GetUsers)
	routes.GET("/:id", h.GetUser)
	routes.PUT("/:id", h.UpdateUser)
	routes.DELETE("/:id", h.DeleteUser)
}
