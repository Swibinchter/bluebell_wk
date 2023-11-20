# makefile是自定义一系列shell命令，然后可以直接输入make xx 快捷调用

# .PHONY是告诉make工具遇到后面这些字段时需执行对应的命令，而不是去查找目录里是否有对应名字的文件
# 也叫伪目标，即代表是一个动作不是一个文件
.PHONY: all build run gotool clean

# 定义一个变量，可以在下面的用${BINARY}调用
BINARY="bluebell"

# all下的命令可以使用make直接执行，当然也可以使用make all
all: gotool build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -0 ${BINARY}

# 如果在名字前加@，表示执行时不会显示命令本身
run:
	@ go run ./main.go conf/config.yaml

gotool:
	go fmt ./
	go vet ./

clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

help:
	@echo "make - 格式化 Go 代码, 并编译生成二进制文件"
	@echo "make build - 编译 Go 代码, 生成二进制文件"
	@echo "make run - 直接运行 Go 代码"
	@echo "make clean - 移除二进制文件和 vim swap files"
	@echo "make gotool - 运行 Go 工具 'fmt' and 'vet'"