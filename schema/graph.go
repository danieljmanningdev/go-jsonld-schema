package schema

type Graph struct {
	Context string `json:"@context"`
	Graph   []any  `json:"@graph"`
}

func NewGraph(nodes ...any) Graph {
	return Graph{
		Context: "https://schema.org",
		Graph:   nodes,
	}
}
