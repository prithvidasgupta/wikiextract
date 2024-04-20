package xsd

import (
	"strings"
)

func (p *Page) RemoveTemplates() *Page {
	var result strings.Builder
	curlyBraces := 0
	strLen := len(p.Revisions[0].Text)
	for i := 0; i < strLen; i++ {
		if i != strLen-1 && p.Revisions[0].Text[i] == '{' && p.Revisions[0].Text[i+1] == '{' {
			curlyBraces += 1
			i += 1
			continue
		} else if i != strLen-1 && p.Revisions[0].Text[i] == '}' && p.Revisions[0].Text[i+1] == '}' {
			curlyBraces -= 1
			i += 1
			continue
		}

		if curlyBraces == 0 {
			result.WriteByte(p.Revisions[0].Text[i])
		}
	}
	p.Revisions[0].Text = result.String()
	return p
}

func (p *Page) RemoveTables() *Page {
	var result strings.Builder
	tables := 0
	strLen := len(p.Revisions[0].Text)
	for i := 0; i < strLen; i++ {
		if i != strLen-1 && p.Revisions[0].Text[i] == '{' && p.Revisions[0].Text[i+1] == '|' {
			tables += 1
			i += 1
			continue
		} else if i != strLen-1 && p.Revisions[0].Text[i] == '|' && p.Revisions[0].Text[i+1] == '}' {
			tables -= 1
			i += 1
			continue
		}

		if tables == 0 {
			result.WriteByte(p.Revisions[0].Text[i])
		}
	}
	p.Revisions[0].Text = result.String()
	return p
}

func cleanPartition(partStr string, splitter string) string {
	parts := strings.Split(partStr, splitter)
	if splitter == "|" && strings.HasPrefix(parts[0], "Category:") {
		return ""
	}
	return parts[len(parts)-1]
}

func (p *Page) HandleInternalLinks() *Page {
	var result strings.Builder
	var partition strings.Builder
	bracketCounter := 0
	strLen := len(p.Revisions[0].Text)
	for i := 0; i < strLen; i++ {
		if i != strLen-1 && p.Revisions[0].Text[i] == '[' && p.Revisions[0].Text[i+1] == '[' {
			bracketCounter += 1
			i += 1
			continue
		} else if i != strLen-1 && p.Revisions[0].Text[i] == ']' && p.Revisions[0].Text[i+1] == ']' {
			bracketCounter -= 1
			i += 1
			if bracketCounter == 0 {
				t := cleanPartition(partition.String(), "|")
				if t != "" {
					result.WriteString(t)
				}
				partition.Reset()
			}
			continue
		}

		if bracketCounter == 0 {
			result.WriteByte(p.Revisions[0].Text[i])
		} else {
			partition.WriteByte(p.Revisions[0].Text[i])
		}
	}
	p.Revisions[0].Text = result.String()
	return p
}

func (p *Page) HandleExternalLinks() *Page {
	var result strings.Builder
	var partition strings.Builder
	bracketCounter := 0
	strLen := len(p.Revisions[0].Text)
	for i := 0; i < strLen; i++ {
		if i != strLen-1 && p.Revisions[0].Text[i] == '[' {
			bracketCounter += 1
			continue
		} else if i != strLen-1 && p.Revisions[0].Text[i] == ']' {
			bracketCounter -= 1
			if bracketCounter == 0 {
				t := cleanPartition(partition.String(), " ")
				if t != "" {
					result.WriteString(t)
				}
				partition.Reset()
			}
			continue
		}

		if bracketCounter == 0 {
			result.WriteByte(p.Revisions[0].Text[i])
		} else {
			partition.WriteByte(p.Revisions[0].Text[i])
		}
	}
	p.Revisions[0].Text = result.String()
	return p
}

func (p *Page) HandleHTMLTags() *Page {
	var result strings.Builder
	bracketCounter := 0
	strLen := len(p.Revisions[0].Text)
	for i := 0; i < strLen; i++ {
		if i != strLen-1 && p.Revisions[0].Text[i] == '<' {
			bracketCounter += 1
			continue
		} else if i != strLen-1 && p.Revisions[0].Text[i] == '>' {
			bracketCounter -= 1
			continue
		}

		if bracketCounter == 0 {
			result.WriteByte(p.Revisions[0].Text[i])
		}
	}
	p.Revisions[0].Text = result.String()
	return p
}
