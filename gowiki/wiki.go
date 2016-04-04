package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"text/template"
)

/*
Creating a global variable name tempaltes, and initialize it with ParseFiles.
tempate.Must is a convenience wrapper that panics when passed a non-nil error
value, otherwise it returns the *Template

template.ParseFiles will read the contents of given html file and returns
a *template.Template
*/
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

/*
global variable to store our validation regex
regexp.MustCompile is distinict from Compile in that it will panic if the
expression compilation fails, while Compile returns an error as a second parameter.
*/
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

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
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
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
func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

/*
SaveHandler function save form data
*/
func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	// Any error occured during p.save() will be reported to the user.
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

/*
RenderTemplate - refactoring view and editHandler function
*/
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

/*
GetTitle validate path and extract the page title
*/
func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		// if the title is invalid, the function will write a "404 not found" error.
		http.NotFound(w, r)
		return "", errors.New("Invalid Page Title")
	}
	// If the title is valid, it will return along with a nil error value.
	return m[2], nil
}

/*
MakeHandler function is a wrapper function that takes a function of the handler
type and returns a function of type http.HandlerFunc
*/
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func main() {
	// using this view handler
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.ListenAndServe(":8080", nil)
}
