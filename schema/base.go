package schema

type Typed struct {
	ID   string `json:"@id,omitempty"`
	Type string `json:"@type"`
}

type Root struct {
	Context string `json:"@context,omitempty"`
	ID      string `json:"@id,omitempty"`
	Type    string `json:"@type"`
}
