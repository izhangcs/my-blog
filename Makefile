shell := /bin/bash

versionDir := "zhangcs/blog/pkg/version"
gitTag = $(shell if [ "`git describe --tags --abbrev=0 2>/dev/null`" != "" ];then git describe --tags --abbrev=0; else git log --pretty=format:'%h' -n 1; fi)
gitCommit = $(shell git log --pretty=format:'%H' -n 1)
gitTreeState = $(shell if git status | grep -q 'git add';then echo dirty; else echo clean; fi)
buildDate = $(shell TZ=Asia/Shanghai date +%FT%T%z)

ldflags = "-w -X ${versionDir}.gitTag=${gitTag} -X ${versionDir}.gitCommit=${gitCommit} -X ${versionDir}.gitTreeState=${gitTreeState} -X ${versionDir}.buildDate=${buildDate}"

all: clean gotool swag
	go build -v -ldflags ${ldflags} .

clean:
	rm -f blog

gotool:
	gofmt -w .
	go vet . 

ca:
	openssl req -new -nodes -x509 -out conf/server.crt -keyout conf/server.key -days 3650 -subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=127.0.0.1/emailAddress=xxxxx@qq.com"

swag:
	swag init

help:
	@echo "make - compile the source code"
	@echo "make clean - remove binary file and vim swp files"
	@echo "make gotool - run go tool 'fmt' and 'vet'"
	@echo "make ca - generate ca files"

.PHONY: clean gotool ca 