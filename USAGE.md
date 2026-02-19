<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := criblcontrolplanesdkgo.New(
		"https://api.example.com",
		criblcontrolplanesdkgo.WithSecurity(components.Security{
			BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
		}),
	)

	res, err := s.LakeDatasets.Create(ctx, "<id>", components.CriblLakeDataset{
		ID: "<id>",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.CountedCriblLakeDataset != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->