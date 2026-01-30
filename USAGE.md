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

	res, err := s.DatabaseConnections.Create(ctx, components.DatabaseConnectionConfig{
		AuthType:          "connectionString",
		ConfigObj:         criblcontrolplanesdkgo.Pointer("<value>"),
		ConnectionString:  criblcontrolplanesdkgo.Pointer("mysql://admin:password123@mysql.example.com:3306/production?ssl=true"),
		ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
		CredsSecrets:      criblcontrolplanesdkgo.Pointer("<value>"),
		DatabaseType:      components.DatabaseConnectionTypeMysql,
		Description:       "Production MySQL database for customer data",
		ID:                "mysql-prod-db",
		Password:          criblcontrolplanesdkgo.Pointer("QpvMa8DI_lUJL_b"),
		RequestTimeout:    criblcontrolplanesdkgo.Pointer[float64](4657.19),
		Tags:              criblcontrolplanesdkgo.Pointer("production,mysql,customer-data"),
		User:              criblcontrolplanesdkgo.Pointer("Dolores.Feil"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.CountedDatabaseConnectionConfig != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->