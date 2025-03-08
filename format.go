package co

// Returns a new color with the format bold.
func (c Color) Bold() Color {
	newFormat := append([]string{bold}, c.format...)
	return Color{text: c.text, bg: c.bg, format: newFormat}
}

// Returns a new color with the format bold.
func Bold() Color {
	newFormat := []string{bold}
	return Color{format: newFormat}
}

// Returns a new color with the format dim.
func (c Color) Dim() Color {
	newFormat := append([]string{dim}, c.format...)
	return Color{text: c.text, bg: c.bg, format: newFormat}
}

// Returns a new color with the format dim.
func Dim() Color {
	newFormat := []string{dim}
	return Color{format: newFormat}
}

// Returns a new color with the format italic.
func (c Color) Italic() Color {
	newFormat := append([]string{italic}, c.format...)
	return Color{text: c.text, bg: c.bg, format: newFormat}
}

// Returns a new color with the format italic.
func Italic() Color {
	newFormat := []string{italic}
	return Color{format: newFormat}
}

// Returns a new color with the format underline.
func (c Color) Underline() Color {
	newFormat := append([]string{underline}, c.format...)
	return Color{text: c.text, bg: c.bg, format: newFormat}
}

// Returns a new color with the format underline.
func Underline() Color {
	newFormat := []string{underline}
	return Color{format: newFormat}
}

// Returns a new color with the format strikethrough.
func (c Color) St() Color {
	newFormat := append([]string{st}, c.format...)
	return Color{text: c.text, bg: c.bg, format: newFormat}
}

// Returns a new color with the format strikethrough.
func St() Color {
	newFormat := []string{st}
	return Color{format: newFormat}
}
