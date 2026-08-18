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

func NewArticle(
	headline string,
	options ...NodeOption,
) Article {
	return Article{
		Node: newNode(
			"Article",
			options...,
		),
		Headline: headline,
	}
}

func NewPersonAuthor(
	name string,
	options ...NodeOption,
) Author {
	return Author{
		Node: newNode(
			"Person",
			options...,
		),
		Name: name,
	}
}

func NewOrganisationAuthor(
	name string,
	options ...NodeOption,
) Author {
	return Author{
		Node: newNode(
			"Organization",
			options...,
		),
		Name: name,
	}
}

func NewPublisher(
	name string,
	options ...NodeOption,
) Publisher {
	return Publisher{
		Node: newNode(
			"Organization",
			options...,
		),
		Name: name,
	}
}

func NewImageObject(
	url string,
	options ...NodeOption,
) ImageObject {
	return ImageObject{
		Node: newNode(
			"ImageObject",
			options...,
		),
		URL: url,
	}
}
