package Encoder

import (
	"bytes"
	"io"

	sc "golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

var GBK = &GBKEncoder{}

type GBKEncoder struct {
}

/*
convert UTF-8 to GBK
*/
func (encoder *GBKEncoder) Encode(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, sc.GBK.NewEncoder())
	d, e := io.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}

/*
convert GBK to UTF-8
*/
func (encoder *GBKEncoder) Decode(s []byte) ([]byte, error) {
	// decoder := mahonia.NewDecoder("GBK")
	// encod := mahonia.NewEncoder("GBK")
	// b, _ := encoder.Encode([]byte("筒条万"))
	// I := bytes.NewReader(b)
	I := bytes.NewReader(s)
	O := transform.NewReader(I, sc.GBK.NewDecoder())

	d, e := io.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
	//return []byte(decoder.ConvertString(encod.ConvertString("筒条万"))), nil
	//return []byte(decoder.ConvertString(string(s))), nil
}
