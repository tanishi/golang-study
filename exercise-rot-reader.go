package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(c, base byte) byte {
	return (c-base+13)%26 + base
}

func (r *rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)

	if err != nil {
		return 0, err
	}

	for i := 0; i < n; i++ {
		if 'A' <= b[i] && b[i] <= 'Z' {
			b[i] = rot13(b[i], 'A')
		} else if 'a' <= b[i] && b[i] <= 'z' {
			b[i] = rot13(b[i], 'a')
		}
	}

	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
