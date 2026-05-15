# 🚀 Golang Fundamentals: RESTful API Task Manager

This repository serves as my personal workspace for mastering **Go (Golang)**. What started as a basic CLI tool has now evolved into a fully functional RESTful API, demonstrating core backend architecture principles.

## 🎯 Core Features Implemented
* **RESTful API Architecture:** Built a local web server utilizing Go's native `net/http` package.
* **HTTP Methods Handling:** Implemented routing logic to handle both `GET` (Read) and `POST` (Create) requests effectively.
* **Data Persistence:** Engineered an automated JSON I/O system (`Marshal`/`Unmarshal`) to permanently store incoming API payloads to the hard drive.
* **Modular Codebase:** Clean separation of concerns using Go Packages (structs and methods isolated in a `models` package).

## 🛠️ How to Test the API
1. Run the server:
   ```bash
   go run main.go
