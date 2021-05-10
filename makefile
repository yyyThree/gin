.PHONY: all env vendor build clean help

PROJECT_NAME="gin_demo"
INSTALL_DIR="/tmp/go"
CONFIG_FILE=${INSTALL_DIR}/config.yaml
MIGRATION_DIR=${INSTALL_DIR}/migration
LOG=/tmp/go/${PROJECT_NAME}.log

all: build

env:
    go env -w GO111MODULE=on
    go env -w GOPROXY=https://goproxy.cn,direct

vendor: env
    go mod vendor

build: clean vendor
    @if [ ! -d ${INSTALL_DIR} ] ; then mkdir ${INSTALL_DIR} ; fi
    cp -f config.yaml ${CONFIG_FILE}
    cp -rf migration ${MIGRATION_DIR}
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${INSTALL_DIR}/${PROJECT_NAME}

clean:
    @if [ -f ${INSTALL_DIR}/${PROJECT_NAME} ] ; then rm ${INSTALL_DIR}/${PROJECT_NAME} ; fi

start:
    @if [ ! -f ${INSTALL_DIR}/${PROJECT_NAME} ] ; then make ; fi
    @sh -c "cd ${INSTALL_DIR} && nohup ./${PROJECT_NAME} > ${LOG} 2>&1 &"

help:
    @echo "make         - 删除旧的二进制文件，重新编译 Go 代码, 生成二进制文件"
    @echo "make env     - 设置go环境变量"
    @echo "make vendor  - 生成依赖包"
    @echo "make build   - 同make"
    @echo "make clean   - 移除二进制文件"
    @echo "make start   - 启动服务"