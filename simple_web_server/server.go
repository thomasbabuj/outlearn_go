/*
Simple Web Server
*/

package main

import (
	"fmt"
	"net/http"
)

/*
Handler function is of the type http.HandleFunc. it takes an http.ResponseWriter
and an http.Request as its arguments.

http.ResponseWriter value assembles the HTTP server's response. by writing to it,
we send data to the HTTP client.

http.Request is a data structure that represents the client HTTP request.
r.URL.Path is the path component of the request URL.

the trailling [1:] means, create a sub-slice of Path from the 1st character to
the end. This drops the leading "/" from the path name.
*/
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

/*
The main function begins with http.HandleFunc which tells the http pacakge to
handle all requests to the web root "/" with handler.

It then calls http.ListenAndServe, specifying that it should listen on port 8080
on any interface. This function will block until the program is terminated.
*/
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
