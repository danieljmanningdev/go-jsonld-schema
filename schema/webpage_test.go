package schema_test

import (
	"encoding/json"
	"testing"

	"github.com/danieljmanningdev/go-jsonld-schema/schema"
)

func TestWebPageMarshal(t *testing.T) {
	const pageURL = "https://example.com/about"

	expectedID := schema.PageID(pageURL)

	page := schema.NewWebPage(
		"About",
		pageURL,
		schema.WithID(expectedID),
	)

	data, err := schema.Marshal(page)
	if err != nil {
		t.Fatalf("marshal WebPage: %v", err)
	}

	var result struct {
		Context string `json:"@context"`
		ID      string `json:"@id"`
		Type    string `json:"@type"`
		Name    string `json:"name"`
		URL     string `json:"url"`
	}

	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("unmarshal WebPage: %v", err)
	}

	if result.Context != schema.SchemaOrgContext {
		t.Errorf(
			"expected @context %q, got %q",
			schema.SchemaOrgContext,
			result.Context,
		)
	}

	if result.ID != expectedID {
		t.Errorf(
			"expected @id %q, got %q",
			expectedID,
			result.ID,
		)
	}

	if result.Type != "WebPage" {
		t.Errorf(
			"expected @type %q, got %q",
			"WebPage",
			result.Type,
		)
	}

	if result.Name != "About" {
		t.Errorf(
			"expected name %q, got %q",
			"About",
			result.Name,
		)
	}

	if result.URL != pageURL {
		t.Errorf(
			"expected URL %q, got %q",
			pageURL,
			result.URL,
		)
	}
}
