package csvutil

import (
	"bytes"
	"reflect"
	"strconv"
	"time"
)

func encodeDurationHack(typ reflect.Type) encodeFunc {
	var b [64]byte

	durationType := reflect.TypeOf(time.Duration(0))
	if typ != durationType {
		return nil
	}

	return func(v reflect.Value, buf *bytes.Buffer, _ bool) (int, error) {
		t := time.Duration(v.Int())
		return buf.Write(strconv.AppendFloat(b[:0], t.Seconds(), 'f', 3, 64))
	}
	return nil
}

func decodeDurationHack(typ reflect.Type) decodeFunc {
	durationType := reflect.TypeOf(time.Duration(0))
	if typ != durationType {
		return nil
	}

	return func(s string, v reflect.Value) error {
		d, err := time.ParseDuration(s + "s")
		if err != nil {
			return &UnmarshalTypeError{Value: s, Type: typ}
		}

		v.SetInt(int64(d))

		return nil
	}
}
