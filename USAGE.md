<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	criblcontrolplanesdkgo "github.com/Cribl-Community/cribl-control-plane-sdk-go"
	"github.com/Cribl-Community/cribl-control-plane-sdk-go/models/components"
	"log"
)

func main() {
	ctx := context.Background()

	s := criblcontrolplanesdkgo.New(
		"https://api.example.com",
	)

	res, err := s.Auth.Tokens.Get(ctx, components.LoginInfo{
		Password: "6j50J9421x29IhO",
		Username: "Lilly_Weissnat",
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