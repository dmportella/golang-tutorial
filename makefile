TEST?=$$(go list ./... | grep -v '/vendor/')
VETARGS?=-all
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
REV?=$$(git rev-parse --short HEAD)
BRANCH?=$$(git rev-parse --abbrev-ref HEAD)
VERSION?="0.0.0"

default: version fmt lint vet test

version:
	@echo "branch: " ${BRANCH}
	@echo "revision: " ${REV}
	@echo "version: " ${VERSION}

ci: tools build
	@echo "travis building..."

tools:
	@echo "installing tools..."
	@go get -u github.com/kardianos/govendor
	@go get -u golang.org/x/tools/cmd/cover
	@go get -u github.com/golang/lint/golint

build: version test
	@echo "building..."
	@go build -ldflags "-X main.Build=${VERSION} -X main.Revision=${REV} -X main.Branch=${BRANCH}" -v -o golang-tutorial .

lint:
	@echo "lint..."
	@golint

test: fmt generate lint vet
	@echo "testing..."
	@go test $(TEST) $(TESTARGS) -timeout=30s -parallel=4 -bench=. -benchmem -cover

cover:
	@echo "cover report..."
	@go tool cover 2>/dev/null; if [ $$? -eq 3 ]; then \
		go get -u golang.org/x/tools/cmd/cover; \
	fi
	@go test $(TEST) -coverprofile=coverage.out
	@go tool cover -html=coverage.out
	@rm coverage.out

generate:
	@echo "generating..."
	@go generate $(go list ./... | grep -v /vendor/)

# vet runs the Go source code static analysis tool `vet` to find
# any common errors.
vet:
	@echo "go vet..."
	@echo "go tool vet $(VETARGS) ."
	@go tool vet $(VETARGS) $$(ls -d */ | grep -v vendor) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

fmt:
	@echo "gotfmt..."
	@gofmt -w $(GOFMT_FILES)

.PHONY: tools default
