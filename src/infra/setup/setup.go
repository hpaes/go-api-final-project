package setup

import (
	"context"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hpaes/go-api-final-project/src/api/controller"
	"github.com/hpaes/go-api-final-project/src/core/application/usecase"
	"github.com/hpaes/go-api-final-project/src/infra/config"
	infraDb "github.com/hpaes/go-api-final-project/src/infra/database"
	"github.com/hpaes/go-api-final-project/src/infra/logger"
	"github.com/hpaes/go-api-final-project/src/infra/repository"
	"github.com/hpaes/go-api-final-project/src/infra/router"
	"github.com/hpaes/go-api-final-project/src/infra/server"
)

type setup struct {
	config    *config.AppConfig
	webServer server.ServerConfig
	db        *repository.UserRepository
	router    router.GinRouter
	logger    logger.LoggerService
}

func NewSetup() *setup {
	return &setup{}
}

func (s *setup) InitLogger() *setup {
	s.logger = logger.NewLoggerService()

	s.logger.Info("Logger initialized")
	return s
}

func (s *setup) WithAppConfig() *setup {
	appConfig, err := config.LoadConfig()
	if err != nil {
		s.logger.Fatal("Error loading config", err)
	}
	s.config = appConfig
	return s
}

func (s *setup) WithDatabase() *setup {
	db, err := infraDb.NewSqlConnection(s.config.MySQL)
	if err != nil {
		s.logger.Fatal("Error connecting to database", err)
	}
	s.db = repository.NewUserRepository(db)
	s.logger.Info("Database initialized")
	return s
}

func (s *setup) WithRouter() *setup {
	rc := controller.NewRegisterUserController(usecase.NewRegisterUser(s.db, s.logger))
	gc := controller.NewGetUserDetailController(usecase.NewGetUserDetail(s.db, s.logger))
	uc := controller.NewUpdateUserController(usecase.NewUpdateUser(s.db, s.logger))
	dc := controller.NewRemoveUserController(usecase.NewRemoveUser(s.db, s.logger))
	lc := controller.NewListUsersController(usecase.NewListUsers(s.db, s.logger))

	s.router = router.NewGinEngine(gin.Default(), rc, gc, dc, lc, uc)
	return s
}

func (s *setup) WithServer() *setup {
	port, err := strconv.ParseInt(s.config.Application.Server.Port, 10, 64)
	if err != nil {
		s.logger.Fatal("Error parsing port", err)
	}

	duration, err := time.ParseDuration(s.config.Application.Server.Timeout + "s")
	if err != nil {
		s.logger.Fatal("Error parsing duration", err)
	}

	s.webServer = server.NewWebServer(s.router, port, duration*time.Second)
	s.logger.Info("Server initialized")
	return s
}

func (s *setup) Run(ctx context.Context, wg *sync.WaitGroup) {
	s.logger.Info("Server running on port %s", s.config.Application.Server.Port)
	s.webServer.Listen(ctx, wg)
}
