package schema_test

import (
	"testing"

	"github.com/danieljmanningdev/go-jsonld-schema/schema"
)

func TestIDHelpers(t *testing.T) {
	tests := []struct {
		name  string
		build func(string) string
		input string
		want  string
	}{
		{
			name:  "website ID without trailing slash",
			build: schema.WebsiteID,
			input: "https://example.com",
			want:  "https://example.com/#website",
		},
		{
			name:  "website ID with trailing slash",
			build: schema.WebsiteID,
			input: "https://example.com/",
			want:  "https://example.com/#website",
		},
		{
			name:  "website ID removes query and fragment",
			build: schema.WebsiteID,
			input: "https://example.com?utm_source=test#old",
			want:  "https://example.com/#website",
		},
		{
			name:  "person ID",
			build: schema.PersonID,
			input: "https://example.com/",
			want:  "https://example.com/#person",
		},
		{
			name:  "page ID without trailing slash",
			build: schema.PageID,
			input: "https://example.com/about",
			want:  "https://example.com/about#webpage",
		},
		{
			name:  "page ID with trailing slash",
			build: schema.PageID,
			input: "https://example.com/about/",
			want:  "https://example.com/about/#webpage",
		},
		{
			name:  "page ID removes query and existing fragment",
			build: schema.PageID,
			input: "https://example.com/about?utm_source=test#section",
			want:  "https://example.com/about#webpage",
		},
		{
			name:  "root page ID preserves root slash",
			build: schema.PageID,
			input: "https://example.com",
			want:  "https://example.com/#webpage",
		},
		{
			name:  "organisation ID",
			build: schema.OrganisationID,
			input: "https://example.com/",
			want:  "https://example.com/#organization",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.build(test.input)

			if got != test.want {
				t.Errorf(
					"expected %q, got %q",
					test.want,
					got,
				)
			}
		})
	}
}
