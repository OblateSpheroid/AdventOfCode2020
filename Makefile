DEP_BIN := $(shell go env GOPATH)/bin/goimports
.PHONY: $(MAKECMDGOALS)

${DEP_BIN}: # ensure dependencies are installed
	@go get golang.org/x/tools/cmd/goimports

$(MAKECMDGOALS): ${DEP_BIN}
	@cd $@ && gofmt -s -w . && goimports -w . && go vet .
	@cd $@ && go run .
