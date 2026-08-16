package parse

import "testing"

func TestParseBasic(t *testing.T) {
	q, err := Parse("name=Alice&city=Beijing")
	if err != nil {
		t.Fatal(err)
	}
	if len(q.Pairs) != 2 || q.Pairs[0].Key != "name" || q.Pairs[1].Value != "Beijing" {
		t.Fatalf("pairs=%v", q.Pairs)
	}
}

func TestParsePlusIsSpace(t *testing.T) {
	q, err := Parse("q=hello+world")
	if err != nil {
		t.Fatal(err)
	}
	if q.Pairs[0].Value != "hello world" {
		t.Fatalf("got %q", q.Pairs[0].Value)
	}
}

func TestParseBadPercent(t *testing.T) {
	if _, err := Parse("q=%ZZ"); err == nil {
		t.Fatal("expected error for bad percent escape")
	}
}

func TestParseEmptyNonNil(t *testing.T) {
	q, err := Parse("")
	if err != nil {
		t.Fatal(err)
	}
	if q == nil || q.Pairs == nil {
		t.Fatal("empty input must return non-nil Query with empty Pairs")
	}
}

func TestParseLeadingQuestion(t *testing.T) {
	q, err := Parse("?a=1")
	if err != nil {
		t.Fatal(err)
	}
	if len(q.Pairs) != 1 || q.Pairs[0].Key != "a" {
		t.Fatalf("pairs=%v", q.Pairs)
	}
}
