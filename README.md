# 🎫 Go Ticket Booking Application

A simple, fast, and concurrent **ticket booking CLI application** built using **Golang**, demonstrating clean architecture, input validation, and real-time booking logic with asynchronous operations using goroutines.

---

## ✨ Features
- Modular code with a helper package for input validation  
- Concurrent operations using **goroutines** and **sync.WaitGroup**  
- Custom `UserData` struct to store booking details  
- Real-time ticket tracking using slice-based storage  
- Simulated email confirmation using concurrency  
- Clean, readable, beginner-friendly Go code design  

---

## 🛠 Tech Stack
- Go (Golang)  
- Goroutines & WaitGroups  
- Structs & Slices  
- Modular Packages  

---

## 🚀 How to Run
```bash
go run main.go helper.go

```bash
go run main.go


📂 Project Structure
booking_app/
 ├── main.go
 ├── go.mod
 └── helper/
      └── validate.go
