## 数据库链接
DB_URL=postgresql://postgres:postgres_booking@localhost:5555/db_bookings?sslmode=disable

## migrate: 生成迁移sql文件, 如: make migrate NAME=init
migrate:
	migrate create -seq -ext=.sql -dir=./dbrepo/migrations $$NAME

## migrate_up: 向上迁移
migrate_up:
	migrate -database ${DB_URL} -path ./dbrepo/migrations -verbose up 1

## migrate_down: 向下迁移
migrate_down:
	migrate -database ${DB_URL} -path ./dbrepo/migrations -verbose down 1

## migrate_force: 版本回滚，例如： make migrate_force V=1
migrate_force:
	migrate -database ${DB_URL} -path ./dbrepo/migrations force $$V

## null: 替换sql.NullXXX 为指针类型
null:
	go run ./cmd/replace_null/*.go
	goimports -w ./dbrepo/postgres

## sqlc: 使用 sqlc generate 生成代码
sqlc:
	sqlc generate
	make null
	
## run: 启动api服务
run: 
	go run ./cmd/api/*.go

