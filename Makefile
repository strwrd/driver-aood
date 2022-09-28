tests:
	go test ./... -count=1
.PHONY: tests

tests-cover:
	go test ./... -count=1 -coverprofile coverage.out
.PHONY: tests