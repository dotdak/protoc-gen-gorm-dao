.PHONY: build

build:
	cd protoc-gen-gorm-dao && go build -o protoc-gen-gorm-dao && mv protoc-gen-gorm-dao ~/go/bin
	buf generate