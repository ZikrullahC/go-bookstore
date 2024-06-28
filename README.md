# Go Bookstore

This is a simple bookstore application written in Go, utilizing the Gorilla Mux router and GORM for database interactions. The application allows for basic CRUD operations on books.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Project Structure](#project-structure)
- [Contributing](#contributing)
- [License](#license)

## Installation

To set up the project on your local machine, follow these steps:

1. **Clone the repository:**
   ```sh
   git clone https://github.com/ZikrullahC/go-bookstore.git
   cd go-bookstore ```
   
2. **Install dependencies:**
Ensure you have Go installed on your machine. Then, run:
   ```sh
   go mod tidy
   ```
3. **Set up the database:**  
Create a MySQL database.  
Update the database configuration in pkg/config/app.go to match your database settings.

## Usage

  To run the application, execute the following command:
   ```sh
   go mod tidy
   ```

 ## API Endpoints
  The application provides the following RESTful API endpoints:  

  Create a new book:
  ```sh
  POST /book/  
  ```
  Request body:
  ```sh
  {
  "name": "Book Name",
  "author": "Author Name",
  "publication": "Publication Date"
 }
  ```  
 Get All Books:
 ```sh
 GET /book/
 ```
 
 Get A Book By ID:
 ```sh
 GET /book/{bookID}
 ```
 
 Update a book by ID:
 ```sh
 PUT /book/{bookID}
 ```
 Request body:
 ```sh
 {
  "name": "Updated Book Name",
  "author": "Updated Author Name",
  "publication": "Updated Publication Date"
}
 ```
Delete a book by ID:

```sh
DELETE /book/{bookId}
```

## Project Structure

```sh
go-bookstore/
├── main.go
├── go.mod
├── go.sum
├── pkg/
│   ├── config/
│   │   └── app.go
│   ├── controllers/
│   │   └── book-controller.go
│   ├── models/
│   │   └── book.go
│   └── routes/
│       └── book-routes.go
```
``main.go``: The entry point of the application.  
``pkg/config``: Contains configuration for the application, including database setup.  
``pkg/controllers``: Contains the handlers for the different API endpoints.  
``pkg/models``: Contains the data models and methods for interacting with the database.  
``pkg/routes``: Contains the routing logic for the application.  


















