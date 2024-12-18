package models

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) BodyString() string {
	return string(p.Body)
}
