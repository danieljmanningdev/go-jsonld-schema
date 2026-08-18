package schema

import "strings"

func WebsiteID(base string) string {
	return strings.TrimRight(base, "/") + "/#website"
}

func PersonID(base string) string {
	return strings.TrimRight(base, "/") + "/#person"
}

func PageID(url string) string {
	return strings.TrimRight(url, "/") + "#webpage"
}
