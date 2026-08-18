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

func NewPerson(name string) Person {
	return Person{
		Node: Node{
			Type: "Person",
		},
		Name: name,
	}
}
