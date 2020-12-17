.PHONY: $(MAKECMDGOALS)

$(MAKECMDGOALS):
	@cd $@ && gofmt -s -w . && goimports -w . && go vet .
	@cd $@ && go run .
