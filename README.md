# Simple Go HTTP Server

A beginner-friendly HTTP server built with Go's standard `net/http` package. The server demonstrates how to register multiple routes and send plain text responses to clients.

## Features

- Built using Go's standard library
- Multiple HTTP routes
- Simple request handlers
- Plain text responses
- Lightweight and easy to understand

## Project Structure

```text
.
├── main.go
└── README.md
```

## Requirements

- Go 1.18 or later

## Installation

Clone the repository:

```bash
git clone <repository-url>
```

Navigate into the project directory:

```bash
cd <project-directory>
```

Run the application:

```bash
go run .
```

The server will start on:

```
http://localhost:3000
```

---

## Available Endpoints

### GET /

There is no handler registered for `/`, so the server returns:

```
404 Not Found
```

---

### GET /user

Returns a welcome message.

**Request**

```http
GET /user HTTP/1.1
Host: localhost:3000
```

**Response**

```
Welcome
```

---

### GET /customer

Returns a greeting for the customer.

**Request**

```http
GET /customer HTTP/1.1
Host: localhost:3000
```

**Response**

```
Hello Customer
```

---

### GET /waiter

Returns a message from the waiter.

**Request**

```http
GET /waiter HTTP/1.1
Host: localhost:3000
```

**Response**

```
What can we get you
```

---

## Testing the Server

Using your browser:

```
http://localhost:3000/user
```

```
http://localhost:3000/customer
```

```
http://localhost:3000/waiter
```

Or using `curl`:

```bash
curl http://localhost:3000/user
```

```bash
curl http://localhost:3000/customer
```

```bash
curl http://localhost:3000/waiter
```

---

## Technologies Used

- Go
- net/http
- fmt
- log

---

## Learning Objectives

This project demonstrates how to:

- Create an HTTP server in Go
- Register routes using `http.NewServeMux`
- Handle incoming HTTP requests
- Write responses using `http.ResponseWriter`
- Start a web server with `http.ListenAndServe`

---

## Future Improvements

- Add JSON responses using `encoding/json`
- Restrict routes to specific HTTP methods (GET, POST, etc.)
- Return custom HTTP status codes
- Add middleware for logging requests
- Create reusable handler functions
- Add unit tests
- Connect to a database
- Build a REST API

---

## Author

**Chekus-dev**
