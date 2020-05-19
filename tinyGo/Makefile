
# git include
include ../../packages/boilerplate/help.mk
include ../../packages/boilerplate/gitr.mk
include ../../packages/boilerplate/gitu.mk


# my env.mk
#include $(PWD)/env.mk

LIB_FSPATH :=?


print: ## print

	@echo LIB_BIN_NAME : $(LIB_BIN_NAME)

	$(MAKE) gitr-print

	$(MAKE) gitu-print

GIT_BRANCH = blah
dep-t:
    $(eval GIT_BRANCH=$(shell git rev-parse --abbrev-ref HEAD))
    @echo Git branch is $(GIT_BRANCH)
	
#VAR1 =: 0
dep:
	#VAR1=$(shell $(MAKE) gitu-clone-master)
	#VAR1 = $(MAKE) gitu-clone-master
	@echo VAR1: $(VAR1)


test:
	go test -v ./...

	
### cloud

cloud-auth:
	gcloud auth login
	gcloud config set project $(CLOUD_PROJECT_ID)

	gcloud config set run/region europe-west1
cloud-open:
	# run
	open https://console.cloud.google.com/run?project=$(CLOUD_PROJECT_ID)
	# container
	open https://console.cloud.google.com/cloud-build/builds?project=$(CLOUD_PROJECT_ID)
	# store
	open https://console.cloud.google.com/storage/browser?project=$(CLOUD_PROJECT_ID)
cloud-build:
	echo $(CLOUD_PROJECT_ID)
	cd $(LIB_FSPATH) && gcloud builds submit --tag gcr.io/$(CLOUD_PROJECT_ID)/identicon-generator
cloud-deploy:
	cd $(LIB_FSPATH) && gcloud beta run deploy --image gcr.io/$(CLOUD_PROJECT_ID)/identicon-generator --platform managed




### build

build: ## build
	go build -o $(LIB_BIN_FSPATH) .

build-tools:
	go build -tags tools -o $(LIB_BIN_FSPATH) .

go-clean: ## go-clean
	# go stuff
	go clean -cache
	go clean -modcache
	go clean -testcache

	# binary in the GO BIN path
	rm $(LIB_BIN_FSPATH)



