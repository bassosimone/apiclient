package apiclient

import "encoding/json"

type stdlibJSONCodec struct{}

func (*stdlibJSONCodec) Encode(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (*stdlibJSONCodec) Decode(b []byte, v interface{}) error {
	return json.Unmarshal(b, v)
}
