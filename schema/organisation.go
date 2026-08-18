package schema

type Organisation struct {
	Node

	Name string `json:"name"`
	URL  string `json:"url,omitempty"`
	Logo string `json:"logo,omitempty"`
}

func NewOrganisation(
	name string,
	url string,
	options ...NodeOption,
) Organisation {
	return Organisation{
		Node: newNode(
			"Organization",
			options...,
		),
		Name: name,
		URL:  url,
	}
}
