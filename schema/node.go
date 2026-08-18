package schema

const SchemaOrgContext = "https://schema.org"

type Node struct {
	ID   string `json:"@id,omitempty"`
	Type string `json:"@type"`
}

type NodeOption interface {
	apply(*Node)
}

type nodeOptionFunc func(*Node)

func (option nodeOptionFunc) apply(node *Node) {
	option(node)
}

func WithID(id string) NodeOption {
	return nodeOptionFunc(func(node *Node) {
		node.ID = id
	})
}

func newNode(schemaType string, options ...NodeOption) Node {
	node := Node{
		Type: schemaType,
	}

	for _, option := range options {
		if option != nil {
			option.apply(&node)
		}
	}

	return node
}
