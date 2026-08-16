package write

import (
	"bytes"
	"errors"
	"testing"

	"url-query/internal/parse"
)

type failWriter struct{}

func (failWriter) Write(p []byte) (int, error) {
	return 0, errors.New("write fail")
}

func TestQueryOK(t *testing.T) {
	q := &parse.Query{Pairs: []parse.Pair{{Key: "a", Value: "1"}}}
	var buf bytes.Buffer
	if err := Query(&buf, q); err != nil {
		t.Fatal(err)
	}
	if buf.String() != "a=1" {
		t.Fatalf("got %q", buf.String())
	}
}

func TestQueryFlushError(t *testing.T) {
	q := &parse.Query{Pairs: []parse.Pair{{Key: "a", Value: "1"}}}
	if err := Query(failWriter{}, q); err == nil {
		t.Fatal("expected flush/write error")
	}
}
