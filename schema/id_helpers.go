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

func WebsiteID(base string) string { return FragmentID(base, websiteIDFragment) }
func WebSiteID(base string) string { return WebsiteID(base) }
func PersonID(base string) string { return FragmentID(base, personIDFragment) }
func PageID(pageURL string) string { return FragmentID(pageURL, pageIDFragment) }
func WebPageID(pageURL string) string { return PageID(pageURL) }
func OrganisationID(base string) string { return FragmentID(base, organisationIDFragment) }
func OrganizationID(base string) string { return OrganisationID(base) }

// FragmentID builds a stable node identifier after removing query/fragment.
func FragmentID(rawURL, fragment string) string {
	rawURL = strings.TrimSpace(rawURL)
	fragment = strings.TrimSpace(strings.TrimPrefix(fragment, "#"))
	if rawURL == "" || fragment == "" {
		return ""
	}
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return ""
	}
	parsedURL.RawQuery = ""
	parsedURL.ForceQuery = false
	parsedURL.Fragment = ""
	parsedURL.RawFragment = ""
	if parsedURL.Path == "" {
		parsedURL.Path = "/"
	}
	parsedURL.Fragment = fragment
	return parsedURL.String()
}
