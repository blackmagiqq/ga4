# ga4

## package for sending events to Google Analytics 4 using the Measurement Protocol

Measurement Protocol (Google Analytics 4) docs reference: https://developers.google.com/analytics/devguides/collection/protocol/ga4

Debugging tools: https://ga-dev-tools.google/ga4/event-builder/

### Usage

```go
package main

import (
    "fmt"
    "github.com/ad/ga4"
)

func main() {
    client := ga4.NewGA4Client("G-XXXXXXXXXX", "xxxxxxxxxxxxxxxxxxxxx", true)
    cid := ClientID("xxx")
    err := client.SendEvent(
        ClientID: cid,
        ga4.Event{
            Name: "test_event",
            Params: map[string]string{
                "param1": "value1",
                "param2": "value2",
            },
        },
    )
    if err != nil {
        fmt.Println(err)
    }
}
```

    