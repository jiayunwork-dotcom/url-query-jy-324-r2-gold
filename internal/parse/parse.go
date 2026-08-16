package parse

import (
	"fmt"
	"strings"
)

// Pair is one decoded key/value.
type Pair struct {
	Key   string
	Value string
}

// Query is a decoded query string. Empty input yields a non-nil empty slice.
type Query struct {
	Pairs []Pair
}

// Parse decodes application/x-www-form-urlencoded text.
// '+' becomes space. Percent sequences must be two hex digits.
func Parse(s string) (*Query, error) {
	s = strings.TrimPrefix(s, "?")
	out := &Query{Pairs: []Pair{}}
	if s == "" {
		return out, nil
	}
	for _, part := range strings.Split(s, "&") {
		if part == "" {
			return nil, fmt.Errorf("parse: empty pair")
		}
		k, v, ok := strings.Cut(part, "=")
		if !ok {
			v = ""
		}
		dk, err := unescape(k)
		if err != nil {
			return nil, err
		}
		dv, err := unescape(v)
		if err != nil {
			return nil, err
		}
		if dk == "" {
			return nil, fmt.Errorf("parse: empty key")
		}
		out.Pairs = append(out.Pairs, Pair{Key: dk, Value: dv})
	}
	return out, nil
}

func unescape(s string) (string, error) {
	s = strings.ReplaceAll(s, "+", " ")
	var b strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] != '%' {
			b.WriteByte(s[i])
			continue
		}
		if i+2 >= len(s) {
			return "", fmt.Errorf("parse: truncated percent escape")
		}
		hi, ok1 := fromHex(s[i+1])
		lo, ok2 := fromHex(s[i+2])
		if !ok1 || !ok2 {
			return "", fmt.Errorf("parse: bad percent escape")
		}
		b.WriteByte(hi<<4 | lo)
		i += 2
	}
	return b.String(), nil
}

func fromHex(c byte) (byte, bool) {
	switch {
	case c >= '0' && c <= '9':
		return c - '0', true
	case c >= 'a' && c <= 'f':
		return c - 'a' + 10, true
	case c >= 'A' && c <= 'F':
		return c - 'A' + 10, true
	}
	return 0, false
}
