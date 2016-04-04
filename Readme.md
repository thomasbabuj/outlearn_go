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

      # Part 5

        - Error Handling

               - Included error handling in renderTemplate and saveHandler functions.
     # Part 6

        - Template caching

     # Part 7

        - Validation

             - fixing the serious security flaw ( a user an supply an arbitrary path to be read/written on the server)

     # Part 8

         - Function Literals and closures

             - using function literals we are abstracting the error handlers repeated
             code.
             - rewriting the function definition of the each handers to accept a
             title string
             - removing the getTitle code because of the makeHandler
