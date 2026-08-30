package schema

type WebSite struct {
	Node
	Name            string        `json:"name"`
	URL             string        `json:"url"`
	Publisher       *Reference    `json:"publisher,omitempty"`
	PotentialAction *SearchAction `json:"potentialAction,omitempty"`
}

type SearchAction struct {
	Node
	Target     string `json:"target"`
	QueryInput string `json:"query-input"`
}

func NewWebSite(name string, url string, options ...NodeOption) WebSite {
	return WebSite{Node: newNode(TypeWebSite, options...), Name: name, URL: url}
}

func NewWebsite(name string, url string, options ...NodeOption) WebSite {
	return NewWebSite(name, url, options...)
}

func NewSearchAction(target string, queryInput string, options ...NodeOption) SearchAction {
	return SearchAction{Node: newNode(TypeSearchAction, options...), Target: target, QueryInput: queryInput}
}
