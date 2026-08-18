package schema

import (
	"encoding/json"
	"fmt"
)

func Marshal(v any) ([]byte, error) {
	return marshalDocument(v, false)
}

func MarshalIndent(v any) ([]byte, error) {
	return marshalDocument(v, true)
}

func marshalDocument(v any, indent bool) ([]byte, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("schema: marshal value: %w", err)
	}

	var document map[string]json.RawMessage

	if err := json.Unmarshal(data, &document); err != nil {
		return nil, fmt.Errorf(
			"schema: value must marshal to a JSON object: %w",
			err,
		)
	}

	if document == nil {
		return nil, fmt.Errorf(
			"schema: value must marshal to a JSON object",
		)
	}

	if _, exists := document["@context"]; !exists {
		context, err := json.Marshal(SchemaOrgContext)
		if err != nil {
			return nil, fmt.Errorf(
				"schema: marshal context: %w",
				err,
			)
		}

		document["@context"] = context
	}

	if indent {
		return json.MarshalIndent(document, "", "  ")
	}

	return json.Marshal(document)
}
