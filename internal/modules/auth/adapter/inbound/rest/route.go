package rest

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.RouterGroup, h *Handler) {
	auth := r.Group("/auth")

	auth.POST("/register", h.Register)
	auth.POST("/login", h.Login)
}
