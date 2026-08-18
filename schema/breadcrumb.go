package schema

type BreadcrumbList struct {
	Node

	ItemListElement []ListItem `json:"itemListElement"`
}

type ListItem struct {
	Node

	Position int    `json:"position"`
	Name     string `json:"name"`
	Item     string `json:"item"`
}

func NewBreadcrumbList(items []ListItem) BreadcrumbList {
	return BreadcrumbList{
		Node: Node{
			Type: "BreadcrumbList",
		},
		ItemListElement: items,
	}
}

func NewListItem(position int, name, item string) ListItem {
	return ListItem{
		Node: Node{
			Type: "ListItem",
		},
		Position: position,
		Name:     name,
		Item:     item,
	}
}
