<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	"log"
)

func main() {
	ctx := context.Background()

	s := criblcontrolplanesdkgo.New(
		"https://api.example.com",
	)

	res, err := s.Auth.Login(ctx, components.LoginInfo{
		Username: "Nikko.Connelly",
		Password: "Ljp4BunfMR9hNyM",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.AuthToken != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->