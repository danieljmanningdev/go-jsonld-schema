package schema

import "strings"

type Reference struct {
	ID string `json:"@id"`
}

func Ref(id string) Reference {
	return Reference{
		ID: strings.TrimSpace(id),
	}
}

func (node Node) Reference() *Reference {
	reference := Ref(node.ID)

	if reference.ID == "" {
		return nil
	}

	return &reference
}
