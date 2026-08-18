package schema

type SocialProfile string

type Person struct {
	Thing

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
		Thing: Thing{
			Node: newNode(
				TypePerson,
				options...,
			),
			Name: name,
		},
	}
}
