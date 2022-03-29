# dashd-go Usage Attempt

Both of the following commands exit with `1` error code and no changes to `go.mod`.

## Trying `go get -u`

```
go get -u github.com/dashevo/dashd-go@master
```

> go get: github.com/dashevo/dashd-go@master (v0.0.0-20220315183651-a9800ac93748) requires github.com/dashevo/dashd-go@v0.23.2, not github.com/dashevo/dashd-go@master (v0.0.0-20220315183651-a9800ac93748)

## Trying `go install`:

```
go install github.com/dashevo/dashd-go@master
```

> go install: github.com/dashevo/dashd-go@master (in github.com/dashevo/dashd-go@v0.0.0-20220315183651-a9800ac93748):
  The go.mod file for the module providing named packages contains one or
  more replace directives. It must not contain directives that would cause
  it to be interpreted differently than if it were the main module.