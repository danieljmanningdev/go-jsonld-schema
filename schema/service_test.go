package schema

import (
	"encoding/json"
	"testing"
)

func TestNewService(t *testing.T) {
	service := NewService("Web Development", WithID("https://example.com/#service"))
	service.URL = "https://example.com/web-development/"
	service.ServiceType = "Web development"
	service.AreaServed = []string{"United Kingdom"}

	data, err := Marshal(service)
	if err != nil {
		t.Fatalf("Marshal() error = %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}

	if got["@context"] != SchemaOrgContext {
		t.Fatalf("@context = %v, want %q", got["@context"], SchemaOrgContext)
	}
	if got["@type"] != TypeService {
		t.Fatalf("@type = %v, want %q", got["@type"], TypeService)
	}
	if got["name"] != "Web Development" {
		t.Fatalf("name = %v, want Web Development", got["name"])
	}
}

func TestNewProfessionalService(t *testing.T) {
	service := NewProfessionalService("Daniel J. Manning")
	service.URL = "https://example.com"
	service.AreaServed = []string{"Leeds", "United Kingdom"}
	service.PriceRange = "££"

	data, err := Marshal(service)
	if err != nil {
		t.Fatalf("Marshal() error = %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}

	if got["@type"] != TypeProfessionalService {
		t.Fatalf("@type = %v, want %q", got["@type"], TypeProfessionalService)
	}
}
