package document

import (
	"bytes"
	"html"
	"mawer/article"

	"baliance.com/gooxml/document"
)

func extractTextContentFromDocx(data []byte) (*article.Article, error) {
	// Create a reader from the byte slice
	reader := bytes.NewReader(data)

	// Open the document from the reader
	doc, err := document.Read(reader, int64(len(data)))
	if err != nil {
		return nil, err
	}

	// Paragraphs are made up of "runs" which are strings with the same formatting.
	// We're ignoring differences in formatting for now and just extracting the text.
	var parsed []article.Content
	for _, para := range doc.Paragraphs() {
		var content string = ""
		isAllBold := true
		for _, run := range para.Runs() {
			content += html.EscapeString(run.Text())

			if !run.Properties().IsBold() {
				isAllBold = false
			}
		}

		if len(bytes.TrimSpace([]byte(content))) == 0 {
			continue
		}

		if isAllBold {
			parsed = append(parsed, article.Heading{
				Title: content,
			})
			continue
		}

		parsed = append(parsed, article.Paragraph{
			Content: content,
		})
	}

	return &article.Article{
		Title:   extractTitle(&parsed),
		Content: parsed,
	}, nil
}

func extractTitle(parsed *[]article.Content) string {
	if len(*parsed) == 0 {
		return "Untitled Article"
	}

	if heading, ok := (*parsed)[0].(article.Heading); ok {
		*parsed = (*parsed)[1:]

		return heading.Title
	}

	return "Untitled Article"
}
