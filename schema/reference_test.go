package schema_test

import (
	"encoding/json"
	"testing"

	"github.com/danieljmanningdev/go-jsonld-schema/schema"
)

func TestWebsitePublisherReferencesOrganisationNode(t *testing.T) {
	const baseURL = "https://example.com"

	organisationID := schema.OrganisationID(baseURL)

	organisation := schema.NewOrganisation(
		"Example Ltd",
		baseURL,
		schema.WithID(organisationID),
	)

	website := schema.NewWebsite(
		"Example",
		baseURL,
		schema.WithID(
			schema.WebsiteID(baseURL),
		),
	)

	website.Publisher = organisation.Reference()

	graph := schema.NewGraph(
		website,
		organisation,
	)

	data, err := schema.Marshal(graph)
	if err != nil {
		t.Fatalf("marshal graph: %v", err)
	}

	var result struct {
		Graph []map[string]any `json:"@graph"`
	}

	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("unmarshal graph: %v", err)
	}

	var websiteNode map[string]any
	var organisationNode map[string]any

	for _, node := range result.Graph {
		switch node["@type"] {
		case "WebSite":
			websiteNode = node

		case "Organization":
			organisationNode = node
		}
	}

	if websiteNode == nil {
		t.Fatal("WebSite node was not found in graph")
	}

	if organisationNode == nil {
		t.Fatal("Organization node was not found in graph")
	}

	publisher, ok := websiteNode["publisher"].(map[string]any)
	if !ok {
		t.Fatalf(
			"expected publisher reference object, got %#v",
			websiteNode["publisher"],
		)
	}

	if got := publisher["@id"]; got != organisationID {
		t.Errorf(
			"expected publisher @id %q, got %v",
			organisationID,
			got,
		)
	}

	if got := organisationNode["@id"]; got != organisationID {
		t.Errorf(
			"expected Organization node @id %q, got %v",
			organisationID,
			got,
		)
	}

	if len(publisher) != 1 {
		t.Errorf(
			"expected publisher to contain only @id, got %#v",
			publisher,
		)
	}
}
