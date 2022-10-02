package z85m

import (
	"errors"
	"reflect"

	"github.com/tilinna/z85"
)

var (
	paddings = [][]byte{
		{1},
		{2, 2},
		{3, 3, 3},
		{},
	}

	InvalidExtension = errors.New("invalid extension")
	InvalidLength    = errors.New("invalid encoded length")
	InvalidPadding   = errors.New("invalid padding")
)

func Encode(src []byte) ([]byte, error) {
	src, pad := pad(src)
	dst := make([]byte, z85.EncodedLen(len(src)))
	_, err := z85.Encode(dst, src)
	if err != nil {
		return nil, err
	}

	if pad != 4 {
		dst = append(dst, byte(pad+'0'))
	}
	return dst, nil
}

func Decode(src []byte) ([]byte, error) {
	src, pad, err := unpad(src)
	if err != nil {
		return nil, err
	}
	dst := make([]byte, z85.DecodedLen(len(src)))
	_, err = z85.Decode(dst, src)
	if err != nil {
		return nil, err
	}

	check := dst[len(dst)-pad:]
	expect := paddings[(pad+3)%4]
	if !reflect.DeepEqual(check, expect) {
		return nil, InvalidPadding
	}

	return dst[:len(dst)-pad], nil
}

func pad(data []byte) ([]byte, int) {
	which := len(data) % 4
	return append(data, paddings[3-which]...), 4 - which
}

func unpad(data []byte) ([]byte, int, error) {
	switch len(data) % 5 {
	case 0:
		return data, 0, nil
	case 1:
		pad := data[len(data)-1] - '0'
		if pad < 1 || pad > 3 {
			return nil, 0, InvalidExtension
		}
		return data[:len(data)-1], int(pad), nil
	default:
		return nil, 0, InvalidLength
	}
}
