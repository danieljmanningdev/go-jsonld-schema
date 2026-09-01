package schema

// Organisation is the British-English Go name for schema:Organization.
type Organisation struct {
	Thing
	Logo string `json:"logo,omitempty"`
}

// Organization is the Schema.org-spelled alias for Organisation.
type Organization = Organisation

func NewOrganisation(name string, url string, options ...NodeOption) Organisation {
	return Organisation{Thing: Thing{Node: newNode(TypeOrganization, options...), Name: name, URL: url}}
}

func NewOrganization(name string, url string, options ...NodeOption) Organization {
	return NewOrganisation(name, url, options...)
}
