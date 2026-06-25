package backend

import (
	"context"
	"testing"
)

// stubBackend confirms the interface is implementable and the streaming contract works.
type stubBackend struct{ id string }

func (s *stubBackend) GenerateStream(_ context.Context, req GenerateRequest, emit func(TokenChunk) error) error {
	max := req.MaxTokens
	if max == 0 {
		max = DefaultMaxTokens
	}
	for _, t := range []string{"hello", " world"} {
		if err := emit(TokenChunk{Text: t}); err != nil {
			return err
		}
	}
	return emit(TokenChunk{Done: true, Meta: map[string]any{"max_tokens": max}})
}

func (s *stubBackend) Embed(_ context.Context, texts []string) ([][]float32, error) {
	out := make([][]float32, len(texts))
	for i := range texts {
		out[i] = []float32{1, 0}
	}
	return out, nil
}

func (s *stubBackend) Health(context.Context) HealthStatus { return HealthStatus{OK: true, Detail: s.id} }
func (s *stubBackend) Close() error                        { return nil }

func TestInterfaceContract(t *testing.T) {
	var b InferenceBackend = &stubBackend{id: "stub"}

	var got []string
	var last TokenChunk
	err := b.GenerateStream(context.Background(), GenerateRequest{Prompt: "hi"}, func(c TokenChunk) error {
		if c.Text != "" {
			got = append(got, c.Text)
		}
		last = c
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 || got[0] != "hello" {
		t.Errorf("streamed = %v", got)
	}
	if !last.Done || last.Meta["max_tokens"] != DefaultMaxTokens {
		t.Errorf("final chunk = %+v (want Done + default max_tokens)", last)
	}

	if v, _ := b.Embed(context.Background(), []string{"a", "b"}); len(v) != 2 {
		t.Errorf("embed len = %d", len(v))
	}
	if h := b.Health(context.Background()); !h.OK || h.Detail != "stub" {
		t.Errorf("health = %+v", h)
	}

	// emit-error propagation (caller cancellation).
	want := context.Canceled
	gotErr := b.GenerateStream(context.Background(), GenerateRequest{}, func(TokenChunk) error { return want })
	if gotErr != want {
		t.Errorf("emit error not propagated: %v", gotErr)
	}
}
