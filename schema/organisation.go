package schema

type Organisation struct {
	Root

	Name string `json:"name"`
	URL  string `json:"url,omitempty"`
	Logo string `json:"logo,omitempty"`
}

func NewOrganisation(name, url string) Organisation {
	return Organisation{
		Root: Root{
			Type: "Organization",
		},
		Name: name,
		URL:  url,
	}
}
