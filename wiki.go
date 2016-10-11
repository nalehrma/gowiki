package main

//importing formating package and io package
import (
	"fmt"
	"io/ioutil"
)

//Page defines struct for each page, needing title and body (the content of the page)
//body is a []byte slice since the io packages read that data type
type Page struct {
	Title string
	Body  []byte
}

//Page function saves the body to a txt file, used by calling Page.save()
func (p *Page) save() error {
	filename := p.Title + ".txt"
	//0600 refers to chmod permission that owner can read and write
	return ioutil.WriteFile(filename, p.Body, 0600)
}

//loadPage creates the file name, reads the body, and returns a pointer to Page with proper formatting
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil

}

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}
