package schema

type Article struct {
	Root

	Headline      string     `json:"headline"`
	Image         string     `json:"image,omitempty"`
	Author        *Author    `json:"author,omitempty"`
	Publisher     *Publisher `json:"publisher,omitempty"`
	DatePublished string     `json:"datePublished,omitempty"`
}

type Author struct {
	Typed

	Name string `json:"name"`
}

type Publisher struct {
	Typed

	Name string       `json:"name"`
	Logo *ImageObject `json:"logo,omitempty"`
}

type ImageObject struct {
	Typed

	URL string `json:"url"`
}

func NewArticle(headline string) Article {
	return Article{
		Root: Root{
			Context: "https://schema.org",
			Type:    "Article",
		},
		Headline: headline,
	}
}
