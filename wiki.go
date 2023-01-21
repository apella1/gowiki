package main

import (
	"fmt"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

// addressing persistent storage - saving pages
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)

}

// loading pages

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample page")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(p2.Body)
}
