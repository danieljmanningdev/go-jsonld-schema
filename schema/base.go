package schema

type Typed struct {
	Type string `json:"@type"`
}

type Root struct {
	Context string `json:"@context"`
	Type    string `json:"@type"`
}
