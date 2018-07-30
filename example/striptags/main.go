package main

import (
	"fmt"
	"github.com/darkoatanasovski/htmltags"
)

func main() {
	original := "<div>This is <strong style=\"font-size:50px\">complex</strong> text with <span>children <i>nodes</i></span></div>"
	allowableTags := []string{"strong", "i"}
	removeInlineAttributes := false
	stripped, _ := htmltags.Strip(original, allowableTags, removeInlineAttributes)

	fmt.Println(stripped)
	fmt.Println(stripped.ToString())
}
