package creating

import (
	"context"
	"errors"
	"zenrows-proxy/kit/command"
)

const UserCommandType command.Type = "command.creating.user"

type UserCommand struct {
	apiKey         string
	expirationDate string
}

func NewUserCommand(apiKey, expirationDate string) UserCommand {
	return UserCommand{
		apiKey:         apiKey,
		expirationDate: expirationDate,
	}
}

func (c UserCommand) Type() command.Type {
	return UserCommandType
}

type UserCommandHandler struct {
	service UserService
}

func NewUserCommandHandler(service UserService) UserCommandHandler {
	return UserCommandHandler{
		service: service,
	}
}

func (h UserCommandHandler) Handle(ctx context.Context, cmd command.Command) error {
	createUserCmd, ok := cmd.(UserCommand)
	if !ok {
		return errors.New("unexpected command")
	}

	return h.service.CreateUser(
		ctx,
		createUserCmd.apiKey,
		createUserCmd.expirationDate,
	)
}
