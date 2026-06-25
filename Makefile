ROLE := backend-sdk
.PHONY: test test-standalone-layout test-go
test: test-standalone-layout test-go
test-standalone-layout:
	./test/scripts/assert-layout.sh $(ROLE)
# Go InferenceBackend contract (coexists with the Python ABC during the cluster migration).
test-go:
	go test ./...
