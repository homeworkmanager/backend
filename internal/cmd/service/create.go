package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"homeworktodolist/internal/config"
	"homeworktodolist/internal/err_handler"
	userHandler "homeworktodolist/internal/http/user"
	userRepo "homeworktodolist/internal/repository/postgres/user"
	userService "homeworktodolist/internal/service/user"
	"homeworktodolist/internal/tx_manager"
	postgres "homeworktodolist/pkg/db/postgres"
	"homeworktodolist/pkg/db/redis"
	"os"
	"os/signal"
)

func createApp() {

	//config
	cfg := config.NewCfg()

	//database
	postgresDb := postgres.Connect(&cfg.PGConfig)
	redisDb := redis.Connect(&cfg.RedisConfig)

	//txmanager
	txmanager := tx_manager.NewTxManager(postgresDb)

	//Repo
	userRepo := userRepo.NewUserRepo(txmanager)

	//Service
	userService := userService.NewUserService(userRepo, cfg)

	//Handlers
	userHandler := userHandler.NewUserHandler(cfg, userService)

	//fiber
	fiberApp := fiber.New(fiber.Config{
		ErrorHandler: err_handler.ErrorHandler,
	})

	//fiber listen

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt)

	go func() {
		if err := fiberApp.Listen(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)); err != nil {
			panic(err)
		}
	}()

	<-exit

	if err := fiberApp.Shutdown(); err != nil {
		panic(err)
	}

}
