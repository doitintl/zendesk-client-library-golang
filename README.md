# Zendesk

[![build](https://github.com/doitintl/zendesk/workflows/build/badge.svg)](https://github.com/doitintl/zendesk/actions?query=workflow%3A"build") [![Go Report Card](https://goreportcard.com/badge/github.com/doitintl/zendesk)](https://goreportcard.com/report/github.com/doitintl/zendesk)
[![GoDoc](http://godoc.org/github.com/doitintl/zendesk/zendesk?status.png)](http://godoc.org/github.com/doitintl/zendesk/zendesk)

Zendesk is a [Zendesk Core API](https://developer.zendesk.com/rest_api/docs/core/introduction) client library (thin wrapper on top of Zendesk REST API) for Go.

This library is used internally at DoiT International, forked initially from the [MEDIGO/go-zendesk](https://github.com/MEDIGO/go-zendesk) repository, and detached from it later on, since the original repository is abandoned and no longer maintained. DoiT will continue to maintain this repository, and will be using it for the next major release of the library.

## Usage

```go
package main

import (
  "log"

  "github.com/doitintl/zendesk/zendesk"
)

func main() {
    client, err := zendesk.NewClient("domain", "username", "password")
    if err != nil {
        log.Fatal(err)
    }
    ticket, err := client.ShowTicket(1)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Requester ID is: %d", *ticket.RequesterID)
}
```

Find the complete API on https://godoc.org/github.com/doitintl/zendesk/zendesk#NewClient


## Development

### Linting

To lint the source code, use the command:

```
$ make lint
```

### Testing

The project contains integration tests that uses the Zendesk API. To execute them you must provide the following values in a `.env` file:

```
ZENDESK_DOMAIN=<your-zendesk-domain>
ZENDESK_USERNAME=<your-zendesk-api-email>
ZENDESK_PASSWORD=<your-zendesk-api-password>
```

Then, to run the test, use the command:

```
$ make test
```

Please note that integration tests will create and alter entities in the configured Zendesk instance.
You most likely want to run them against a [Zendesk Sandbox](https://support.zendesk.com/hc/en-us/articles/203661826-Testing-changes-in-your-sandbox-Enterprise-) instance.
