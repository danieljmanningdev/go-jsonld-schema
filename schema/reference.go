package schema

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Reference is an @id-only JSON-LD node reference.
type Reference struct {
	ID string `json:"@id"`
}

// Ref creates an @id-only reference.
func Ref(id string) Reference {
	return Reference{ID: strings.TrimSpace(id)}
}

// Validate checks that the reference has an identifier.
func (reference Reference) Validate() error {
	trimmedID := strings.TrimSpace(reference.ID)
	if trimmedID == "" {
		return fmt.Errorf("schema: reference ID must not be empty")
	}
	if trimmedID != reference.ID {
		return fmt.Errorf("schema: reference ID must not contain surrounding whitespace")
	}
	return nil
}

// MarshalJSON validates and serializes the @id-only reference.
func (reference Reference) MarshalJSON() ([]byte, error) {
	if err := reference.Validate(); err != nil {
		return nil, err
	}
	type referenceAlias Reference
	return json.Marshal(referenceAlias(reference))
}

// Reference returns an @id-only reference to node, or nil when node has no ID.
func (node Node) Reference() *Reference {
	reference := Ref(node.ID)
	if reference.ID == "" {
		return nil
	}
	return &reference
}
