package main

import (
	"fmt"
	"io/ioutil"
)

/*
 Defining the data structures for our wiki.Page struct
 has two fields title and body. The body element is a byte slice
 rather then a string becuase that is the type expected by the
 io libraries.
*/
type Page struct {
	Title string
	Body  []byte // a byte slice
}

/*
Save method is to persist Page data
This method signature reads - "This is a method named save
that takes as its receiver p a pointer to Page. It take no parameter and
returns a value of type error"

WriteFile is a standard library function that writes a byte slice to a file

While writing the file, if anything goes wrong WriteFile will return an error.
If all goes well, save method will return nil ( zero value for pointers,
interfaces and some other types.)
*/
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

/*
Loadpage function read a given file retruns a pointer to a Page literal
constructed with the proper title and body values.

function can retrun multiple values, the standard library function
io.ReadFile returns []byte and error. Here you can collect the body slice
and ignore the error by using underscore symbol (blank identifier).
*/
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample page")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}
