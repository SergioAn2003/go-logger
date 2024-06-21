.PHONY: doc lint test test-cover test-cover-svg test-cover-html install-deps
.SILENT:


# Show godoc in browser
doc:
	@echo documentation is available at http://localhost:6060/pkg/gitlab.app.cube/medcore/go-utils/logger/
	@$(CURDIR)/temp/bin/godoc -http=localhost:6060



# Lint
lint:
	@$(CURDIR)/temp/bin/golangci-lint run -c ./.golangci.yaml --path-prefix . --fix



# Test
test:
	@go test --cover --coverprofile=$(CURDIR)/temp/coverage.out $(TEST_COVER_EXCLUDE_DIR) --race
test-cover:
	@go tool cover -func=$(CURDIR)/temp/coverage.out | grep total | grep -oE '[0-9]+(\.[0-9]+)?%'
test-cover-svg:
	@$(CURDIR)/temp/bin/go-cover-treemap -coverprofile $(CURDIR)/temp/coverage.out > $(CURDIR)/temp/coverage.svg
test-cover-html:
	@go tool cover -html=$(CURDIR)/temp/coverage.out



# Dependencies
install-deps:
	@GOBIN=$(CURDIR)/temp/bin go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@GOBIN=$(CURDIR)/temp/bin go install golang.org/x/tools/cmd/godoc@latest
	@GOBIN=$(CURDIR)/temp/bin go install github.com/nikolaydubina/go-cover-treemap@latest
	@GOBIN=$(CURDIR)/temp/bin go install go.uber.org/mock/mockgen@latest
	@go mod tidy



# Config
TEST_COVER_EXCLUDE_DIR := `go list ./... | grep -v -E '/cmd/|/mocks/|/bin'`
