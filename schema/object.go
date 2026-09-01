package schema

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Object is a generic Schema.org node. It can represent any Schema.org type
// and property, including vocabulary added after this package was released.
type Object struct {
	Node
	types      []string
	properties map[string]any
	err        error
}

// NewObject creates a generic Schema.org node.
func NewObject(schemaType string, options ...NodeOption) *Object {
	object := &Object{
		Node:       newNode(strings.TrimSpace(schemaType), options...),
		properties: make(map[string]any),
	}
	if object.Type == "" {
		object.err = fmt.Errorf("schema: object type must not be blank")
	}
	return object
}

// AddType appends an additional @type. Duplicate and blank types are ignored.
func (object *Object) AddType(schemaType string) *Object {
	if object == nil {
		return object
	}
	schemaType = strings.TrimSpace(schemaType)
	if schemaType == "" {
		if object.err == nil {
			object.err = fmt.Errorf("schema: object type must not be blank")
		}
		return object
	}
	if schemaType == object.Type {
		return object
	}
	for _, existing := range object.types {
		if existing == schemaType {
			return object
		}
	}
	object.types = append(object.types, schemaType)
	return object
}

// Set assigns a Schema.org property.
func (object *Object) Set(property string, value any) *Object {
	if object == nil {
		return object
	}
	property = strings.TrimSpace(property)
	if property == "" {
		if object.err == nil {
			object.err = fmt.Errorf("schema: property name must not be blank")
		}
		return object
	}
	if strings.HasPrefix(property, "@") {
		if object.err == nil {
			object.err = fmt.Errorf("schema: JSON-LD keyword %q cannot be set as a property", property)
		}
		return object
	}
	object.properties[property] = value
	return object
}

// Add appends a value to an array-valued Schema.org property.
func (object *Object) Add(property string, value any) *Object {
	if object == nil {
		return object
	}
	property = strings.TrimSpace(property)
	if property == "" || strings.HasPrefix(property, "@") {
		return object.Set(property, value)
	}
	current, exists := object.properties[property]
	if !exists {
		object.properties[property] = []any{value}
		return object
	}
	if values, ok := current.([]any); ok {
		object.properties[property] = append(values, value)
	} else {
		object.properties[property] = []any{current, value}
	}
	return object
}

// Err returns the first structural builder error.
func (object *Object) Err() error {
	if object == nil {
		return fmt.Errorf("schema: nil object")
	}
	return object.err
}

// Reference returns an @id-only reference when the object has an identifier.
func (object *Object) Reference() *Reference {
	if object == nil {
		return nil
	}
	return object.Node.Reference()
}

// MarshalJSON encodes the generic object while keeping JSON-LD keywords
// separate from ordinary Schema.org properties.
func (object Object) MarshalJSON() ([]byte, error) {
	if object.err != nil {
		return nil, object.err
	}
	if strings.TrimSpace(object.Type) == "" {
		return nil, fmt.Errorf("schema: object type must not be blank")
	}
	document := make(map[string]any, len(object.properties)+2)
	if object.ID != "" {
		document["@id"] = object.ID
	}
	types := append([]string{object.Type}, object.types...)
	if len(types) == 1 {
		document["@type"] = types[0]
	} else {
		document["@type"] = types
	}
	for property, value := range object.properties {
		document[property] = value
	}
	return json.Marshal(document)
}
