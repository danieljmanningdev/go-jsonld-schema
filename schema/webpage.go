package schema

type WebPage struct {
	Node

	Name string `json:"name"`
	URL  string `json:"url"`

	IsPartOf   *Reference `json:"isPartOf,omitempty"`
	MainEntity *Reference `json:"mainEntity,omitempty"`
}

func NewWebPage(
	name string,
	url string,
	options ...NodeOption,
) WebPage {
	return WebPage{
		Node: newNode(
			"WebPage",
			options...,
		),
		Name: name,
		URL:  url,
	}
}
