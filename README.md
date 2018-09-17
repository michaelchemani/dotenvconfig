# dotenvconfig

Super simple Go(lang) implementation of dotenv _(inspired by: [https://github.com/alexsasharegan/dotenv](https://github.com/alexsasharegan/dotenv))_ with a special environment based twist.

It is often helpful to have different configuration values based on the environment where the application is running. For example, you may wish to use a different cache driver locally than you do on your production server.

To make this a cinch, one can use dotenvconfig.

## Installation

```sh
go get github.com/michaelchemani/dotenvconfig
```

## Usage

In your environment file (named `{your_environment}.env`):

```sh
SECRET_KEY=mya======somesecretKey&^%

MESSAGE="A message containing important spaces."

# A comment line that will be ignored
DB_PORT=github.com
```

In your application:

```go
package main

import (
    "os"

    "github.com/michaelchemani/dotenvconfig"
)

func main() {
    environment := os.Getenv("APP_ENV")
    //Environment = staging for example

    config.Load("config/" + environment + ".env")

    fmt.Println(os.Getenv("DB_NAME"))
}
```

## Documentation

[https://godoc.org/github.com/michaelchemani/dotenvconfig](https://godoc.org/github.com/michaelchemani/dotenvconfig)