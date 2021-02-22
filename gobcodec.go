package apiclient

import (
	"bytes"
	"encoding/gob"
)

type stdlibGobCodec struct{}

func (*stdlibGobCodec) Encode(v interface{}) ([]byte, error) {
	var bb bytes.Buffer
	if err := gob.NewEncoder(&bb).Encode(v); err != nil {
		return nil, err
	}
	return bb.Bytes(), nil
}

func (*stdlibGobCodec) Decode(b []byte, v interface{}) error {
	return gob.NewDecoder(bytes.NewReader(b)).Decode(v)
}
