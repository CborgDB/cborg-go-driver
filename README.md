# CborgDB Go Driver

![version](https://img.shields.io/github/v/tag/cborgdb/cborg-go-driver?color=red&label=cborg-go-driver)
[![Go Report Card](https://goreportcard.com/badge/github.com/cborgdb/cborg-go-driver)](https://goreportcard.com/report/github.com/cborgdb/cborg-go-driver)
![Go Dependency](https://img.shields.io/badge/go->=1.19-blue)

The CborgDB Go Driver is 🚧 still under development 🚧 and [CborgDB](https://github.com/cborgdb/cborg) also.

## Requirements

- Go 1.19 or higher
- CborgDB v0.6.0
  - [GitHub](https://github.com/CborgDB/cborg/releases/tag/v0.6.0)
  - [Docker](https://hub.docker.com/layers/cborgdb/cborg/0.6.0/images/sha256-b4c8c69423c1bdf42f0319737ce5a8d64432caffee1827b25ffc79e078d552f7)

## Install

- Download

`go get github.com/cborgdb/cborg-go-driver`

- Import

`import "github.com/cborgdb/cborg-go-driver/cborg"`

## Usage

1) Run docker container

```console
docker run -d -p 30000:30000 cborgdb/cborg:0.6.0
```

2) Run the following sample code

```golang
package main

import (
	"os"

	"github.com/cborgdb/cborg-go-driver/cborg"
)

func main() {
	client := cborg.NewClient("127.0.0.1", "30000")
	err := client.Connect()
	if err != nil {
		os.Exit(1)
	}
	defer client.Disconnect()

	// Create Database
	result, _ := client.CreateDatabase("cborgdb")
	// Create Collection using previous result
	result.Database.CreateCollection("numbers")

	// Init "number" Collection
	numbers := client.Database("cborgdb").Collection("numbers")
	// Insert an item
	numbers.InsertOne(1992)
}
```

## License

Copyright © 2022 Adil Benhlal 

The CborgDB Go Driver is licensed under the [Apache License](LICENSE).
