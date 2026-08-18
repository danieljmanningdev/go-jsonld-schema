package schema

type WebSite struct {
	Node

	Name string `json:"name"`
	URL  string `json:"url"`

	PotentialAction *SearchAction `json:"potentialAction,omitempty"`
}

type SearchAction struct {
	Node

	Target     string `json:"target"`
	QueryInput string `json:"query-input"`
}

func NewWebsite(
	name string,
	url string,
	options ...NodeOption,
) WebSite {
	return WebSite{
		Node: newNode(
			"WebSite",
			options...,
		),
		Name: name,
		URL:  url,
	}
}
