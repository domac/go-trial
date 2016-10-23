package main

import (
	"github.com/gosuri/uiprogress"
	"io"
)

type Progress struct {
	io.Reader
	Total int64
	Recv  int
	Bar   *uiprogress.Bar
}

func (ptr *Progress) Read(p []byte) (int, error) {
	n, err := ptr.Reader.Read(p)
	if n > 0 {
		ptr.Recv += n
		ptr.Bar.Set(ptr.Recv)
	}
	return n, err
}
