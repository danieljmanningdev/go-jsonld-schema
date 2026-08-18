package schema

type Organisation struct {
	Node

	Name string `json:"name"`
	URL  string `json:"url,omitempty"`
	Logo string `json:"logo,omitempty"`
}

func NewOrganisation(name, url string) Organisation {
	return Organisation{
		Node: Node{
			Type: "Organization",
		},
		Name: name,
		URL:  url,
	}
}
