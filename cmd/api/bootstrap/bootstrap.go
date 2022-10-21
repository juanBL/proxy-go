package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"zenrows-proxy/internal/createUserRequest"
	"zenrows-proxy/internal/creating"
	"zenrows-proxy/internal/findByApiKey"
	"zenrows-proxy/internal/platform/bus/inmemory"
	"zenrows-proxy/internal/platform/server"
	"zenrows-proxy/internal/platform/storage/mysql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
)

func Run() error {
	var cfg config
	err := envconfig.Process("PROXY", &cfg)
	if err != nil {
		return err
	}

	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbPort, cfg.DbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	var (
		commandBus = inmemory.NewCommandBus()
		queryBus   = inmemory.NewQueryBus()
		eventBus   = inmemory.NewEventBus()
	)

	userRepository := mysql.NewUserRepository(db, cfg.DbTimeout)
	userRequestRepository := mysql.NewUserRequestRepository(db, cfg.DbTimeout)

	creatingUserService := creating.NewUserService(userRepository, eventBus)
	getUserService := findByApiKey.NewUserService(userRepository)
	createUserRequestService := createUserRequest.NewUserRequestService(userRequestRepository)

	createUserCommandHandler := creating.NewUserCommandHandler(creatingUserService)
	createUserRequestCommandHandler := createUserRequest.NewUserRequestCreatorCommandHandler(getUserService, createUserRequestService)

	commandBus.Register(creating.UserCommandType, createUserCommandHandler)
	commandBus.Register(createUserRequest.UserRequestCommandType, createUserRequestCommandHandler)

	ctx, srv := server.New(context.Background(), cfg.Host, cfg.Port, cfg.ShutdownTimeout, commandBus, queryBus)
	return srv.Run(ctx)
}

type config struct {
	// Server configuration
	Host            string        `default:"localhost"`
	Port            uint          `default:"8080"`
	ShutdownTimeout time.Duration `default:"10s"`
	// Database configuration
	DbUser    string        `default:"zenrows"`
	DbPass    string        `default:"zenrows"`
	DbHost    string        `default:"localhost"`
	DbPort    uint          `default:"3306"`
	DbName    string        `default:"zenrows"`
	DbTimeout time.Duration `default:"5s"`
}
