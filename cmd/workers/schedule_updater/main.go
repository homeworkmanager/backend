package main

import (
	"github.com/robfig/cron/v3"
	"homeworktodolist/internal/config"
	"homeworktodolist/internal/cron/schedule_updater"
	classRepo "homeworktodolist/internal/repository/postgres/class"
	groupRepo "homeworktodolist/internal/repository/postgres/group"
	homeworkRepo "homeworktodolist/internal/repository/postgres/homework"
	homeworkStatusRepo "homeworktodolist/internal/repository/postgres/homework_status"
	subjectRepo "homeworktodolist/internal/repository/postgres/subject"
	subjectNoteRepo "homeworktodolist/internal/repository/postgres/subjectnote"
	userRepo "homeworktodolist/internal/repository/postgres/user"
	userRedisRepo "homeworktodolist/internal/repository/redis/user"
	adminService "homeworktodolist/internal/service/admin"
	classService "homeworktodolist/internal/service/class"
	groupService "homeworktodolist/internal/service/group"
	homeworkService "homeworktodolist/internal/service/homework"
	homeworkStatusService "homeworktodolist/internal/service/homework_status"
	subjectService "homeworktodolist/internal/service/subject"
	subjectNoteService "homeworktodolist/internal/service/subjectnote"
	userService "homeworktodolist/internal/service/user"
	"homeworktodolist/internal/tx_manager"
	"homeworktodolist/pkg/db/postgres"
	"homeworktodolist/pkg/db/redis"
	"homeworktodolist/pkg/logger"
)

func main() {
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
	subjectNoteRepo := subjectNoteRepo.NewSubjectNoteRepo(txmanager)
	homeworkRepo := homeworkRepo.NewHomeworkRepo(txmanager)
	homeworkStatusRepo := homeworkStatusRepo.NewHomeworkStatusRepo(txmanager)

	//logger
	logger := logger.InitLogger(cfg)

	//Service
	userService := userService.NewUserService(userRepo, userRedisRepo, cfg)

	groupService := groupService.NewGroupService(groupRepo)

	subjectService := subjectService.NewSubjectService(subjectRepo)

	subjectNoteService := subjectNoteService.NewSubjectNoteService(subjectNoteRepo, subjectService)

	classService := classService.NewClassService(groupService, classRepo, subjectService, txmanager)

	homeworkStatusService := homeworkStatusService.NewHomeworkStatusService(homeworkStatusRepo)

	homeworkService := homeworkService.NewHomeworkService(homeworkRepo, homeworkStatusService, txmanager)

	adminService := adminService.NewAdminService(groupService, classService, subjectService, homeworkService, userService, subjectNoteService, txmanager)
	job := schedule_updater.NewCronJob(adminService, logger)

	c := cron.New()
	_, err := c.AddFunc("0 0 * * *", job.Run)
	logger.Info("Worker start")
	if err != nil {
		panic(err)
	}
	c.Start()
	defer c.Stop()

	select {}
}
