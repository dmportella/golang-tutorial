TEST?=$$(go list ./... | grep -v '/vendor/')
VETARGS?=-all
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
REV?=$$(git rev-parse --short HEAD)
BRANCH?=$$(git rev-parse --abbrev-ref HEAD)
VERSION?="0.0.0"

default: version fmt lint vet test

# Git commands
commit:
	@git add . 
	-@git rm $(git ls-files --others --deleted --exclude-standard)
	@git commit 
push:
	@git push origin ${BRANCH}

version:
	@echo "SOFTWARE VERSION"
	@echo "\tbranch:\t\t" ${BRANCH}
	@echo "\trevision:\t" ${REV}
	@echo "\tversion:\t" ${VERSION}

ci: tools build
	@echo "CI BUILD..."

tools:
	@echo "GO TOOLS installation..."
	@go get -u github.com/kardianos/govendor
	@go get -u golang.org/x/tools/cmd/cover
	@go get -u github.com/golang/lint/golint

build: version test
	@echo "GO BUILD..."
	@go build -ldflags "-X main.Build=${VERSION} -X main.Revision=${REV} -X main.Branch=${BRANCH}" -v -o golang-tutorial .

lint:
	@echo "GO LINT..."
	@golint

test: fmt generate lint vet
	@echo "GO TEST..."
	@go test $(TEST) $(TESTARGS) -timeout=30s -parallel=4 -bench=. -benchmem -cover

cover:
	@echo "GO TOOL COVER..."
	@go tool cover 2>/dev/null; if [ $$? -eq 3 ]; then \
		go get -u golang.org/x/tools/cmd/cover; \
	fi
	@go test $(TEST) -coverprofile=coverage.out
	@go tool cover -html=coverage.out
	@rm coverage.out

generate:
	@echo "GO GENERATE..."
	@go generate $(go list ./... | grep -v /vendor/)

# vet runs the Go source code static analysis tool `vet` to find
# any common errors.
vet:
	@echo "GO VET..."
	@go tool vet $(VETARGS) $$(ls -d */ | grep -v vendor) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

fmt:
	@echo "GO FMT..."
	@gofmt -w $(GOFMT_FILES)

.PHONY: tools default
