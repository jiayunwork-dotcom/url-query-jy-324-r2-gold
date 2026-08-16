// Command url-query parses or encodes application/x-www-form-urlencoded text.
package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"url-query/internal/encode"
	"url-query/internal/parse"
	"url-query/internal/write"
)

func main() {
	mode := flag.String("mode", "parse", "parse or encode")
	in := flag.String("in", "-", "input path, or - for stdin")
	out := flag.String("out", "-", "output path, or - for stdout")
	flag.Parse()

	raw, err := readAll(*in)
	if err != nil {
		fatal("read: %v", err)
	}
	text := string(raw)

	w := os.Stdout
	if *out != "-" {
		f, err := os.Create(*out)
		if err != nil {
			fatal("create %s: %v", *out, err)
		}
		defer f.Close()
		w = f
	}

	switch *mode {
	case "parse":
		q, err := parse.Parse(text)
		if err != nil {
			fatal("%v", err)
		}
		for _, p := range q.Pairs {
			fmt.Fprintf(w, "%s\t%s\n", p.Key, p.Value)
		}
	case "encode":
		q, err := parse.Parse(text)
		if err != nil {
			fatal("%v", err)
		}
		if err := write.Query(w, q); err != nil {
			fatal("write: %v", err)
		}
		_ = encode.Encode(q)
	default:
		fatal("unknown -mode %q", *mode)
	}
}

func readAll(path string) ([]byte, error) {
	if path == "-" {
		return io.ReadAll(os.Stdin)
	}
	return os.ReadFile(path)
}

func fatal(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, "url-query: "+format+"\n", a...)
	os.Exit(1)
}
