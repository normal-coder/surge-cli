.PHONY: all build compress macos linux windows compress-macos compress-linux compress-windows help

BUILD_PATH=./build
APPNAME=surge

all: build compress
build: macos linux windows
compress: compress-macos compress-linux compress-windows

macos:
	@echo "开始 macOS 可执行程序编译..."
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o $(BUILD_PATH)/$(APPNAME)-darwin-x86_64
	@echo "macOS 可执行程序编译完成..."

install:
	@echo "Installing..."
	cp build/surge-darwin-x86_64 /Users/normal/.dotfiles/bash-extend/surge
	@echo "Installed at /Users/normal/.dotfiles/bash-extend/surge"

linux:
	@echo "开始 Linux 可执行程序编译..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $(BUILD_PATH)/$(APPNAME)-linux-x86_64
	@echo "Linux 可执行程序编译完成..."

windows:
	@echo "开始 Windows 可执行程序编译..."
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o $(BUILD_PATH)/$(APPNAME)-windows-x86_64.exe
	@echo "Windows 可执行程序编译完成..."

compress-macos:
	@echo "macOS 可执行程序压缩开始..."
	mv $(BUILD_PATH)/$(APPNAME)-darwin-x86_64 $(BUILD_PATH)/$(APPNAME)-darwin-x86_64_tmp
	upx --best -o $(BUILD_PATH)/$(APPNAME)-darwin-x86_64 $(BUILD_PATH)/$(APPNAME)-darwin-x86_64_tmp
	rm -fr $(BUILD_PATH)/$(APPNAME)-darwin-x86_64_tmp
	@echo "macOS 可执行程序压缩完成..."

compress-linux:
	@echo "Linux 可执行程序压缩开始..."
	mv $(BUILD_PATH)/$(APPNAME)-linux-x86_64 $(BUILD_PATH)/$(APPNAME)-linux-x86_64_tmp
	upx --best -o $(BUILD_PATH)/$(APPNAME)-linux-x86_64 $(BUILD_PATH)/$(APPNAME)-linux-x86_64_tmp
	rm -fr $(BUILD_PATH)/$(APPNAME)-linux-x86_64_tmp
	@echo "Linux 可执行程序压缩完成..."

compress-windows:
	@echo "Windows 可执行程序压缩开始..."
	mv $(BUILD_PATH)/$(APPNAME)-windows-x86_64.exe $(BUILD_PATH)/$(APPNAME)-windows-x86_64.exe_tmp
	upx --best -o $(BUILD_PATH)/$(APPNAME)-windows-x86_64.exe $(BUILD_PATH)/$(APPNAME)-windows-x86_64.exe_tmp
	rm -fr $(BUILD_PATH)/$(APPNAME)-windows-x86_64.exe_tmp
	@echo "Windows 可执行程序压缩完成..."

help:
	@echo "构建命令集，用于快速构建服务。"
	@echo ""
	@echo "用法："
	@echo ""
	@echo "	make <command>"
	@echo ""
	@echo "可用命令如下：:"
	@echo "	make all              格式化代码, 编译生成并压缩全部平台的可执行程序"
	@echo "	make build            编译生成全部可执行程序"
	@echo "	make compress         压缩全部可执行程序"
	@echo "	make macos      构建 masOS 可执行程序"
	@echo "	make linux      构建 Linux 可执行程序"
	@echo "	make windows    构建 Windows 可执行程序"
	@echo "	make compress-macos   压缩 masOS 可执行程序"
	@echo "	make compress-linux   压缩 Linux 可执行程序"
	@echo "	make compress-windows 压缩 Windows 可执行程序"
	@echo "	make help             查看帮助"
