package schema

const SchemaOrgContext = "https://schema.org"

type Typed struct {
	ID   string `json:"@id,omitempty"`
	Type string `json:"@type"`
}

type Root struct {
	ID   string `json:"@id,omitempty"`
	Type string `json:"@type"`
}
