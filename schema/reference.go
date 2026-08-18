package schema

type Reference struct {
	ID string `json:"@id"`
}

func Ref(id string) Reference {
	return Reference{ID: id}
}
