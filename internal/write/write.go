package write

import (
	"bufio"
	"io"

	"url-query/internal/encode"
	"url-query/internal/parse"
)

// Query writes the encoded form and flushes.
func Query(w io.Writer, q *parse.Query) (err error) {
	bw := bufio.NewWriter(w)
	defer func() {
		if ferr := bw.Flush(); ferr != nil && err == nil {
			err = ferr
		}
	}()
	_, err = bw.WriteString(encode.Encode(q))
	return err
}
