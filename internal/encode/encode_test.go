package encode

import (
	"testing"

	"url-query/internal/parse"
)

func TestEncodeRoundTrip(t *testing.T) {
	q := &parse.Query{Pairs: []parse.Pair{{Key: "name", Value: "Alice"}}}
	s := Encode(q)
	if s != "name=Alice" {
		t.Fatalf("got %q", s)
	}
}

func TestEncodeSpaceAsPlus(t *testing.T) {
	q := &parse.Query{Pairs: []parse.Pair{{Key: "q", Value: "hello world"}}}
	s := Encode(q)
	if s != "q=hello+world" {
		t.Fatalf("got %q", s)
	}
}

func TestEncodeNilQuery(t *testing.T) {
	if Encode(nil) != "" {
		t.Fatal("nil query must encode to empty string")
	}
}

func TestEncodeDoesNotMutateInput(t *testing.T) {
	q := &parse.Query{Pairs: []parse.Pair{{Key: "a", Value: "1"}}}
	_ = Encode(q)
	if q.Pairs[0].Key != "a" || q.Pairs[0].Value != "1" {
		t.Fatalf("encode mutated input: %+v", q.Pairs)
	}
}
