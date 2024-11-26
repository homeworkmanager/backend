
MIGRATION_FOLDER = ./internal/db/migrations

local_migration_up:
	goose -dir $(MIGRATION_FOLDER) postgres "user=postgres password=admin database=homeworkmanager host=192.168.0.69 port=5432" up

local_migration_down:
	goose -dir $(MIGRATION_FOLDER) postgres "user=postgres password=admin database=homeworkmanager host=192.168.0.69 port=5432" down




