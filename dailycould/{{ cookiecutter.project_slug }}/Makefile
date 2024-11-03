BIN := {{ cookiecutter.project_slug }}
PREFIX := github.com/{{ cookiecutter.github_username }}/$(BIN)/version

SRC := $(shell find . -name '*.go')

DATE := $(shell date +"%Y-%m-%dT%H:%M:%SZ")
GOVERSION := $(shell go version)
VERSION := $(shell git describe --tags --abbrev=8 --dirty --always --long)
SHORT_SHA := $(shell git rev-parse --short HEAD)
FULL_SHA := $(shell git rev-parse HEAD)
export GOVERSION # goreleaser wants this

LDFLAGS = -s -w
LDFLAGS += -X $(PREFIX).Version=$(VERSION)
LDFLAGS += -X '$(PREFIX).Date=$(DATE)'
LDFLAGS += -X '$(PREFIX).GoVersion=$(GOVERSION)'
LDFLAGS += -X $(PREFIX).ShortGitSHA=$(SHORT_SHA)
LDFLAGS += -X $(PREFIX).FullGitSHA=$(FULL_SHA)

.DEFAULT_GOAL := iterate

# Directory to store timestamps for tracking operations
STAMPS := .stamps

$(STAMPS):
	mkdir -p $(STAMPS)

.PHONY: all # Build, check and install
all: check $(BIN) install

.PHONY: iterate # Run check and build
iterate: check $(BIN)

.PHONY: check # Run all code quality checks
check: $(STAMPS)/check

# Aggregate timestamp for all checks
$(STAMPS)/check: $(STAMPS)/goimports $(STAMPS)/tidy $(STAMPS)/fmt $(STAMPS)/lint $(STAMPS)/vet | $(STAMPS)
	@touch $@

.PHONY: build # Build the binary
build: $(BIN)

$(BIN): $(SRC) go.mod go.sum
	go build -ldflags "$(LDFLAGS)" -o $@

.PHONY: goreleaser # Run goreleaser with clean
goreleaser:
	goreleaser --clean

# Individual check timestamps
$(STAMPS)/goimports: $(SRC) | $(STAMPS)
	goimports -w $(SRC)
	goimports-reviser -output=file -set-alias -rm-unused -format $(SRC)
	@touch $@

$(STAMPS)/tidy: go.mod go.sum | $(STAMPS)
	go mod tidy
	@touch $@

$(STAMPS)/fmt: $(SRC) | $(STAMPS)
	gofumpt -extra -w $(SRC)
	@touch $@

$(STAMPS)/golines: $(SRC) | $(STAMPS)
	golines -w $(SRC)
	@touch $@

$(STAMPS)/lint: $(SRC) | $(STAMPS)
	golangci-lint run
	@touch $@

$(STAMPS)/vet: $(SRC) | $(STAMPS)
	go vet ./...
	@touch $@

.PHONY: test # Run tests
test:
	go test ./...

.PHONY: install # Install the binary
install: $(BIN)
	go install -ldflags "$(LDFLAGS)"

.PHONY: help # Show this help
help:
	@grep '^.PHONY: .* #' Makefile | sed 's/\.PHONY: \(.*\) # \(.*\)/\1 \2/' | expand -t20

.PHONY: clean # Remove built binary and stamps
clean:
	$(RM) -r $(BIN) $(STAMPS)

# Convenience targets that force execution regardless of timestamps
.PHONY: goimports tidy fmt golines lint vet
goimports: | $(STAMPS)
	@rm -f $(STAMPS)/goimports
	@$(MAKE) $(STAMPS)/goimports

tidy: | $(STAMPS)
	@rm -f $(STAMPS)/tidy
	@$(MAKE) $(STAMPS)/tidy

fmt: | $(STAMPS)
	@rm -f $(STAMPS)/fmt
	@$(MAKE) $(STAMPS)/fmt

golines: | $(STAMPS)
	@rm -f $(STAMPS)/golines
	@$(MAKE) $(STAMPS)/golines

lint: | $(STAMPS)
	@rm -f $(STAMPS)/lint
	@$(MAKE) $(STAMPS)/lint

vet: | $(STAMPS)
	@rm -f $(STAMPS)/vet
	@$(MAKE) $(STAMPS)/vet