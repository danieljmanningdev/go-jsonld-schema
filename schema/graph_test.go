package schema_test

import (
	"encoding/json"
	"testing"

	"github.com/danieljmanningdev/go-jsonld-schema/schema"
)

func TestGraphMarshalUsesSingleContext(t *testing.T) {
	website := schema.NewWebsite(
		"Example",
		"https://example.com",
	)
	website.ID = schema.WebsiteID(website.URL)

	person := schema.NewPerson("Daniel Manning")
	person.ID = schema.PersonID(website.URL)

	graph := schema.NewGraph(
		website,
		person,
	)

	data, err := schema.Marshal(graph)
	if err != nil {
		t.Fatalf("marshal graph: %v", err)
	}

	var result struct {
		Context string           `json:"@context"`
		Graph   []map[string]any `json:"@graph"`
	}

	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("unmarshal graph: %v", err)
	}

	if result.Context != "https://schema.org" {
		t.Errorf(
			"expected graph context %q, got %q",
			"https://schema.org",
			result.Context,
		)
	}

	if len(result.Graph) != 2 {
		t.Fatalf(
			"expected 2 graph nodes, got %d",
			len(result.Graph),
		)
	}

	for index, node := range result.Graph {
		if _, exists := node["@context"]; exists {
			t.Errorf(
				"graph node %d should not contain its own @context",
				index,
			)
		}
	}
}
