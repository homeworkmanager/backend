package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/zap"
	"homeworktodolist/internal/config"
	"homeworktodolist/internal/err_handler"
	adminHandlers "homeworktodolist/internal/http/admin"
	groupHandlers "homeworktodolist/internal/http/group"
	homeworkStatusHandlers "homeworktodolist/internal/http/homework_status"
	moderatorHandlers "homeworktodolist/internal/http/moderator"
	scheduleHandlers "homeworktodolist/internal/http/schedule"
	subjectHandlers "homeworktodolist/internal/http/subject"
	userHandlers "homeworktodolist/internal/http/user"
	middleware "homeworktodolist/internal/middleware"
	classRepo "homeworktodolist/internal/repository/postgres/class"
	groupRepo "homeworktodolist/internal/repository/postgres/group"
	homeworkRepo "homeworktodolist/internal/repository/postgres/homework"
	homeworkStatusRepo "homeworktodolist/internal/repository/postgres/homework_status"
	subjectRepo "homeworktodolist/internal/repository/postgres/subject"
	userRepo "homeworktodolist/internal/repository/postgres/user"
	userRedisRepo "homeworktodolist/internal/repository/redis/user"
	adminService "homeworktodolist/internal/service/admin"
	classService "homeworktodolist/internal/service/class"
	groupService "homeworktodolist/internal/service/group"
	homeworkService "homeworktodolist/internal/service/homework"
	homeworkStatusService "homeworktodolist/internal/service/homework_status"
	moderatorService "homeworktodolist/internal/service/moderator"
	scheduleService "homeworktodolist/internal/service/schedule"
	subjectService "homeworktodolist/internal/service/subject"
	userService "homeworktodolist/internal/service/user"
	"homeworktodolist/internal/tx_manager"
	postgres "homeworktodolist/pkg/db/postgres"
	"homeworktodolist/pkg/db/redis"
	"homeworktodolist/pkg/logger"
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
	homeworkRepo := homeworkRepo.NewHomeworkRepo(txmanager)
	homeworkStatusRepo := homeworkStatusRepo.NewHomeworkStatusRepo(txmanager)

	//Service
	userService := userService.NewUserService(userRepo, userRedisRepo, cfg)

	groupService := groupService.NewGroupService(groupRepo)

	subjectService := subjectService.NewSubjectService(subjectRepo)

	classService := classService.NewClassService(groupService, classRepo, subjectService, txmanager)

	homeworkService := homeworkService.NewHomeworkService(homeworkRepo)

	adminService := adminService.NewAdminService(groupService, classService, subjectService, homeworkService, txmanager)

	moderatorService := moderatorService.NewModeratorService(homeworkService)

	scheduleService := scheduleService.NewScheduleService(classService, homeworkService)

	homeworkStatusService := homeworkStatusService.NewHomeworkStatusService(homeworkStatusRepo)

	//Handlers
	userHandler := userHandlers.NewUserHandler(cfg, userService)

	adminHandler := adminHandlers.NewAdminHandler(adminService)

	groupHandler := groupHandlers.NewGroupHandler(groupService)

	moderatorHandler := moderatorHandlers.NewModeratorHandler(moderatorService)

	scheduleHandler := scheduleHandlers.NewScheduleHandler(scheduleService)

	homeworkStatusHandler := homeworkStatusHandlers.NewHomeworkStatusHandler(homeworkStatusService)

	subjectHandler := subjectHandlers.NewSubjectHandler(subjectService)

	//fiber
	fiberApp := fiber.New(fiber.Config{
		ErrorHandler: err_handler.ErrorHandler,
	})

	fiberApp.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5000",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Content-Type, Authorization",
		AllowCredentials: true,
	}))

	//middleware
	mw := middleware.NewMwManager(userRedisRepo)

	//logger
	logger := logger.InitLogger(cfg)

	fiberApp.Use(func(c *fiber.Ctx) error {
		err := c.Next()
		if err != nil {
			logger.Errorw("Unhandled error occurred",
				zap.String("method", c.Method()),
				zap.String("path", c.Path()),
				zap.Error(err),
			)
			return err
		}

		return nil
	}, mw.RequestLogger(logger))

	//groups
	userGroup := fiberApp.Group("/user")
	userHandlers.MapUserRoutes(userGroup, userHandler, mw)

	adminGroup := fiberApp.Group("/admin")
	adminHandlers.MapAdminRoutes(adminGroup, adminHandler, mw)

	groupGroup := fiberApp.Group("/group")
	groupHandlers.MapGroupRoutes(groupGroup, groupHandler)

	moderatorGroup := fiberApp.Group("/moderator")
	moderatorHandlers.MapModeratorRoutes(moderatorGroup, moderatorHandler, mw)

	scheduleGroup := fiberApp.Group("/schedule")
	scheduleHandlers.MapScheduleRoutes(scheduleGroup, scheduleHandler, mw)

	homeworkGroup := fiberApp.Group("/homework")
	homeworkStatusHandlers.MapHomeworkStatusRoutes(homeworkGroup, homeworkStatusHandler, mw)

	subjectGroup := fiberApp.Group("/subject")
	subjectHandlers.MapSubjectRoutes(subjectGroup, subjectHandler, mw)

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
