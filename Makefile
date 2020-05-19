
# git include
include ../packages/boilerplate/gitr.mk

# remove the "v" prefix
VERSION ?= $(shell echo $(TAGGED_VERSION) | cut -c 2-)



print: ## print

	$(MAKE) gitr-print


