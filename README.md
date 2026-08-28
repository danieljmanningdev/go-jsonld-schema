# go-jsonld-schema

[![CI](https://github.com/danieljmanningdev/go-jsonld-schema/actions/workflows/ci.yml/badge.svg)](https://github.com/danieljmanningdev/go-jsonld-schema/actions/workflows/ci.yml)

`go-jsonld-schema` is a lightweight Go package for building and
marshaling Schema.org JSON-LD structured data.

It provides typed Go structs and constructors for common structured data
types so JSON-LD can be generated with Go instead of manually writing
JSON inside HTML templates.

## Supported Schemas

Currently supported:

-   WebSite
-   Person
-   Organization
-   Article
-   BreadcrumbList
-   SearchAction
-   ImageObject
-   ListItem

## Installation

``` bash
go get github.com/danieljmanningdev/go-jsonld-schema
```

Import the schema package:

``` go
import "github.com/danieljmanningdev/go-jsonld-schema/schema"
```

## WebSite

Create a basic WebSite schema:

``` go
site := schema.NewWebsite(
    "Example Site",
    "https://example.com",
)

data, err := schema.MarshalIndent(site)
if err != nil {
    log.Fatal(err)
}

fmt.Println(string(data))
```

Output:

``` json
{
  "@context": "https://schema.org",
  "@type": "WebSite",
  "name": "Example Site",
  "url": "https://example.com"
}
```

### WebSite with SearchAction

``` go
site := schema.NewWebsite(
    "Example Site",
    "https://example.com",
)

site.PotentialAction = &schema.SearchAction{
    Typed: schema.Typed{
        Type: "SearchAction",
    },
    Target:     "https://example.com/search?q={search_term_string}",
    QueryInput: "required name=search_term_string",
}
```

## Person

``` go
person := schema.NewPerson("Daniel Manning")

person.URL = "https://example.com"
person.PictureURL = "https://example.com/profile.jpg"
person.JobTitle = "Digital Product Designer & Engineer"

person.SocialProfiles = []schema.SocialProfile{
    "https://github.com/example",
    "https://linkedin.com/in/example",
}
```

## Organization

The Go API uses the British spelling `Organisation`, while the generated
Schema.org `@type` is `Organization`.

``` go
organisation := schema.NewOrganisation(
    "Example Ltd",
    "https://example.com",
)

organisation.Logo = "https://example.com/logo.png"
```

## Person with Organization

``` go
person := schema.NewPerson("Daniel Manning")

organisation := schema.Organisation{
    Root: schema.Root{
        Type: "Organization",
    },
    Name: "Example Ltd",
    URL:  "https://example.com",
}

person.WorksFor = &organisation
```

## BreadcrumbList

``` go
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
    schema.NewListItem(
        3,
        "Article",
        "https://example.com/blog/article",
    ),
})
```

## Article

``` go
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

Organization authors are also supported:

``` go
author := schema.NewOrganisationAuthor("Example Ltd")
article.Author = &author
```

## Marshaling

The package provides standard and indented JSON marshaling helpers.

### Compact JSON

``` go
data, err := schema.Marshal(site)
if err != nil {
    log.Fatal(err)
}
```

### Indented JSON

``` go
data, err := schema.MarshalIndent(site)
if err != nil {
    log.Fatal(err)
}
```

## Using with Go HTML Templates

The generated JSON can be passed into a Go HTML template and rendered
inside a JSON-LD script element.

For applications using `html/template`, take care to integrate generated
JSON using an appropriate trusted template type only after the data has
been produced by `encoding/json`. Do not build JSON-LD by concatenating
untrusted strings manually.

Example template:

``` html
{{ with .JSONLD }}
<script type="application/ld+json">{{ . }}</script>
{{ end }}
```

## Development

Format, vet and test the package with:

``` bash
gofmt -w .
go vet ./...
go test ./...
```

## Project Structure

``` text
.
├── go.mod
├── LICENSE
├── README.md
└── schema
    ├── article.go
    ├── base.go
    ├── breadcrumb.go
    ├── organisation.go
    ├── person.go
    ├── schema.go
    ├── schema_test.go
    └── website.go
```

## Goals

The project aims to:

-   Provide a small, idiomatic Go API for common Schema.org JSON-LD
    types.
-   Reduce manually written JSON-LD in Go web applications.
-   Use Go's `encoding/json` package for serialization.
-   Make common structured-data objects easy to compose.
-   Keep the package lightweight and dependency-free.
-   Make invalid or inconsistent schema values harder to introduce
    through constructors.

## Status

The package is currently under active development. The public API may
change while additional schema types, validation and tests are added.

## License

Licensed under the MIT License. See [LICENSE](LICENSE) for details.
