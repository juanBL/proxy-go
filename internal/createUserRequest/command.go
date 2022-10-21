package createUserRequest

import (
	"context"
	"errors"
	"zenrows-proxy/internal/findByApiKey"
	"zenrows-proxy/kit/command"
)

const UserRequestCommandType command.Type = "command.create.user_request"

type UserRequestCommand struct {
	apiKey  string
	url     string
	headers []string
}

func NewUserRequestCommand(apiKey string, url string, headers []string) UserRequestCommand {
	return UserRequestCommand{
		apiKey:  apiKey,
		url:     url,
		headers: headers,
	}
}

func (q UserRequestCommand) Type() command.Type {
	return UserRequestCommandType
}

type UserRequestCreatorCommandHandler struct {
	service            findByApiKey.UserService
	userRequestService UserRequestService
}

func NewUserRequestCreatorCommandHandler(service findByApiKey.UserService, userRequestService UserRequestService) UserRequestCreatorCommandHandler {
	return UserRequestCreatorCommandHandler{
		service:            service,
		userRequestService: userRequestService,
	}
}

func (h UserRequestCreatorCommandHandler) Handle(ctx context.Context, q command.Command) error {
	requestCommand, ok := q.(UserRequestCommand)
	if !ok {
		return errors.New("unexpected query")
	}

	user, err := h.service.FindByApiKey(ctx, requestCommand.apiKey)

	if err == nil {
		err := h.userRequestService.CreateUserRequest(ctx, user, requestCommand.url, requestCommand.headers)
		if err != nil {
			return err
		}
	}

	return err
}
