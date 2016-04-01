package main

import (
	"io/ioutil"
	"net/http"
	"text/template"
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
	// handling non-existent pages
	p, err := loadPage(title) // Load the page data
	if err != nil {
		// if there is no page, then redirect to edit route and add 302 status in
		// localtion header
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	// formating the page with a string of simple HTML, and writes it to w, the
	// http.ResponseWriter.
	// instead of html string using html/template package to load html
	renderTemplate(w, "view", p)
}

/*
EditHandler function loads the page or if it doesn't exist, create an empty Page
struct and displays an HTML form.

html/template package is part of the Go standard library. we can use html/template
to keep the HTML in a separate file, allowing us to change the layout of our edit
page without modifying the underlying Go code.
*/
func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

/*
RenderTemplate - refactoring view and editHandler function
*/
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	// template.ParseFiles will read the contents of given html file and returns
	// a *template.Template
	t, _ := template.ParseFiles(tmpl + ".html")
	// this method executes the template, writing the generated HTML to the
	// http.ResponseWriter.
	t.Execute(w, p)
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
	//http.HandleFunc("/save/", saveHandler)
	http.ListenAndServe(":8080", nil)
}
