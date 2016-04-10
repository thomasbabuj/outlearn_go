Step 1

  A Basic Web Server

    - RESTful service starts with fundamentally being a web service first.

Step 2

  Adding a Router

    - we will be using mux router package. this Package  implements a request router and dispatcher.
    - now we have created a basic router, adds the route "/" and assigns the Index handler to run when the endpoint is called.

  Creating some basic routes

    - added two routes ( todoIndex and todoShow )

Step 3

   Creating a basic Model

    - In Go a struct will typically serve as model.

Step 4

   Send back some JSON

Step 5

    Refactoring

      - the Todo Model, by adding the struct tags we can control how our struct will be marshalled to JSON

      - breaking the code into separate files.
      - After refactoring the code into multiple files, we just need to use *go build* or *go install* to compiled the source code.
      - For *go run* need to manually specify all the files.

Step 6

    Outputting a Web Log

      - in go there is no web logging package or functionality in the standard library, so we have to create it.
      - We will be passing our handler to this function, which will then wrap the passed handler with logging and timing functionality.

Step 7

    Applying the logger decorator

Step 8

    Refactoring the Routes

       - Separate the NewRouter into a new files
       - adding the content-type and telling the client to expect json and we are explicitly setting the status code.

Step 9

    Adding a mock database

      - added additional handlers like create, destroy and find handlers
      - added Id to the Todo model
      - in the body of the request, we use io.LimitReader which protect against malicious attacks on the server.
      - after read the body we Unmarshal it to our Todo struct, it that fails, we will respond with the appropriate status code, but will also send back the error in a json string.
      - if all gone well, then we return 201 status code which means the entity was successfully created.
