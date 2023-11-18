.PHONY: build

build:
	mkdir -p bin
	go build -o bin/protoc-gen-gorm-dao
	buf generate --template bootstrap.yaml --exclude-path examples/example_proto
	buf generate --exclude-path gorm/gorm.proto

bootstrap:
	buf generate --template bootstrap.yaml --exclude-path examples/example_proto