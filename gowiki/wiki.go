package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
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

/*
ViewHandler will allow users to view a wiki page. It will handle URLs prefixed
with "/view/".
extract the page title from r.URL.Path, the path component of the request URL.
The Path is re-sliced with [len("/view/"):] to drop the leading "/view/"
component of the request path. We are slicing the path because, the path will
invariably begin with "/view/", which is not part of the page's title.
*/
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title) // Load the page data
	// formating the page with a string of simple HTML, and writes it to w, the
	// http.ResponseWriter.
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

/*
EditHandler function loads the page or if it doesn't exist, create an empty Page
struct and displays an HTML form.
*/
func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	fmt.Fprintf(w,
		"<h1>Editing %s</h1>"+
			"<form action=\"/save/$s\" method=\"POST\">"+
			"<textarea name=\"body\">%s</textarea><br />"+
			"<input type=\"submit\" value=\"Save\">"+
			"</form>",
		p.Title, p.Title, p.Body)
}

/*
SaveHandler function save form data
*/
func saveHandler(w http.ResponseWriter, r *http.Response) {

}

func main() {
	// using this view handler
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	http.ListenAndServe(":8080", nil)
}
