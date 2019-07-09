.POSIX:

SHELL:=/bin/sh
VERSION:=`cat ./VERSION.txt`
GOCMD:=go
TF_PLUGINS:=${HOME}/.terraform.d/plugins
WEBSITE_REPO:=github.com/hashicorp/terraform-website
PKG_NAME:=routeros

.PHONY: default
default: all

.PHONY: all
all: install


.PHONY: install
install: dist
	@echo "Locally installing terraform provider plugins into \"${TF_PLUGINS}/\"..."
	@$(GOCMD) install .
	@cp "./.bin/linux_amd64/terraform-provider-routeros_v$(VERSION)_x4" "$(TF_PLUGINS)/linux_amd64/terraform-provider-routeros_v$(VERSION)_x4"
	@cp "./.bin/darwin_amd64/terraform-provider-routeros_v$(VERSION)_x4" "$(TF_PLUGINS)/darwin_amd64/terraform-provider-routeros_v$(VERSION)_x4"
	@cp "./.bin/windows_amd64/terraform-provider-routeros_v$(VERSION)_x4.exe" "$(TF_PLUGINS)/windows_amd64/terraform-provider-routeros_v$(VERSION)_x4.exe"

.PHONY: clean
clean:
	@$(GOCMD) clean
	@# All auto-generated files have been prefixed with zzz_
	@echo "Removing auto-generated sources..."
	@rm -f "./routeros/zzz_"*".go"
	@echo "Removing output directory \"./.bin/\"..."
	@rm -rf "./.bin/"

.PHONY: prepare
prepare:
	@mkdir -p "./.bin/linux_amd64"
	@mkdir -p "$(TF_PLUGINS)/linux_amd64"
	@mkdir -p "./.bin/darwin_amd64"
	@mkdir -p "$(TF_PLUGINS)/darwin_amd64"
	@mkdir -p "./.bin/windows_amd64"
	@mkdir -p "$(TF_PLUGINS)/windows_amd64"
	@echo "Building code generator..."
	@$(GOCMD) build -o ./.bin/terraform-routeros-binding-generator ./generator/

.PHONY: generate
generate: prepare
	@$(GOCMD) generate ./...
	@$(GOCMD) fmt ./routeros/zzz_*.go > /dev/null

.PHONY: dist
dist: generate
	@GOOS=linux GOARCH=amd64 $(GOCMD) build -o ./.bin/linux_amd64/terraform-provider-routeros_v$(VERSION)_x4 .
	@GOOS=darwin GOARCH=amd64 $(GOCMD) build -o ./.bin/darwin_amd64/terraform-provider-routeros_v$(VERSION)_x4 .
	@GOOS=windows GOARCH=amd64 $(GOCMD) build -o ./.bin/windows_amd64/terraform-provider-routeros_v$(VERSION)_x4.exe .

.PHONY: website
website:
ifeq (,$(wildcard $(GOPATH)/src/$(WEBSITE_REPO)))
	echo "$(WEBSITE_REPO) not found in your GOPATH (necessary for layouts and assets), get-ting..."
	git clone https://$(WEBSITE_REPO) $(GOPATH)/src/$(WEBSITE_REPO)
endif
	@$(MAKE) -C $(GOPATH)/src/$(WEBSITE_REPO) website-provider PROVIDER_PATH=$(shell pwd) PROVIDER_NAME=$(PKG_NAME)

website-test:
ifeq (,$(wildcard $(GOPATH)/src/$(WEBSITE_REPO)))
	echo "$(WEBSITE_REPO) not found in your GOPATH (necessary for layouts and assets), get-ting..."
	git clone https://$(WEBSITE_REPO) $(GOPATH)/src/$(WEBSITE_REPO)
endif
	@$(MAKE) -C $(GOPATH)/src/$(WEBSITE_REPO) website-provider-test PROVIDER_PATH=$(shell pwd) PROVIDER_NAME=$(PKG_NAME)
