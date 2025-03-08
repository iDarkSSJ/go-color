package co

import "fmt"

// Returns a new color with the format of the first color and the format of the second color.
func (c Color) Join(cc Color) Color {

	text := c.text
	bg := c.bg
	if c.text == "" {
		text = cc.text
	}
	if c.bg == "" {
		bg = cc.bg
	}

	newFormat := append(c.format, cc.format...)
	return Color{text: text, bg: bg, format: newFormat}
}

// Returns a string with the color and the text.
func (c Color) Paint(s string) string {

	text := c.text
	bg := c.bg

	if text == "" {
		text = "39" // default
	}
	if bg == "" {
		bg = "49" // default
	}
	if c.format != nil {
		for _, f := range c.format {
			text = f + ";" + text
		}
	}

	return fmt.Sprintf("%s%s;%sm%s%s", base, bg, text, s, base+reset)
}