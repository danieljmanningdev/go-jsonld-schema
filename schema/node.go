package schema

const SchemaOrgContext = "https://schema.org"

type Node struct {
	ID   string `json:"@id,omitempty"`
	Type string `json:"@type"`
}
