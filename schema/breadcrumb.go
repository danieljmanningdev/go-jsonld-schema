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
			Type: "BreadcrumbList",
		},
		ItemListElement: items,
	}
}

func NewListItem(position int, name, item string) ListItem {
	return ListItem{
		Typed: Typed{
			Type: "ListItem",
		},
		Position: position,
		Name:     name,
		Item:     item,
	}
}
