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

	if result["@type"] != "WebSite" {
		t.Errorf("expected WebSite, got %v", result["@type"])
	}
}

func TestPersonMarshal(t *testing.T) {
	person := schema.NewPerson("Daniel Manning")

	data, err := schema.Marshal(person)
	if err != nil {
		t.Fatal(err)
	}

	var result map[string]any

	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatal(err)
	}

	if result["@type"] != "Person" {
		t.Errorf("expected Person, got %v", result["@type"])
	}

	if result["name"] != "Daniel Manning" {
		t.Errorf("unexpected name: %v", result["name"])
	}
}

func TestOrganisationMarshal(t *testing.T) {
	org := schema.NewOrganisation(
		"Example Ltd",
		"https://example.com",
	)

	data, err := schema.Marshal(org)
	if err != nil {
		t.Fatal(err)
	}

	var result map[string]any

	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatal(err)
	}

	if result["@type"] != "Organization" {
		t.Errorf("expected Organization, got %v", result["@type"])
	}
}

func TestBreadcrumbListMarshal(t *testing.T) {
	breadcrumbs := schema.NewBreadcrumbList([]schema.ListItem{
		schema.NewListItem(
			1,
			"Home",
			"https://example.com/",
		),
		schema.NewListItem(
			2,
			"Blog",
			"https://example.com/blog",
		),
	})

	data, err := schema.Marshal(breadcrumbs)
	if err != nil {
		t.Fatal(err)
	}

	var result map[string]any

	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatal(err)
	}

	if result["@type"] != "BreadcrumbList" {
		t.Errorf("expected BreadcrumbList, got %v", result["@type"])
	}
}

func TestArticleMarshal(t *testing.T) {
	article := schema.NewArticle("Example Article")

	author := schema.NewPersonAuthor("Daniel Manning")
	publisher := schema.NewPublisher("Example Ltd")
	logo := schema.NewImageObject("https://example.com/logo.png")

	publisher.Logo = &logo
	article.Author = &author
	article.Publisher = &publisher

	data, err := schema.Marshal(article)
	if err != nil {
		t.Fatal(err)
	}

	var result map[string]any

	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatal(err)
	}

	if result["@type"] != "Article" {
		t.Errorf("expected Article, got %v", result["@type"])
	}
}
