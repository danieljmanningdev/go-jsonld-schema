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

func NewBreadcrumbList(
	items []ListItem,
	options ...NodeOption,
) BreadcrumbList {
	return BreadcrumbList{
		Node: newNode(
			"BreadcrumbList",
			options...,
		),
		ItemListElement: items,
	}
}

func NewListItem(
	position int,
	name string,
	item string,
	options ...NodeOption,
) ListItem {
	return ListItem{
		Node: newNode(
			"ListItem",
			options...,
		),
		Position: position,
		Name:     name,
		Item:     item,
	}
}
