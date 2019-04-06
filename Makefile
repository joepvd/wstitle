TARGET := wstitle
GOFILES := $(shell find . -name '*.go')
TESTS := $(shell find . -name '*_test.go')
SOURCEFILES := $(filter-out $(TESTS), $(GOFILES))

.PHONY: all
all: clean build

.PHONY: build
build: $(TARGET)

$(TARGET): $(SOURCEFILES)
	go build -o wstitle cmd/wstitle/main.go

.PHONY: clean
clean:
	rm -f $(TARGET)

.PHONY: test
test: fmt lint

.PHONY: fmt
fmt:
	gofmt -w $(GOFILES)
	git diff --quiet --exit-code

.PHONY: lint
lint: .ensure-lint
	golint -set_exit_status ./...

.PHONY: .ensure-lint
.ensure-lint:
	@if ! command -v go lint >/dev/null 2>&1; then go get -u golang.org/x/lint/golint; fi
