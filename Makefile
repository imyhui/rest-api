SHELL := /bin/bash
BASEDIR = $(shell pwd)

# build with verison infos
versionDir = "rest-api/pkg/version"
gitTag = $(shell if [ "`git describe --tags --abbrev=0 2>/dev/null`" != "" ];then git describe --tags --abbrev=0; else git log --pretty=format:'%h' -n 1; fi)
buildDate = $(shell TZ=Asia/Shanghai date +%FT%T%z)
gitCommit = $(shell git log --pretty=format:'%H' -n 1)
gitTreeState = $(shell if git status|grep -q 'clean';then echo clean; else echo dirty; fi)

ldflags="-w -X ${versionDir}.gitTag=${gitTag} -X ${versionDir}.buildDate=${buildDate} -X ${versionDir}.gitCommit=${gitCommit} -X ${versionDir}.gitTreeState=${gitTreeState}"

all: fmt
	@go build -v -ldflags ${ldflags} .
clean:
	rm -f rest-api
	find . -name "[._]*.s[a-w][a-z]" | xargs rm -f
fmt:
	gofmt -w .
	go vet . 2>&1 | grep -v vendor || true
ca:
	openssl req -new -nodes -x509 -out conf/server.crt -keyout conf/server.key -days 3650 -subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=127.0.0.1/emailAddress=xxxxx@qq.com"
doc:
	swag init
help:
	@echo "make - compile the source code"
	@echo "make clean - remove binary file and vim swp files"
	@echo "make fmt - run go tool 'fmt' and 'vet'"
	@echo "make ca - generate ca files"
	@echo "make doc - generate doc files"

.PHONY: clean fmt ca help doc