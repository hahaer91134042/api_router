package Encoder

type IEncoder interface {
	Encode(s []byte) ([]byte, error)
	Decode(s []byte) ([]byte, error)
}
