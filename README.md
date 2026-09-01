# go-jsonld-schema

[![CI](https://github.com/danieljmanningdev/go-jsonld-schema/actions/workflows/ci.yml/badge.svg)](https://github.com/danieljmanningdev/go-jsonld-schema/actions/workflows/ci.yml)

`go-jsonld-schema` is a lightweight Go package for building and marshaling Schema.org JSON-LD structured data.

It provides typed Go structs and constructors for common structured-data types so JSON-LD can be generated with Go instead of manually writing JSON inside HTML templates.

## Supported Schemas

Currently supported:

- WebSite
- Person
- Organization
- ProfessionalService
- Service
- Offer
- Article
- BreadcrumbList
- SearchAction
- ImageObject
- ListItem
- Graph
- Reference helpers

## Installation

```bash
go get github.com/danieljmanningdev/go-jsonld-schema
```

Import the schema package:

```go
import "github.com/danieljmanningdev/go-jsonld-schema/schema"
```

## WebSite

```go
site := schema.NewWebsite(
    "Example Site",
    "https://example.com",
)

data, err := schema.MarshalIndent(site)
if err != nil {
    log.Fatal(err)
}
```

`schema.Marshal` and `schema.MarshalIndent` automatically add the root JSON-LD context when it is not already present:

```json
{
  "@context": "https://schema.org",
  "@type": "WebSite",
  "name": "Example Site",
  "url": "https://example.com"
}
```

## Person

```go
person := schema.NewPerson("Daniel Manning")
person.URL = "https://example.com"
person.PictureURL = "https://example.com/profile.jpg"
person.JobTitle = "Digital Product Designer & Engineer"
person.SocialProfiles = []schema.SocialProfile{
    "https://github.com/example",
}
```

## Organization

The Go API uses the British spelling `Organisation`, while the generated Schema.org `@type` is `Organization`.

```go
organisation := schema.NewOrganisation(
    "Example Ltd",
    "https://example.com",
)
organisation.Logo = "https://example.com/logo.png"
```

## Service

```go
service := schema.NewService(
    "Web Development",
    schema.WithID("https://example.com/#web-development"),
)
service.URL = "https://example.com/web-development/"
service.Description = "Custom web development services."
service.ServiceType = "Web development"
service.AreaServed = []string{"United Kingdom"}
```

## ProfessionalService

```go
business := schema.NewProfessionalService("Example Studio")
business.URL = "https://example.com"
business.Image = "https://example.com/logo.png"
business.AreaServed = []string{"Leeds", "United Kingdom"}
business.PriceRange = "££"
```

## Offer

```go
offer := schema.NewOffer()
offer.URL = "https://example.com/services/web-development"
offer.Price = "1500"
offer.PriceCurrency = "GBP"
```

Offers can be attached to a `Service` through its `Offers` field.

## BreadcrumbList

```go
breadcrumbs := schema.NewBreadcrumbList([]schema.ListItem{
    schema.NewListItem(1, "Home", "https://example.com/"),
    schema.NewListItem(2, "Blog", "https://example.com/blog"),
    schema.NewListItem(3, "Article", "https://example.com/blog/article"),
})
```

## Article

```go
article := schema.NewArticle("How HTMX Works")
author := schema.NewPersonAuthor("Daniel Manning")
publisher := schema.NewPublisher("Example Ltd")
logo := schema.NewImageObject("https://example.com/logo.png")

publisher.Logo = &logo
article.Author = &author
article.Publisher = &publisher
article.Image = "https://example.com/article.jpg"
article.DatePublished = "2026-08-18"
```

## Graph

Use a graph when one page needs several related JSON-LD nodes:

```go
graph := schema.NewGraph(
    person,
    site,
    breadcrumbs,
)

data, err := schema.Marshal(graph)
```

## IDs and References

Nodes can be given stable `@id` values with `schema.WithID(...)`. Use the reference helpers when one node should point at another existing node rather than embedding the whole object repeatedly.

## Marshaling

Compact JSON:

```go
data, err := schema.Marshal(site)
```

Indented JSON:

```go
data, err := schema.MarshalIndent(site)
```

Both helpers require the root value to marshal to a JSON object and ensure `@context` is present.

## Using with Go HTML Templates

Generate JSON through the package first, then pass the resulting bytes/string to your template. Do not construct JSON-LD by concatenating untrusted strings.

```html
{{with .JSONLD}}
<script type="application/ld+json">{{.}}</script>
{{end}}
```

When using `html/template`, only convert generated JSON to a trusted template type after it has been serialized safely.

## Development

```bash
gofmt -w .
go vet ./...
go test ./...
```

## Goals

- Provide a small, idiomatic Go API for common Schema.org JSON-LD types.
- Reduce manually written JSON-LD in Go applications.
- Use Go's `encoding/json` package for serialization.
- Make common structured-data objects easy to compose.
- Keep the package lightweight and dependency-free.
- Make invalid or inconsistent schema values harder to introduce through constructors.

## Status

The package is under active development. The public API may still evolve as additional Schema.org types and validation helpers are added.

## License

Licensed under the MIT License. See [LICENSE](LICENSE) for details.
