.PHONY: build

build:
	mkdir -p bin
	cd protoc-gen-gorm-dao && go build -o protoc-gen-gorm-dao && mv protoc-gen-gorm-dao ../bin/
	buf generate --template bootstrap.yaml --exclude-path examples/example_proto
	buf generate --exclude-path gorm/gorm.proto

bootstrap:
	buf generate --template bootstrap.yaml --exclude-path examples/example_proto