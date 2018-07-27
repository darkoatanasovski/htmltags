HTML Strip tags
=====================

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


If you want to keep the inline attributes of the tags, set user third parameter to `false`
```go
stripped, _ := htmltags.Strip("<h1>Header text with <span style=\"color:red\">color</span></h1>", []strings{"span"}, false)
```

Or if you want to strip all tags from the string, and get a pure text, the second parameter has to be
empty array

```go
stripped, _ := htmltags.Strip("<h1>Header text with <span style=\"color:red\">color</span></h1>", []strings{}, false)
```

```go
import(
    "strings"
    "fmt"
    "github.com/darkoatanasovski/htmltags"
)

func main() {
    original := "<div>This is <strong style=\"font-size:50px\">complex</strong> text with <span>children <i>nodes</i></span></div>"
    allowedTags := []strings{"strong", "i"}
    removeInlineAttributes := false
    stripped, _ := htmltags.Strip(original, allowedTags, removeInlineAttributes)
    
    fmt.Println(stripped) //will output Node structure
    fmt.Println(stripped.toString()) //will output string: This is <strong>complex</strong> text with children <i>nodes</i>
}
```

## Development
If you have cloned this repo you will probably need the dependency:

`go get golang.org/x/net/html`


[license-svg]: https://img.shields.io/badge/license-BSD--style+patent--grant-blue.svg
[license-link]: https://github.com/darkoatanasovski/htmltags/blob/master/LICENSE