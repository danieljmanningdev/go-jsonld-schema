package schema_test

import (
	"encoding/json"
	"testing"

	"github.com/danieljmanningdev/go-jsonld-schema/schema"
)

func TestWithIDAppliesIdentifierDuringConstruction(t *testing.T) {
	const baseURL = "https://example.com"

	expectedID := schema.WebsiteID(baseURL)

	website := schema.NewWebsite(
		"Example",
		baseURL,
		schema.WithID(expectedID),
	)

	if website.ID != expectedID {
		t.Fatalf(
			"expected node ID %q, got %q",
			expectedID,
			website.ID,
		)
	}

	data, err := schema.Marshal(website)
	if err != nil {
		t.Fatalf("marshal website: %v", err)
	}

	var result struct {
		ID string `json:"@id"`
	}

	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("unmarshal website: %v", err)
	}

	if result.ID != expectedID {
		t.Errorf(
			"expected marshaled @id %q, got %q",
			expectedID,
			result.ID,
		)
	}
}
