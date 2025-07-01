<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
	ctx := context.Background()

	s := criblcontrolplanesdkgo.New(
		"https://api.example.com",
	)

	res, err := s.Diag.GetHealthInfo(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.HealthStatus != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->