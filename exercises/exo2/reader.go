package main

import (
	"io"
	"os"
	"strings"
)

type spaceEraser struct {
	r io.Reader
}

func (tr spaceEraser) Read(buffer []byte) (int, error) {
	n, err := tr.r.Read(buffer)
	var slice []string
	toread := strings.Split(string(buffer), "")
	for _, el := range toread {
		if el != " " {
			slice = append(slice, el)
		}
	}
	final := []byte(strings.Join(slice, ""))
	copy(buffer, final)
	return n, err
}

func main() {
	s := strings.NewReader("H e l l o w o r l d!")
	r := spaceEraser{s}
	io.Copy(os.Stdout, &r)
}
