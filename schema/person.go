package schema

type SocialProfile string

type Person struct {
	Root

	Name           string          `json:"name"`
	URL            string          `json:"url,omitempty"`
	PictureURL     string          `json:"image,omitempty"`
	SocialProfiles []SocialProfile `json:"sameAs,omitempty"`
	JobTitle       string          `json:"jobTitle,omitempty"`
	WorksFor       *Organisation   `json:"worksFor,omitempty"`
}

func NewPerson(name string) Person {
	return Person{
		Root: Root{
			Context: "https://schema.org/",
			Type:    "Person",
		},
		Name: name,
	}
}
