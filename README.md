# go-loadenvconf

This is a Go repository to parse from a given file (f.e. .env) and pars it into a given structure.
It is supposed to parse only string values

## Installation
As a library:

```shell
go get github.com/dzwiedz90/go-loadenvconf
```

## Usage

Add your's application configuration to the file (f.e. ".env") into desired location (may be root location of your project), like here:</br>
```shell
USERNAME=JohnDoe
PASSWORD=p@$$w0rD
```
Then in your Go application you can implement it like shown below (please see also examples/example.go):

```go
package main

import (
	"fmt"

	"github.com/dzwiedz90/go-loadenvconf/loadenvconf"
)

type Config struct {
	SECRET_KEY  string
	DB_NAME     string
	DB_USER     string
	DB_PASSWORD string
}

func main() {
	cfg := Config{}
	loadenvconf.LoadEnvConfig("loadenvconf/.env", &cfg)
	fmt.Println(cfg)
}
```