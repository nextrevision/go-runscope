# go-runscope

[![Circle CI](https://circleci.com/gh/nextrevision/go-runscope.svg?style=svg)](https://circleci.com/gh/nextrevision/go-runscope)
[![GoDoc](https://godoc.org/github.com/nextrevision/go-runscope?status.svg)](https://godoc.org/github.com/nextrevision/go-runscope)

Go Library for interacting with the Runscope API

## Usage

To use library in your project, first create a new client:

```
package main

import "github.com/nextrevision/go-runscope"


func main() {
  token := os.Getenv("RUNSCOPE_TOKEN")
  client := runscope.NewClient(&runscope.Options{
    Token: token,
  })
}
```

Once you have a client, you can start interacting with the Runscope API. For example, to list all buckets:

```
buckets, err := client.ListBuckets()
if err != nil {
  ...
}

for _, bucket := range *buckets {
  println(bucket.Name)
}
```

To list all tests in a bucket:

```
tests, _ := client.ListTests(bucket.Key)
for _, test := range *tests {
  println(test.Name)
}
```

For a comprehensive list of all the methods available, please referece [https://godoc.org/github.com/nextrevision/go-runscope](https://godoc.org/github.com/nextrevision/go-runscope).

## Developing

Dependencies are managed with [glide](https://github.com/Masterminds/glide) using the new vendoring support in Go. To add a new dependency, simply type:

```
glide get <package_name>
```

To test, run:

```
go vet
go test -v
```
