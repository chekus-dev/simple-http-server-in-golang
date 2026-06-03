# Simple Go HTTP Server

A minimalist web server written in Go (Golang) that serves a custom welcome message on your local machine.

##  What This Code Does

* **Starts a Local Web Server**: Listens on port `8080` of your computer.
* **Handles Routing**: Routes all traffic from the root path (`/`) to a specific function.
* **Serves HTML**: Returns a web page containing a `welcome to simple http server in golang` header when accessed.

---

##  Prerequisites

Make sure you have Go installed on your machine. You can check by running:
```bash
go version
```

---

##  How to Run the Server

Follow these steps to get your server up and running:

### 1. Save the Code
Save the Go code into a file named `main.go`.

### 2. Run the Server
Open your terminal or command prompt, navigate to the directory containing your file, and run:
```bash
go run main.go
```
You should see the message: `starting simple http server in golang on localhost:8080`

### 3. View the Results
Open your preferred web browser and navigate to:
```text
http://localhost:8080
```

### 4. Stop the Server
To shut down the server, press `Ctrl + C` in your terminal window.
