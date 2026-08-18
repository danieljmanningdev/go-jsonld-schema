package schema

type Article struct {
	Node

	Headline      string     `json:"headline"`
	Image         string     `json:"image,omitempty"`
	Author        *Author    `json:"author,omitempty"`
	Publisher     *Publisher `json:"publisher,omitempty"`
	DatePublished string     `json:"datePublished,omitempty"`
}

type Author struct {
	Node

	Name string `json:"name"`
}

type Publisher struct {
	Node

	Name string       `json:"name"`
	Logo *ImageObject `json:"logo,omitempty"`
}

type ImageObject struct {
	Node

	URL string `json:"url"`
}

func NewArticle(headline string) Article {
	return Article{
		Node: Node{
			Type: "Article",
		},
		Headline: headline,
	}
}

func NewPersonAuthor(name string) Author {
	return Author{
		Node: Node{
			Type: "Person",
		},
		Name: name,
	}
}

func NewOrganisationAuthor(name string) Author {
	return Author{
		Node: Node{
			Type: "Organization",
		},
		Name: name,
	}
}

func NewPublisher(name string) Publisher {
	return Publisher{
		Node: Node{
			Type: "Organization",
		},
		Name: name,
	}
}

func NewImageObject(url string) ImageObject {
	return ImageObject{
		Node: Node{
			Type: "ImageObject",
		},
		URL: url,
	}
}
