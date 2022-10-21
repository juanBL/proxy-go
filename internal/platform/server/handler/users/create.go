package users

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	proxy "zenrows-proxy/internal"
	"zenrows-proxy/internal/creating"
	"zenrows-proxy/kit/command"
)

type createRequest struct {
	ApiKey         string `json:"api_key" binding:"required"`
	ExpirationDate string `json:"expiration_date" binding:"required"`
}

func CreateUserHandler(commandBus command.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err := commandBus.Dispatch(ctx, creating.NewUserCommand(
			req.ApiKey,
			req.ExpirationDate,
		))

		if err != nil {
			switch {
			case errors.Is(err, proxy.ErrInvalidApiKey),
				errors.Is(err, proxy.ErrEmptyExpirationDate):
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			default:
				ctx.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}

		ctx.Status(http.StatusCreated)
	}
}
