# Candango fsok

Candango **fsok** is a Go library designed to make file and directory
operations simple, reliable, and idiomatic. It provides helpers for common
tasks such as checking existence, copying directories, and working with files —
all with Go best practices in mind.

## Installation

Use `go get` to install:

```sh
go get github.com/candango/fsok
```

Then import in your code:

```go
import "github.com/candango/fsok"
```

## Usage

### Check if a directory exists

```go
import "github.com/candango/fsok/dir"

if dir.Exists("/tmp/mydir") {
    // Directory exists
}
```

### Copy all contents of a directory

```go
import "github.com/candango/fsok/dir"

err := dir.CopyAll("source", "destination")
if err != nil {
    log.Fatal(err)
}
```

### Check if a file exists

```go
import "github.com/candango/fsok/file"

if file.Exists("myfile.txt") {
    // File exists
}
```

## Contributing

Contributions are welcome! Please open issues or pull requests on GitHub.

## License

[MIT](LICENSE)
