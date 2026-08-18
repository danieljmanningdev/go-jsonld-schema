package schema

type WebSite struct {
	Context string `json:"@context"`
	Type    string `json:"@type"`
	Name    string `json:"name"`
	URL     string `json:"url"`

	PotentialAction *SearchAction `json:"potentialAction,omitempty"`
}

type SearchAction struct {
	Type       string `json:"@type"`
	Target     string `json:"target"`
	QueryInput string `json:"query-input"`
}

func NewWebsite(name, url string) WebSite {
	return WebSite{
		Context: "https://schema.org",
		Type:    "WebSite",
		Name:    name,
		URL:     url,
	}
}
