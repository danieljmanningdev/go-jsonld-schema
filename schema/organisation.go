package schema

type Organisation struct {
	Thing

	Logo string `json:"logo,omitempty"`
}

func NewOrganisation(
	name string,
	url string,
	options ...NodeOption,
) Organisation {
	return Organisation{
		Thing: Thing{
			Node: newNode(
				TypeOrganization,
				options...,
			),
			Name: name,
			URL:  url,
		},
	}
}
