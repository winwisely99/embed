GO_OS 		:= $(shell go env GOOS)
GO_ARCH		:= $(shell go env GOARCH)

GIT_VERSION	:= $(shell git describe --tags)

# Signing
AND_KEYSTORE_APP_NAME=main
AND_KEYSTORE_FSPATH=$(PWD)/_keys/$(AND_KEYSTORE_APP_NAME)

DEP_LIB=github.com/DimitarPetrov/stegify
DEP_FSAPTH=$(GOAPTH)/src/$(DEP_LIB)

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
	@echo SIGNING --
	@echo AND_KEYSTORE_APP_NAME: $(AND_KEYSTORE_APP_NAME)
	@echo AND_KEYSTORE_FSPATH: $(AND_KEYSTORE_FSPATH)
	@echo



bootstrap:
	# move to bootstrap repo later
	
dep:
	go get -u github.com/DimitarPetrov/stegify



