VERSION = 1.0.0
PACKAGE = wechat
GOPATH  = $(CURDIR)/vendor:$(CURDIR)
BASE    = $(CURDIR)/src/$(PACKAGE)
DATE   ?= $(shell date +%FT%T%z)

export GOPATH

M = $(shell printf "\033[34;1m▶\033[0m")

.PHONY: dep
dep: ; $(info $(M) retrieving dependencies…)
	go get -d -v ./src/...


.PHONY: fmt
fmt: ; $(info $(M) running gofmt…) @ ## Run gofmt on all source files
	@ret=0 && for d in $$(go list -f '{{.Dir}}' ./...); do \
		gofmt -l -w $$d/*.go || ret=$$? ; \
	 done ; exit $$ret


.PHONY: clean
clean: ; $(info $(M) cleaning…)	@ ## Cleanup everything
	@rm -rf bin
	@rm -rf vendor


.PHONY: all
all: fmt dep | $(BASE) ; $(info $(M) building executable…) @ ## Build program binary
	$Q cd $(BASE) && go build \
		-tags release \
		-ldflags '-X $(PACKAGE)/cmd.Version=$(VERSION) -X $(PACKAGE)/cmd.BuildDate=$(DATE)' \
		-o $(CURDIR)/bin/$(PACKAGE) main.go
