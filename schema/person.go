package schema

type SocialProfile string

type Person struct {
	Node

	Name           string          `json:"name"`
	URL            string          `json:"url,omitempty"`
	PictureURL     string          `json:"image,omitempty"`
	SocialProfiles []SocialProfile `json:"sameAs,omitempty"`
	JobTitle       string          `json:"jobTitle,omitempty"`
	WorksFor       *Organisation   `json:"worksFor,omitempty"`
}

func NewPerson(
	name string,
	options ...NodeOption,
) Person {
	return Person{
		Node: newNode(
			"Person",
			options...,
		),
		Name: name,
	}
}
