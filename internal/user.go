package proxy

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"zenrows-proxy/kit/event"
)

var ErrInvalidApiKey = errors.New("invalid api key")

type ApiKey struct {
	value string
}

func NewApiKey(value string) (ApiKey, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return ApiKey{}, fmt.Errorf("%w: %s", ErrInvalidApiKey, value)
	}

	return ApiKey{
		value: v.String(),
	}, nil
}

func (apiKey ApiKey) String() string {
	return apiKey.value
}

var ErrEmptyExpirationDate = errors.New("the field expiration date can not be empty")

type ExpirationDate struct {
	value string
}

func NewExpirationDate(value string) (ExpirationDate, error) {
	if value == "" {
		return ExpirationDate{}, ErrEmptyExpirationDate
	}

	return ExpirationDate{
		value: value,
	}, nil
}

func (expirationDate ExpirationDate) String() string {
	return expirationDate.value
}

type User struct {
	apiKey         ApiKey
	expirationDate ExpirationDate

	events []event.Event
}

type UserRepository interface {
	SearchAll(ctx context.Context) ([]User, error)
	FindByApiKey(ctx context.Context, apiKey ApiKey) ([]User, error)
	Save(ctx context.Context, user User) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=UserRepository

func NewUser(apiKey, expirationDate string) (User, error) {
	apiKeyVO, err := NewApiKey(apiKey)
	if err != nil {
		return User{}, err
	}

	expirationDateVO, err := NewExpirationDate(expirationDate)
	if err != nil {
		return User{}, err
	}

	u := User{
		apiKey:         apiKeyVO,
		expirationDate: expirationDateVO,
	}
	return u, nil
}

func (user User) ApiKey() ApiKey {
	return user.apiKey
}

func (user User) ExpirationDate() ExpirationDate {
	return user.expirationDate
}

// Record records a new domain event.
func (user *User) Record(event event.Event) {
	user.events = append(user.events, event)
}

// PullEvents returns all the recorded domain events.
func (user User) PullEvents() []event.Event {
	evt := user.events
	user.events = []event.Event{}

	return evt
}
