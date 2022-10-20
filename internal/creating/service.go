package creating

import (
	"context"
	proxy "zenrows-proxy/internal"

	"zenrows-proxy/kit/event"
)

type UserService struct {
	userRepository proxy.UserRepository
	eventBus       event.Bus
}

func NewUserService(userRepository proxy.UserRepository, eventBus event.Bus) UserService {
	return UserService{
		userRepository: userRepository,
		eventBus:       eventBus,
	}
}

func (userService UserService) CreateUser(ctx context.Context, apiKey, expirationDate string) error {
	u, err := proxy.NewUser(apiKey, expirationDate)
	if err != nil {
		return err
	}

	if err := userService.userRepository.Save(ctx, u); err != nil {
		return err
	}

	return userService.eventBus.Publish(ctx, u.PullEvents())
}
