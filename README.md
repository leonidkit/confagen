## Confagen
Confagen is a CLI utility for Go that generates a golang configuration file from a yaml file.
It is enough to specify the name of the source file and the file to be generated.

### Install
```
go get -u github.com/pkg/errors
go get -u github.com/spf13/viper
go get -u github.com/leonidkit/confagen
```
This will install the `confagen` binary to your $GOPATH/bin directory.

### CLI usage
```
Usage:
  confagen [flags]

Examples:
        confagen --src=/path/to/yaml/config/config.yaml --dst=/path/to/destination/file/config.go

Flags:
      --dst string   destination file
  -h, --help         help for confagen
      --src string   source file
```