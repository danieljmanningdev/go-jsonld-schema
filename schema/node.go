package schema

import (
	"fmt"
	"strings"
)

// SchemaOrgContext is the canonical Schema.org JSON-LD context URL.
const SchemaOrgContext = "https://schema.org"

// Node contains the JSON-LD identity and primary type shared by typed nodes.
type Node struct {
	ID   string `json:"@id,omitempty"`
	Type string `json:"@type"`
}

// NodeOption customizes a node during construction.
type NodeOption interface {
	apply(*Node)
}

type nodeOptionFunc func(*Node)

func (option nodeOptionFunc) apply(node *Node) {
	option(node)
}

// WithID assigns an @id to a node during construction.
func WithID(id string) NodeOption {
	return nodeOptionFunc(func(node *Node) {
		node.ID = strings.TrimSpace(id)
	})
}

func newNode(schemaType string, options ...NodeOption) Node {
	node := Node{Type: strings.TrimSpace(schemaType)}
	for _, option := range options {
		if option != nil {
			option.apply(&node)
		}
	}
	return node
}

// Validate checks the node identity and primary type.
func (node Node) Validate() error {
	trimmedType := strings.TrimSpace(node.Type)
	if trimmedType == "" {
		return fmt.Errorf("schema: node type must not be empty")
	}
	if trimmedType != node.Type {
		return fmt.Errorf("schema: node type must not contain surrounding whitespace")
	}
	if strings.TrimSpace(node.ID) != node.ID {
		return fmt.Errorf("schema: node ID must not contain surrounding whitespace")
	}
	return nil
}
