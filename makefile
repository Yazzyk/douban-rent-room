.PHONY: build build_linux build_windows build_darwin build_linux_arm build_darwin_arm build_docker

APPNAME = douban-rent-room

all: build

build:
	rm -rf ./build/
	make build_linux
	make build_windows
	make build_darwin
	make build_linux_arm
	make build_darwin_arm

build_linux:
	@echo "linux版"
	export GO111MODULE=on; \
	export GOPROXY="https://goproxy.cn,direct"; \
	export GOOS=linux; \
	export GOARCH=amd64; \
	go mod tidy -compat=1.20; \
	go build -o ./build/$(APPNAME)_linux_amd64/$(APPNAME) cmd/main.go
	cp config.toml ./build/$(APPNAME)_linux_amd64/config.toml

build_windows:
	@echo "windows版"
	export GO111MODULE=on; \
	export GOPROXY="https://goproxy.cn,direct"; \
	export GOOS=windows; \
	export GOARCH=amd64; \
	go mod tidy -compat=1.20; \
	go build -o ./build/$(APPNAME)_win_amd64/$(APPNAME) cmd/main.go
	cp config.toml ./build/$(APPNAME)_win_amd64/config.toml

build_darwin:
	@echo "darwin版"
	export GO111MODULE=on; \
	export GOPROXY="https://goproxy.cn,direct"; \
	export GOOS=darwin; \
	export GOARCH=amd64; \
	go mod tidy -compat=1.20; \
	go build -o ./build/$(APPNAME)_win_amd64/$(APPNAME) cmd/main.go
	cp config.toml ./build/$(APPNAME)_win_amd64/config.toml

build_linux_arm:
	@echo "linux arm版"
	export GO111MODULE=on; \
	export GOPROXY="https://goproxy.cn,direct"; \
	export GOOS=linux; \
	export GOARCH=arm; \
	go mod tidy -compat=1.20; \
	go build -o ./build/$(APPNAME)_linux_arm/$(APPNAME) cmd/main.go
	cp config.toml ./build/$(APPNAME)_linux_arm/config.toml

build_darwin_arm:
	@echo "darwin arm版"
	export GO111MODULE=on; \
	export GOPROXY="https://goproxy.cn,direct"; \
	export GOOS=darwin; \
	export GOARCH=arm64; \
	go mod tidy -compat=1.20; \
	go build -o ./build/$(APPNAME)_darwin_arm64/$(APPNAME) cmd/main.go
	cp config.toml ./build/$(APPNAME)_darwin_arm64/config.toml

build_docker:
	docker build -t $(APPNAME) .