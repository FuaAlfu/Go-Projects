---
stack: Go, cli
lang: all
---

# modern cli in golang
In this code, the *cobra* library is used to create a root command with a description and a subcommand for printing the version number of the application. 
The *viper* library is used to read configuration files, which are expected to be in YAML format and located in the user's home directory.


## pkg
```
go mod tidy
```

### cobra
```
go get github.com/spf13/cobra
```

### viper
```
 go get github.com/spf13/viper
```