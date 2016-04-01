Building Web services in Go - Curated articles by Mat Ryer in Outlearn.com

1) Writing Web applications
     - By reading the article from golang.org (https://golang.org/doc/articles/wiki/)

     # Part 1

        - Created a simple data structure and the ability to save to and load from a file.

     # Part 2

         - Introducing the net/http package - create a simple web server

     # Part 3

         - Using net/http to serve wiki pages

     # Part 4

         - Ability to edit pages.

               - creating a new handler for edit page
               - using html/template package instead of hard-coded HTML and modify the editHandler function to use html/template
               - creating template for viewHanlder and modify handler function
               - refactoring view and editHanlder code, by introducing renderTemplate function since both contains same templating code.
               - handling the non-existent pages
               - saveHanlders will handle the form submission.
