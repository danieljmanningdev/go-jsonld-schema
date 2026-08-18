package schema

import "encoding/json"

func Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func MarshalIndent(v any) ([]byte, error) {
	return json.MarshalIndent(v, "", "  ")
}
