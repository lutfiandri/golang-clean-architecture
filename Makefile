.PHONY: test clean coverage

TEST_TARGET ?= ./test/http/...

test:
	go test -v -coverprofile=./test/coverage.out $(TEST_TARGET) \
		./internal/delivery/http/controller \
		./internal/usecase \
		./internal/repository
	go tool cover -html ./test/coverage.out -o ./test/coverage.html