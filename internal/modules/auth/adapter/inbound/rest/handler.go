package rest

import (
	"tradex/internal/modules/auth/adapter/inbound/rest/dto"
	"tradex/internal/modules/auth/application/inbound"
	"tradex/internal/shared/response"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	register inbound.Register
	login inbound.Login
}

func NewHandler(
	register inbound.Register,
	login inbound.Login,
) *Handler {
	return &Handler{
		register: register,
		login: login,
	}
}

func (h *Handler) Register(ctx *gin.Context) {
	var request dto.RegisterRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	result, err := h.register.Register(ctx.Request.Context(), request)
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	response.Created(ctx, "user registered successfully", result)
}

func (h *Handler) Login(ctx *gin.Context) {
	var request dto.LoginRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	result, err := h.login.Login(ctx.Request.Context(), request)
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	response.OK(ctx, "login successfully", result)
}
