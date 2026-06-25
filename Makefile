ROLE := backend-sdk
.PHONY: test test-standalone-layout test-go
test: test-standalone-layout test-go
test-standalone-layout:
	./test/scripts/assert-layout.sh $(ROLE)
# Go InferenceBackend contract (Go-only; the Python ABC was removed post-migration).
test-go:
	go test ./...
