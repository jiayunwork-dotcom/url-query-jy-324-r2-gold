package encode

import (
	"strings"

	"url-query/internal/parse"
)

// Encode writes pairs as application/x-www-form-urlencoded.
// Space becomes '+'. Bytes outside unreserved set are percent-encoded.
func Encode(q *parse.Query) string {
	if q == nil {
		return ""
	}
	parts := make([]string, 0, len(q.Pairs))
	for _, p := range q.Pairs {
		parts = append(parts, escape(p.Key)+"="+escape(p.Value))
	}
	return strings.Join(parts, "&")
}

func escape(s string) string {
	var b strings.Builder
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == ' ' {
			b.WriteString("%20")
			continue
		}
		if unreserved(c) {
			b.WriteByte(c)
			continue
		}
		b.WriteByte('%')
		b.WriteByte(toHex(c >> 4))
		b.WriteByte(toHex(c & 0x0f))
	}
	return b.String()
}

func unreserved(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9') || c == '-' || c == '_' || c == '.' || c == '~'
}

func toHex(n byte) byte {
	if n < 10 {
		return '0' + n
	}
	return 'A' + (n - 10)
}
