package findByApiKey

import (
	"context"
	"errors"
	proxy "zenrows-proxy/internal"
	"zenrows-proxy/kit/query"
)

const UserQueryType query.Type = "query.getting.user"

type UserFindByApiKeyQuery struct {
	apiKey string
}

func NewUserFindByApiKeyQuery(apiKey string) UserFindByApiKeyQuery {
	return UserFindByApiKeyQuery{
		apiKey: apiKey,
	}
}

func (q UserFindByApiKeyQuery) Type() query.Type {
	return UserQueryType
}

type UserFindByApiKeyQueryHandler struct {
	service UserService
}

func NewUserFindByApiKeyQueryHandler(service UserService) UserFindByApiKeyQueryHandler {
	return UserFindByApiKeyQueryHandler{
		service: service,
	}
}

func (h UserFindByApiKeyQueryHandler) Handle(ctx context.Context, q query.Query) (query.Response, error) {
	findUserByApiKeyQuery, ok := q.(UserFindByApiKeyQuery)
	if !ok {
		return proxy.User{}, errors.New("unexpected query")
	}

	return h.service.FindByApiKey(
		ctx,
		findUserByApiKeyQuery.apiKey,
	)
}
