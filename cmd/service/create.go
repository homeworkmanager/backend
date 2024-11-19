package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"homeworktodolist/internal/config"
	"homeworktodolist/internal/err_handler"
	adminHandlers "homeworktodolist/internal/http/admin"
	userHandlers "homeworktodolist/internal/http/user"
	middleware "homeworktodolist/internal/middleware"
	classRepo "homeworktodolist/internal/repository/postgres/class"
	groupRepo "homeworktodolist/internal/repository/postgres/group"
	subjectRepo "homeworktodolist/internal/repository/postgres/subject"
	userRepo "homeworktodolist/internal/repository/postgres/user"
	userRedisRepo "homeworktodolist/internal/repository/redis/user"
	adminService "homeworktodolist/internal/service/admin"
	classService "homeworktodolist/internal/service/class"
	groupService "homeworktodolist/internal/service/group"
	subjectService "homeworktodolist/internal/service/subject"
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
	groupRepo := groupRepo.NewGroupRepo(txmanager)
	classRepo := classRepo.NewClassRepo(txmanager)
	subjectRepo := subjectRepo.NewSubjectRepo(txmanager)

	//Service
	userService := userService.NewUserService(userRepo, userRedisRepo, cfg)

	groupService := groupService.NewGroupService(groupRepo)

	subjectService := subjectService.NewSubjectService(subjectRepo)

	classService := classService.NewClassService(groupService, classRepo, subjectService, txmanager)

	adminService := adminService.NewAdminService(groupService, classService, subjectService)

	//Handlers
	userHandler := userHandlers.NewUserHandler(cfg, userService)

	adminHandler := adminHandlers.NewAdminHandler(adminService)

	//fiber
	fiberApp := fiber.New(fiber.Config{
		ErrorHandler: err_handler.ErrorHandler,
	})

	//middleware
	mw := middleware.NewMwManager(userRedisRepo)
	//groups
	userGroup := fiberApp.Group("/user")
	userHandlers.MapUserRoutes(userGroup, userHandler, mw)

	adminGroup := fiberApp.Group("/admin")
	adminHandlers.MapAdminRoutes(adminGroup, adminHandler, mw)

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
