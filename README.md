# Go
1. Running:
```
go run main.go
```

2. Creating a mod
```
go mod init github.com/YourUsername/YourRepository
```

2. Building
```
go build -o hello
```

# Identifiers
- function starting with lowerCase (Eg: sayHello): means **private**
- function starting with upperCase (Eg: SayHello): means **public**

# Slices
- São "arrays" de tamanho variável
- Exemplo:
```go
func main() {
  var gavetas[] string
  gavetas = append(gavetas, "copos", "panos", "pratos")
  fmt.Println(gavetas[:2]) // do índice 0 até 1: copos, panos
  fmt.Println(gavetas[2:]) // do índice 2 até o final: pratos
}
```
