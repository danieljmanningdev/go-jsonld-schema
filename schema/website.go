package schema

type WebSite struct {
	Root

	Name string `json:"name"`
	URL  string `json:"url"`

	PotentialAction *SearchAction `json:"potentialAction,omitempty"`
}

type SearchAction struct {
	Typed

	Target     string `json:"target"`
	QueryInput string `json:"query-input"`
}

func NewWebsite(name, url string) WebSite {
	return WebSite{
		Root: Root{
			Type: "WebSite",
		},
		Name: name,
		URL:  url,
	}
}
