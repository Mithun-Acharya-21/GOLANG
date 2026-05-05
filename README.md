# Golang Learning Project

A comprehensive Go learning project covering fundamental Go concepts and patterns including goroutines, channels, context handling, pointers, and more.

## 📋 Project Overview

This project demonstrates core Go programming concepts through practical examples and implementations. It serves as a learning resource for understanding Go's concurrency model, memory management, and channel-based communication patterns.

## 📁 File Structure

| File | Purpose |
|------|---------|
| `main.go` | Entry point of the application |
| `pointers.go` | Demonstrates pointer usage with structs and pass-by-reference |
| `pipelines.go` | Shows how to build data processing pipelines using channels and goroutines |
| `context.go` | Illustrates context cancellation and worker coordination |
| `select.go` | Demonstrates channel multiplexing with select statements |
| `convstr.go` | String conversion utilities and examples |
| `first.go` | Basic Go concepts and first examples |
| `loan.go` | Loan-related calculations or data structures |
| `PDR.go` | Additional utility functions |
| `go.mod` | Go module definition |

## 🚀 Features

- **Pointers & References**: Learn how to work with pointers and modify struct data
- **Goroutines**: Concurrent programming with lightweight threads
- **Channels**: Channel-based communication between goroutines
- **Pipelines**: Build composable data processing pipelines
- **Context**: Proper context management for cancellation and deadlines
- **Select Statements**: Multiplexing operations on multiple channels

## 🛠️ Prerequisites

- Go 1.25.5 or higher
- Git

## 📦 Installation

1. Clone the repository:
```bash
git clone https://github.com/Mithun-Acharya-21/GOLANG.git
cd GOLANG
```

2. Install dependencies (if any):
```bash
go mod download
```

## ▶️ Running the Project

Run the main application:
```bash
go run main.go
```

To run specific examples:
```bash
go run main.go pointers.go
go run main.go pipelines.go
go run main.go context.go
```

## 🧪 Examples

### Pointers
Demonstrates how to use pointers to modify struct fields:
```go
updateBalance(&user, -200) // pass pointer to modify
```

### Pipelines
Shows data flowing through channels:
```go
for n := range sq(gen(2, 3, 4)) {
    fmt.Println(n) // 4, 9, 16
}
```

### Context & Cancellation
Safely stop goroutines using context:
```go
ctx, cancel := context.WithCancel(context.Background())
cancel() // stop worker cleanly
```

## 📚 Learning Outcomes

After studying this project, you'll understand:
- How Go's concurrency model works with goroutines and channels
- Proper use of pointers for efficient memory management
- Building composable data processing pipelines
- Context-based cancellation patterns
- Channel multiplexing with select

## 💡 Use Cases

This project is perfect for:
- Beginners learning Go fundamentals
- Understanding Go's concurrency patterns
- Reference implementation for best practices
- Interview preparation for Go development roles

## 🤝 Contributing

Feel free to fork this repository and submit pull requests with improvements or additional examples.

## 📝 License

This project is open source and available under the MIT License.

## 👤 Author

**Mithun Acharya**

- GitHub: [@Mithun-Acharya-21](https://github.com/Mithun-Acharya-21)
- Repository: [GOLANG](https://github.com/Mithun-Acharya-21/GOLANG)

## 📞 Support

If you have questions or suggestions, feel free to open an issue on the [GitHub repository](https://github.com/Mithun-Acharya-21/GOLANG/issues).

---

**Happy Learning! 🎓**
