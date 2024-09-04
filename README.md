# FirstGoApp

This project demonstrates a basic web server in Go that uses handlers, renders templates, and serves HTML pages. The project structure includes a main entry point, handler functions, and a template rendering utility.

## Project Structure

```
.
├── main.go
├── pkg
│   ├── handlers
│   │   └── handlers.go
│   └── render
│       └── render.go
└── templates
    ├── home.page.tmpl
    └── about.page.tmpl
```

## Files

### `main.go`

This is the entry point of the application. It sets up the HTTP server and routes.

```go
package main

import (
	"fmt"
	"log"
	"main/pkg/handlers"
	"net/http"
)

const portNumber = ":8000"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting server on port %s", portNumber))
	if err := http.ListenAndServe(portNumber, nil); err != nil {
		log.Fatal(err)
	}
}
```

### `pkg/render/render.go`

This file contains the template rendering logic. It parses and executes HTML templates.

```go
package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error Parsing template: ", err)
		return
	}
}
```

### `pkg/handlers/handlers.go`

This file contains the handler functions for different routes.

```go
package handlers

import (
	"main/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
```

### `templates/home.page.tmpl`

This is the HTML template for the home page.

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Home Page</title>
</head>
<body>
    <h1>Welcome to the Home Page</h1>
</body>
</html>
```

### `templates/about.page.tmpl`

This is the HTML template for the about page.

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>About Page</title>
</head>
<body>
    <h1>Welcome to the About Page</h1>
</body>
</html>
```

## Running the Project

1. Ensure you have Go installed on your machine.
2. Clone the repository or copy the project structure to your local machine.
3. Navigate to the project directory.
4. Run the application using the following command:

```sh
go run main.go
```

5. Open your web browser and navigate to `http://localhost:8000` to see the home page.
6. Navigate to `http://localhost:8000/about` to see the about page.

## Conclusion

This project provides a simple experience for me of how to set up a basic web server in Go with handlers, template rendering, and HTML templates. You can expand this project by adding more routes, templates, and functionality as needed.

---
