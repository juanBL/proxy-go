package users

import (
	"net/http"
	"zenrows-proxy/internal/findByApiKey"
	"zenrows-proxy/kit/query"

	"github.com/gin-gonic/gin"
	proxy "zenrows-proxy/internal"
)

type findUserRequest struct {
	ApiKey string `json:"api_key"`
}

type getResponse struct {
	ApiKey         string `json:"api_key"`
	ExpirationDate string `json:"expiration_date"`
}

func FindUserByApiKeyHandler(queryBus query.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req findUserRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		var queryResponse, err = queryBus.Ask(ctx, findByApiKey.NewUserFindByApiKeyQuery(
			req.ApiKey,
		))

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, []getResponse{})
			return
		}
		users, ok := queryResponse.([]proxy.User)
		if ok {
			var response = make([]getResponse, 0, len(users))
			for _, user := range users {
				response = append(response, getResponse{
					ApiKey:         user.ApiKey().String(),
					ExpirationDate: user.ExpirationDate().String(),
				})
			}
			ctx.JSON(http.StatusOK, response)
			return
		}
		ctx.JSON(http.StatusInternalServerError, []getResponse{})
	}
}
