package client

import (
	"bytes"
	"encoding/json"
	"io"
)

func encodeJSON(v any) (io.Reader, error) {
	if v == nil {
		return nil, nil
	}

	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(b), nil
}
