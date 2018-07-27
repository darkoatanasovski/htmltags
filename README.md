HTML Strip tags
=====================

[![Docs][docs-godoc-svg]][docs-godoc-link]
[![Go Report Card][goreport-svg]][goreport-link]
[![License][license-svg]][license-link]

This is a Go package which removes HTML tags from a string. Also, you can provide an array of `allowedTags` that can be
skipped.
Strip HTML tags library is very useful if you work with web crawlers, and want to remove tags,
or want to keep specific tags.

## Instalation
```bash
$ go get github.com/darkoatanasovski/htmltags
``` 
## Usage


If you want to keep the inline attributes of the tags, set third parameter to `false`
```go
stripped, _ := htmltags.Strip("<h1>Header text with <span style=\"color:red\">color</span></h1>", []string{"span"}, false)
```

Or if you want to strip all tags from the string, and get a pure text, the second parameter has to be
empty array

```go
stripped, _ := htmltags.Strip("<h1>Header text with <span style=\"color:red\">color</span></h1>", []string{}, false)
```

```go
import(
    "fmt"
    "github.com/darkoatanasovski/htmltags"
)

func main() {
    original := "<div>This is <strong style=\"font-size:50px\">complex</strong> text with <span>children <i>nodes</i></span></div>"
    allowedTags := []string{"strong", "i"}
    removeInlineAttributes := false
    stripped, _ := htmltags.Strip(original, allowedTags, removeInlineAttributes)
    
    fmt.Println(stripped) //output: Node structure
    fmt.Println(stripped.ToString()) //output string: This is <strong>complex</strong> text with children <i>nodes</i>
}
```

## Development
If you have cloned this repo you will probably need the dependency:

`go get golang.org/x/net/html`

[docs-godoc-svg]: https://img.shields.io/badge/docs-godoc-blue.svg
[docs-godoc-link]: https://godoc.org/github.com/darkoatanasovski/htmltags
[goreport-svg]: https://goreportcard.com/badge/github.com/darkoatanasovski/htmltags
[goreport-link]: https://goreportcard.com/report/github.com/darkoatanasovski/htmltags
[license-svg]: https://img.shields.io/badge/license-BSD--style+patent--grant-blue.svg
[license-link]: https://github.com/darkoatanasovski/htmltags/blob/master/LICENSE