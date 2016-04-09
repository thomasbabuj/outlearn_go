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
