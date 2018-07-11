# GopherEnv
## A golang package for working with a .env file in the most simple way possible

### The most simple example
.env file
```
HELLO=World
GO4BROKE=Another string
64SD=Will return an empty string, illegal name, starts with a number
```
main.go
```go
import (
    "github.com/theclocker/gopherenv"
    "fmt"
)

func main() {
    fmt.Print(gopherenv.Env("HELLO")) // will print "World"
}
```