package schema

type BreadcrumbList struct {
	Root

	ItemListElement []ListItem `json:"itemListElement"`
}

type ListItem struct {
	Typed

	Position int    `json:"position"`
	Name     string `json:"name"`
	Item     string `json:"item"`
}

func NewBreadcrumbList(items []ListItem) BreadcrumbList {
	return BreadcrumbList{
		Root: Root{
			Context: "https://schema.org/",
			Type:    "BreadcrumbList",
		},
		ItemListElement: items,
	}
}
