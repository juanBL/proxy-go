package users

import (
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	proxy "zenrows-proxy/internal"
	"zenrows-proxy/internal/createUserRequest"
	"zenrows-proxy/kit/command"

	"github.com/gin-gonic/gin"
)

type proxyRequest struct {
	ApiKey string `json:"api_key"`
	Url    string `json:"url"`
}

func CreateUserRequestHandler(commandBus command.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestHeaders := strings.Split(ctx.Request.Header["Headers"][0], ",")
		var req proxyRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err := commandBus.Dispatch(ctx, createUserRequest.NewUserRequestCommand(
			req.ApiKey,
			req.Url,
			requestHeaders,
		))
		if err != nil {
			switch {
			case errors.Is(err, proxy.ErrInvalidApiKey),
				errors.Is(err, proxy.ErrEmptyExpirationDate):
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			case errors.Is(err, proxy.ErrUserNotFound):
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			case errors.Is(err, proxy.ErrExpirationDate):
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			default:
				ctx.JSON(http.StatusInternalServerError, "not user find with this apikey")
				return
			}
		}

		if err == nil {
			req, err := http.NewRequest("GET", req.Url, nil)
			if err != nil {
				log.Fatalln(err)
			}

			for i, value := range requestHeaders {
				key := strconv.Itoa(i + 1)
				req.Header.Add("Header"+key, value)
			}

			client := &http.Client{}
			resp, err := client.Do(req)

			if err != nil {
				log.Fatalln(err)
			}

			defer resp.Body.Close()

			b, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatalln(err)
			}

			ctx.Status(http.StatusCreated)
			ctx.String(http.StatusOK, string(b))
			return
		}
	}
}
