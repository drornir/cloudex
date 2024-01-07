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

.PHONY: build 
build: # build into bin directory
	mkdir -p bin
	go build -o bin/${BINARY} .

.PHONY: run
run: ## runs the build binary
	bin/${BINARY}

##@ Development

.PHONY: test
test: build ## go test ./...
	go test ./...


.PHONY: dev
dev: gen fmt lint build run ## Main dev command. Run `make setup-dev` once before runnign this

.PHONY: dev-run
dev-run: ## Runs the server with templ watch
	templ generate --watch --cmd="$(MAKE) build run"

.PHONY: setup-dev
setup-dev: upgrade-go-tools goget ## setup dev env


.PHONY: goget
goget: ## go get ./...
	go get ./...
	go mod tidy

.PHONY: goget-u
goget-u: ## go get -u ./...
	go get -u ./...

.PHONY: upgrade-app
upgrade-app: upgrade-go goget-u ## update all dependencies to latest
	go mod tidy

.PHONY: upgrade-global
upgrade-global: upgrade-go-tools upgrade-prettier brew-install ## upgrade global dev tooling

.PHONY: upgrade-go
upgrade-go: ## upgrade to latest go and toolchain
	go get go@latest
	go get toolchain@patch

.PHONY: upgrade-prettier
upgrade-prettier: ## upgrade prettier and plugin
	npm update --include=dev prettier prettier-plugin-tailwindcss

.PHONY: upgrade-go-tools
upgrade-go-tools: ## install dev deps 
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/drornir/factor3@latest
	go install golang.org/x/tools/cmd/stringer@latest
	go install github.com/campoy/jsonenums@latest
	go install github.com/a-h/templ/cmd/templ@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

.PHONY: brew-install
brew-install: 
	brew install tailwindcss


.PHONY: gen 
gen: gen-go sqlc tailwind templ fmt ## all code generation scripts

.PHONY: gen-go
gen-go:
	go generate ./...

.PHONY: templ
templ:
	templ generate

.PHONY: sqlc
sqlc:
	rm pkg/db/*.sql.go || true
	sqlc generate --file pkg/db/sqlc.yaml

.PHONY: tailwind
tailwind: 
	tailwindcss -i css/main.css -o assets/main.css


.PHONY: lint
lint: ## go vet
	go vet ./...

.PHONY: fmt
fmt: ## goimports 
	goimports -w .

