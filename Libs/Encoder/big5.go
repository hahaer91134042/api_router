package Encoder

import (
	"bytes"
	"io"

	tc "golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

var Big5 = &Big5Encoder{}

type Big5Encoder struct {
}

/*
convert UTF-8 to BIG5
*/
func (encoder *Big5Encoder) Encode(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, tc.Big5.NewEncoder())
	d, e := io.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}

/*
convert BIG5 to UTF-8
*/
func (encoder *Big5Encoder) Decode(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, tc.Big5.NewDecoder())
	d, e := io.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}
