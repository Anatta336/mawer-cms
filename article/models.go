package article

type Article struct {
	Title   string
	Content []Content
}

type Content interface {
	String() string
	Html() string
}

type Paragraph struct {
	Content string
}

type Heading struct {
	Title string
}

func (p Paragraph) String() string {
	return "[Paragraph] " + p.Content
}

func (p Paragraph) Html() string {
	return "<p>" + p.Content + "</p>"
}

func (h Heading) String() string {
	return "[Title] " + h.Title
}

func (h Heading) Html() string {
	return "<h2>" + h.Title + "</h2>"
}

func (a Article) String() string {
	return a.Title
}
