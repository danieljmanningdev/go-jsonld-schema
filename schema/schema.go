package schema

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type validator interface {
	Validate() error
}

func Marshal(v any) ([]byte, error) {
	return marshalDocument(v, SchemaOrgContext, false)
}

func MarshalIndent(v any) ([]byte, error) {
	return marshalDocument(v, SchemaOrgContext, true)
}

func MarshalWithContext(v any, context any) ([]byte, error) {
	return marshalDocument(v, context, false)
}

func MarshalIndentWithContext(v any, context any) ([]byte, error) {
	return marshalDocument(v, context, true)
}

func marshalDocument(v any, context any, indent bool) ([]byte, error) {
	if isNilValue(v) {
		return nil, fmt.Errorf("schema: value must marshal to a JSON object")
	}
	if candidate, ok := v.(validator); ok {
		if err := candidate.Validate(); err != nil {
			return nil, fmt.Errorf("schema: validate value: %w", err)
		}
	}
	data, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("schema: marshal value: %w", err)
	}
	var document map[string]json.RawMessage
	if err := json.Unmarshal(data, &document); err != nil {
		return nil, fmt.Errorf("schema: value must marshal to a JSON object: %w", err)
	}
	if document == nil {
		return nil, fmt.Errorf("schema: value must marshal to a JSON object")
	}
	if _, exists := document["@context"]; !exists {
		if isNilValue(context) {
			context = SchemaOrgContext
		}
		encodedContext, err := json.Marshal(context)
		if err != nil {
			return nil, fmt.Errorf("schema: marshal context: %w", err)
		}
		document["@context"] = encodedContext
	}
	if indent {
		return json.MarshalIndent(document, "", "  ")
	}
	return json.Marshal(document)
}

func isNilValue(value any) bool {
	if value == nil {
		return true
	}
	reflected := reflect.ValueOf(value)
	switch reflected.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return reflected.IsNil()
	default:
		return false
	}
}
