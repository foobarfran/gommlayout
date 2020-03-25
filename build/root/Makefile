# Modules Config
APP_MODULES= cmd/fooapp 
LIB_MODULES= pkg/fooapi

# Modules Commands
BASE_DIR= $(shell pwd)
ALL_MODULES= $(APP_MODULES) $(LIB_MODULES)

default: clean build

clean: 
	@echo "\033[1mCLEAN\033[0m"
	@rm -rf target 

deps: set-deps $(ALL_MODULES)
test: set-test $(ALL_MODULES)
build-libs: set-build-lib $(LIB_MODULES)
build-apps: set-build-app $(APP_MODULES)
build: build-libs build-apps


set-build-lib set-build-app set-test set-deps:
	$(eval MAKE_TARGET := $(subst set-,,$@))

$(ALL_MODULES):
	@$(MAKE) $(MAKE_TARGET) -f ${BASE_DIR}/build/root/Makefile.module MODULE=$@ BASE_DIR=${BASE_DIR} -C ./$@ 


# Modules Release
release: check-release-version
	@for dir in $(ALL_MODULES); do \
		TAG_NAME=$$dir/v${RELEASE_VERSION}; \
		git tag $$TAG_NAME; \
		git push origin $$TAG_NAME; \
	done

check-release-version:
ifndef RELEASE_VERSION
	$(error RELEASE_VERSION flag is required to perform the release)
endif

.PHONY: clean deps test build build-libs build-apps release $(ALL_MODULES)
