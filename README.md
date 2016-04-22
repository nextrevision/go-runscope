# go-runscope

[![Circle CI](https://circleci.com/gh/nextrevision/go-runscope.svg?style=svg)](https://circleci.com/gh/nextrevision/go-runscope)

Go Library for interacting with the Runscope API

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
