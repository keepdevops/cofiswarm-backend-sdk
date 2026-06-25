// Package backend is the Go InferenceBackend contract — every engine (llama.cpp, MLX,
// vLLM, ...) implements it. Go port of the Python cofiswarm_backend.base ABC; the two
// coexist (like cofiswarm-observer-sdk's dual Go/Python clients) during the MLX-cluster
// migration to Go.
package backend

import "context"

// DefaultMaxTokens mirrors GenerateRequest.max_tokens=512.
const DefaultMaxTokens = 512

// GenerateRequest is one generation request (ports GenerateRequest).
type GenerateRequest struct {
	Prompt      string
	MaxTokens   int // 0 => DefaultMaxTokens (the backend applies the default)
	Temperature float64
	Stop        []string
	Extra       map[string]any
}

// TokenChunk is one streamed unit; the final chunk has Done=true (ports TokenChunk).
type TokenChunk struct {
	Text string
	Done bool
	Meta map[string]any
}

// HealthStatus is a liveness probe result (ports HealthStatus).
type HealthStatus struct {
	OK     bool
	Detail string
}

// InferenceBackend is the common surface for all inference engines.
//
// GenerateStream calls emit for each TokenChunk (the Go idiom for Python's
// AsyncIterator[TokenChunk]); the final chunk has Done=true. If emit returns an error,
// GenerateStream stops and returns it (caller-driven cancellation, e.g. client disconnect).
type InferenceBackend interface {
	GenerateStream(ctx context.Context, req GenerateRequest, emit func(TokenChunk) error) error
	Embed(ctx context.Context, texts []string) ([][]float32, error)
	Health(ctx context.Context) HealthStatus
	Close() error
}
