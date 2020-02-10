GO_OS 		:= $(shell go env GOOS)
GO_ARCH		:= $(shell go env GOARCH)

GIT_VERSION	:= $(shell git describe --tags)

# Signing
AND_KEYSTORE_APP_NAME=main
AND_KEYSTORE_FSPATH=$(PWD)/_keys/$(AND_KEYSTORE_APP_NAME)

DEP_LIB=github.com/vishen/go-chromecast
DEP_LIB_FSAPTH=$(GOPATH)/src/$(DEP_LIB)

print:

	@echo
	@echo GO_OS: $(GO_OS)
	@echo GO_OS: $(GO_ARCH)
	@echo

	@echo
	@echo GIT_VERSION: $(GIT_VERSION)
	@echo GIT_OTHER: $(shell git describe --all)
	@echo

	@echo
	@echo DEP_LIB: $(DEP_LIB)
	@echo DEP_FSAPTH: $(DEP_LIB_FSAPTH)
	@echo

	@echo
	@echo SIGNING --
	@echo AND_KEYSTORE_APP_NAME: $(AND_KEYSTORE_APP_NAME)
	@echo AND_KEYSTORE_FSPATH: $(AND_KEYSTORE_FSPATH)
	@echo




bootstrap:
	# move to bootstrap repo later

dep-save:
	go mod init
	go mod vendor
run:
	go run main.go
build:
	go build -o $(PWD)/bin/cast .

run-gui:
	./bin/cast ui
run-ls:
	./bin/cast ls
run-stat:
	./bin/cast status

	#go-chromecast status -u b03f1b498b8cf803ffe4da91d97f055e
run-vid:
	./bin/cast load https://www.youtube.com/watch?v=kWxnadQI5Qw -u b03f1b498b8cf803ffe4da91d97f055e

