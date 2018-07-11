# GopherEnv
## A golang package for working with a .env file in the most simple way possible

### The most simple example
.env file **(at project root)**
```
HELLO=World
GO4BROKE=Another string
64SD=Will log fatal and return empty string, illegal name, starts with a number
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

### Set a custom path
config/.env file
```
REDIS_HOST=127.0.0.1
REDIS_PORT=6379
```
main.go
```go
import (
    "github.com/theclocker/gopherenv"
)

func main() {
    gopherenv.SetPath("./config")
}
```

cache.go
```go
import (
    "github.com/go-redis/redis"
    "fmt"
)

func Connect() *redis.Client {
    return redis.NewClient(&redis.Options{
        Addr: fmt.Sprintf("%s:%s", gopherenv.Env("REDIS_HOST"), gopherenv.Env("REDIS_PORT")),
        Password: "",
        DB: 0,
    })
}

```