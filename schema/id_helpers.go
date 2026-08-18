package schema

import (
	"net/url"
	"strings"
)

const (
	websiteIDFragment      = "website"
	personIDFragment       = "person"
	pageIDFragment         = "webpage"
	organisationIDFragment = "organization"
)

func WebsiteID(base string) string {
	return buildID(base, websiteIDFragment)
}

func PersonID(base string) string {
	return buildID(base, personIDFragment)
}

func PageID(pageURL string) string {
	return buildID(pageURL, pageIDFragment)
}

func buildID(rawURL, fragment string) string {
	parsedURL, err := url.Parse(strings.TrimSpace(rawURL))
	if err != nil {
		return ""
	}

	// Tracking parameters and existing fragments must not become
	// part of the stable identity of a JSON-LD node.
	parsedURL.RawQuery = ""
	parsedURL.ForceQuery = false
	parsedURL.Fragment = ""
	parsedURL.RawFragment = ""

	// Preserve existing page paths and trailing slashes, but make
	// sure the site root has its canonical slash.
	if parsedURL.Path == "" {
		parsedURL.Path = "/"
	}

	parsedURL.Fragment = fragment

	return parsedURL.String()
}

func OrganisationID(base string) string {
	return buildID(base, organisationIDFragment)
}
