package tozstd

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/klauspost/compress/zstd"

	"github.com/klauspost/compress/snappy"
)

func TestSnappy_Convert(t *testing.T) {
	in, err := os.Open("testdata/xml")
	if err != nil {
		t.Fatal(err)
	}

	var comp bytes.Buffer
	w := snappy.NewBufferedWriter(&comp)
	_, err = io.Copy(w, in)
	if err != nil {
		t.Fatal(err)
	}
	err = w.Close()
	if err != nil {
		t.Fatal(err)
	}
	s := Snappy{}
	var dst bytes.Buffer
	n, err := s.Convert(&comp, &dst)
	if err != io.EOF {
		t.Fatal(err)
	}
	t.Log(n, dst.Len())

	dec, err := zstd.NewReader(nil)
	if err != nil {
		t.Fatal(err)
	}
	decoded, err := dec.DecodeAll(dst.Bytes(), nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(len(decoded))

}
