# CRM API using Fiber Framework

This project is a simple CRM (Customer Relationship Management) API built using the **Fiber** framework in **Go**, with **GORM** for database management and **SQLite** as the database. The application allows users to manage leads with basic CRUD operations.

## Features
- Retrieve all leads
- Retrieve a specific lead by ID
- Create a new lead
- Delete a lead

## Project Structure
```
crm-go/
│── main.go         # Entry point of the application
│── database/
│   └── database.go # Database connection setup
│── lead/
│   └── lead.go     # Lead model and handlers
│── go.mod          # Go module dependencies
│── go.sum          # Checksums for dependencies
```

## Installation & Setup
### Prerequisites
Ensure you have Go installed on your machine. You can download it from [golang.org](https://golang.org/dl/).

### Clone the repository
```sh
git clone https://github.com/Aman-Shitta/crm-go.git
cd crm-go
```

### Install dependencies
```sh
go mod tidy
```

### Run the application
```sh
go run main.go
```

The server will start on `http://localhost:8000`.

## API Endpoints
| Method | Endpoint            | Description            |
|--------|---------------------|------------------------|
| GET    | `/api/v1/lead`      | Get all leads         |
| GET    | `/api/v1/lead/:id`  | Get a lead by ID      |
| POST   | `/api/v1/lead`      | Create a new lead     |
| DELETE | `/api/v1/lead/:id`  | Delete a lead         |

## Code Overview

### `main.go`
- Initializes the database connection
- Sets up API routes
- Starts the Fiber web server

### `database/database.go`
- Manages the connection to the SQLite database

### `lead/lead.go`
- Defines the `Lead` model
- Implements CRUD operations for leads

## Learning Resources
This project is inspired by the tutorial:
["Building a CRM with Golang and Fiber"](https://youtu.be/jFfo23yIWac)

## Dependencies
- [Fiber](https://gofiber.io/) - Web framework for Go
- [GORM](https://gorm.io/) - ORM for Go
- [SQLite](https://www.sqlite.org/) - Database

## License
This project is open-source and available under the Apache2.0 License.

