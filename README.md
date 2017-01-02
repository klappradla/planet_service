# Planet Service

Super tiny JSON API based on the [Gin](https://github.com/gin-gonic/gin) framework, returning an in memory representation for star wars planets - used to simplify benchmarking my [jRuby Ratpack Examples](https://github.com/klappradla/jruby_ratpack_examples).

## Run

Run in dev mode
```sh
$ go run main.go

# visit http://localhost:3000/planets
```

Run in release mode
```sh
# start server in release mode
$ export GIN_MODE=release
$ ./planet_service

# optional: start without logging output
$ export GIN_MODE=release
$ ./planet_service > /dev/null
```
