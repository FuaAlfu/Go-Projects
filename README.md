---
stack: GO
lng: all
---

# GO Projects
all about GO Projects, Restfull API, business layers, auth and more..

## online crul builder
[click here](https://tools.w3cub.com/curl-builder)

## uuid generator
[click here](https://www.uuidgenerator.net/)

## go docs
use go docs to look for code, inside any pkg
```
go doc github.com/go-pkg/example
```

## go module
```
go mod init folder-name or www.github.com/userName/repo-name
```

## latest packs
```
go mod tidy
```

## get version of module - analyzing a package to see if it's go module compatible
go list -m -version pkg-name
- example:
```
go list -m -versions github.com/dgrijalva/jwt-go
```

## Got a problem with lunch? GOPATH should be set to
```
export GOPATH=$GOROOT
unset GOROOT
```

##  GO111MODULE set "on" or "off"
```
go env -w GO111MODULE=off
```