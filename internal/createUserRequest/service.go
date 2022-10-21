package createUserRequest

import (
	"context"
	"time"
	proxy "zenrows-proxy/internal"
)

type UserRequestService struct {
	userRequestRepository proxy.UserRequestRepository
}

func NewUserRequestService(userRequestRepository proxy.UserRequestRepository) UserRequestService {
	return UserRequestService{
		userRequestRepository: userRequestRepository,
	}
}

func (service UserRequestService) CreateUserRequest(ctx context.Context, user proxy.User, url string, headers []string) error {
	ur, err := proxy.NewUserRequest(user.ApiKey().String(), url, headers)

	if err == nil {
		expirationDate, _ := time.Parse("2006-01-02 15:04:05", user.ExpirationDate().String())
		now := time.Now()

		if now.After(expirationDate) {
			return proxy.ErrExpirationDate
		}

		err := service.userRequestRepository.Save(ctx, ur)
		if err != nil {
			return err
		}
	}

	return err
}
