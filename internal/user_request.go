package proxy

import (
	"context"
	"encoding/json"
	"errors"
)

var ErrEmptyUrl = errors.New("the field url can not be empty")

type Url struct {
	value string
}

func NewUrl(value string) (Url, error) {
	if value == "" {
		return Url{}, ErrEmptyUrl
	}

	return Url{
		value: value,
	}, nil
}

func (url Url) String() string {
	return url.value
}

func (headers Headers) String() string {
	value, _ := json.Marshal(headers)
	return string(value)
}

type Header struct {
	HeaderName  int
	HeaderValue string
}

type Headers []Header

func NewHeaders(requestHeaders []string) (Headers, error) {
	headers := Headers{}
	for i, value := range requestHeaders {
		header := Header{i, value}
		headers = append(headers, header)
	}

	return headers, nil
}

type UserRequest struct {
	apiKey  ApiKey
	url     Url
	headers []Header
}

type UserRequestRepository interface {
	Save(ctx context.Context, userRequest UserRequest) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=UserRepository

func NewUserRequest(apiKey, url string, headers []string) (UserRequest, error) {
	apiKeyVO, err := NewApiKey(apiKey)
	if err != nil {
		return UserRequest{}, err
	}

	urlVO, err := NewUrl(url)
	if err != nil {
		return UserRequest{}, err
	}

	headersVO, err := NewHeaders(headers)
	if err != nil {
		return UserRequest{}, err
	}

	u := UserRequest{
		apiKey:  apiKeyVO,
		url:     urlVO,
		headers: headersVO,
	}
	return u, nil
}

func (userRequest UserRequest) ApiKey() ApiKey {
	return userRequest.apiKey
}

func (userRequest UserRequest) Url() Url {
	return userRequest.url
}

func (userRequest UserRequest) Headers() Headers {
	return userRequest.headers
}
