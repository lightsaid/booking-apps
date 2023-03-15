gen:
	protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=.

server:
	go run ./cmd/server/*.go

client:
	go run ./cmd/client/*.go

## replace: 替换sql.NullXXX 为指针类型
replace:
	go run ./cmd/replace/*.go
	goimports -w ./db/postgres

## sqlc: 使用 sqlc generate 生成代码
sqlc:
	sqlc generate
	make replace
	

.PHONY: gen server client replace sqlc