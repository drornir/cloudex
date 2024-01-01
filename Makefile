SHELL:=/bin/zsh

BINARY=cloudex

default: help


##@ General

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk commands is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php

.PHONY: help
help: ## Show help for each of the Makefile recipes.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<command>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: dev
dev: gen fmt lint build run ## Main dev command. Run `make setup-dev` once before runnign this

.PHONY: dev-run
dev-run: ## Runs the server with templ hot reload
	templ generate --watch --cmd="make build run"

.PHONY: test
test: build ## go test ./...
	go test ./...

.PHONY: setup-dev
setup-dev: install-dev-deps goget ## setup dev env
	
.PHONY: goget
goget: ## go get ./...
	go get ./...
	go mod tidy

.PHONY: install-dev-deps
install-dev-deps: ## install dev deps 
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/drornir/factor3@latest
	go install golang.org/x/tools/cmd/stringer@latest
	go install github.com/campoy/jsonenums@latest
	go install github.com/a-h/templ/cmd/templ@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

.PHONY: gen 
gen: go-gen sqlc templ ## all code generation scripts

.PHONY: go-gen
go-gen:
	go generate ./...

.PHONY: templ
templ:
	templ generate

.PHONY: sqlc
sqlc:
	rm pkg/db/*.sql.go
	sqlc generate --file pkg/db/sqlc.yaml
	

.PHONY: build 
build: # build into bin directory
	mkdir -p bin
	go build -o bin/${BINARY} .

.PHONY: lint
lint: ## go vet
	go vet ./...

.PHONY: fmt
fmt: ## goimports 
	goimports -w .

.PHONY: run
run: ## runs the build binary
	bin/${BINARY}