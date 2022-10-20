package findByApiKey

import (
	"context"
	proxy "zenrows-proxy/internal"
)

type UserService struct {
	userRepository proxy.UserRepository
}

func NewUserService(userRepository proxy.UserRepository) UserService {
	return UserService{
		userRepository: userRepository,
	}
}

func (s UserService) FindByApiKey(ctx context.Context, apiKey string) ([]proxy.User, error) {
	ak, _ := proxy.NewApiKey(apiKey)
	return s.userRepository.FindByApiKey(ctx, ak)
}
