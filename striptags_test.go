package htmltags

import (
	"testing"
)

func TestStrip(t *testing.T) {
	tests := []struct {
		Input                 string
		AllowedTags           []string
		StripInlineAttributes bool
		Want                  string
	}{
		{"", []string{}, true, ""},
		{"<p>Content</p>", []string{}, true, "Content"},
		{"<p>This is content <span>with child tags</span></p>", []string{}, true, "This is content with child tags"},
		{"<p>This is content <span>with child tags</span></p>", []string{"span"}, true, "This is content <span>with child tags</span>"},
		{"<p>This is content <span style=\"font-size:20px;\">with child tags</span></p>", []string{"span"}, true, "This is content <span>with child tags</span>"},
		{"<p>This is content <span style=\"font-size:20px;\">with child tags</span></p>", []string{"span"}, false, "This is content <span style=\"font-size:20px;\">with child tags</span>"},
		{"<p>This <p><span> foo <i>is</i> bar</span></p> content <span>with <em>child</em> tags</span></p>", []string{"span", "i", "em"}, true, "This <span> foo <i>is</i> bar</span> content <span>with <em>child</em> tags</span>"},
	}

	for _, test := range tests {
		if got, _ := Strip(test.Input, test.AllowedTags, test.StripInlineAttributes); got.ToString() != test.Want {
			t.Errorf("%q: want %q, got %q", test.Input, test.Want, got.ToString())
		}
	}
}
