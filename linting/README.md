# Lint - Static code analysis

**Install:**
```sh
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

When you run the command below, nothing will happen.

**Run:**
```sh
go run main.go
```
However, running the command below will report a failure due to a deprecated library.

**Run:**
```sh
golangci-lint
```

**Message error**
```sh
main.go:11:2: SA1019: httputil.NewClientConn has been deprecated since Go 1.0: Use the Client or Transport in package [net/http] instead. (staticcheck)
        httputil.NewClientConn(nil, nil)
```

# References
[golang-lint](https://golangci-lint.run/)