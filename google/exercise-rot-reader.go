package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader *rot13Reader) Read(buf []byte) (int, error) {
	n, err := reader.r.Read(buf)
	if err != nil {
		return 0, err
	}

	for i := 0; i < n; i++ {
		buf[i] = reader.rot13(buf[i])
	}
	return n, nil
}

func (reader *rot13Reader) rot13(b byte) byte {
	if ('a' <= b && b <= 'm') || ('A' <= b && b <= 'M') {
		return b + 13
	}
	if ('n' <= b && b <= 'z') || ('N' <= b && b <= 'Z') {
		return b - 13
	}
	return b
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
