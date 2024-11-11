package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"homeworktodolist/internal/config"
	"homeworktodolist/internal/err_handler"
	userHandlers "homeworktodolist/internal/http/user"
	userRepo "homeworktodolist/internal/repository/postgres/user"
	userRedisRepo "homeworktodolist/internal/repository/redis/user"
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
	redisClient := redis.Connect(&cfg.RedisConfig)

	//txmanager
	txmanager := tx_manager.NewTxManager(postgresDb)

	//Repo
	userRepo := userRepo.NewUserRepo(txmanager)
	userRedisRepo := userRedisRepo.NewUserRepo(redisClient, cfg)

	//Service
	userService := userService.NewUserService(userRepo, userRedisRepo, cfg)

	//Handlers
	userHandler := userHandlers.NewUserHandler(cfg, userService)

	//fiber
	fiberApp := fiber.New(fiber.Config{
		ErrorHandler: err_handler.ErrorHandler,
	})

	//groups
	userGroup := fiberApp.Group("/user")
	userHandlers.MapUserRoutes(userGroup, userHandler)

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
