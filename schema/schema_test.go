package schema_test

import (
	"encoding/json"
	"testing"

	"github.com/danieljmanningdev/go-jsonld-schema/schema"
)

func TestWebsiteMarshal(t *testing.T) {
	site := schema.NewWebsite(
		"Example",
		"https://example.com",
	)

	data, err := schema.Marshal(site)
	if err != nil {
		t.Fatal(err)
	}

	var result map[string]any

	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatal(err)
	}

	if result["@context"] != "https://schema.org" {
		t.Errorf("unexpected @context: %v", result["@context"])
	}

	if result["@type"] != "WebSite" {
		t.Errorf("unexpected @type: %v", result["@type"])
	}

	if result["name"] != "Example" {
		t.Errorf("unexpected name: %v", result["name"])
	}
}
