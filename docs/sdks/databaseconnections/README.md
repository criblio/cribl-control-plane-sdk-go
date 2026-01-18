# DatabaseConnections

## Overview

### Available Operations

* [Create](#create) - Create Database Connection
* [Delete](#delete) - Delete a Database Connection
* [Get](#get) - Get a Database Connection
* [Update](#update) - Update a Database Connection

## Create

Create a new Database Connection.

### Example Usage

<!-- UsageSnippet language="go" operationID="createDatabaseConnectionConfig" method="post" path="/lib/database-connections" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        AuthType: "connectionString",
        ConfigObj: criblcontrolplanesdkgo.Pointer("<value>"),
        ConnectionString: criblcontrolplanesdkgo.Pointer("mysql://admin:password123@mysql.example.com:3306/production?ssl=true"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
        CredsSecrets: criblcontrolplanesdkgo.Pointer("<value>"),
        DatabaseType: components.DatabaseConnectionTypeMysql,
        Description: "Production MySQL database for customer data",
        ID: "mysql-prod-db",
        Password: criblcontrolplanesdkgo.Pointer("QpvMa8DI_lUJL_b"),
        RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](4657.19),
        Tags: criblcontrolplanesdkgo.Pointer("production,mysql,customer-data"),
        User: criblcontrolplanesdkgo.Pointer("Dolores.Feil"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedDatabaseConnectionConfig != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [components.DatabaseConnectionConfig](../../models/components/databaseconnectionconfig.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.CreateDatabaseConnectionConfigResponse](../../models/operations/createdatabaseconnectionconfigresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete the specified Database Connection.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteDatabaseConnectionConfigById" method="delete" path="/lib/database-connections/{id}" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.DatabaseConnections.Delete(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedDatabaseConnectionConfig != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                 | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `ctx`                                                     | [context.Context](https://pkg.go.dev/context#Context)     | :heavy_check_mark:                                        | The context to use for the request.                       |
| `id`                                                      | *string*                                                  | :heavy_check_mark:                                        | The <code>id</code> of the Database Connection to delete. |
| `opts`                                                    | [][operations.Option](../../models/operations/option.md)  | :heavy_minus_sign:                                        | The options for this request.                             |

### Response

**[*operations.DeleteDatabaseConnectionConfigByIDResponse](../../models/operations/deletedatabaseconnectionconfigbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Get the specified Database Connection.

### Example Usage

<!-- UsageSnippet language="go" operationID="getDatabaseConnectionConfigById" method="get" path="/lib/database-connections/{id}" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.DatabaseConnections.Get(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedDatabaseConnectionConfig != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The <code>id</code> of the Database Connection to get.   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetDatabaseConnectionConfigByIDResponse](../../models/operations/getdatabaseconnectionconfigbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update the specified Database Connection.</br></br>Provide a complete representation of the Database Connection that you want to update in the request body. This endpoint does not support partial updates. Cribl removes any omitted fields when updating the Database Connection.</br></br>Confirm that the configuration in your request body is correct before sending the request. If the configuration is incorrect, the updated Database Connection might not function as expected.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateDatabaseConnectionConfigById" method="patch" path="/lib/database-connections/{id}" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.DatabaseConnections.Update(ctx, "<id>", components.DatabaseConnectionConfig{
        AuthType: "connectionString",
        ConfigObj: criblcontrolplanesdkgo.Pointer("<value>"),
        ConnectionString: criblcontrolplanesdkgo.Pointer("mysql://admin:password123@mysql.example.com:3306/production?ssl=true"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
        CredsSecrets: criblcontrolplanesdkgo.Pointer("<value>"),
        DatabaseType: components.DatabaseConnectionTypeMysql,
        Description: "Production MySQL database for customer data",
        ID: "mysql-prod-db",
        Password: criblcontrolplanesdkgo.Pointer("Fu8u0O8uNNPcA9S"),
        RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](7946.16),
        Tags: criblcontrolplanesdkgo.Pointer("production,mysql,customer-data"),
        User: criblcontrolplanesdkgo.Pointer("Rubie93"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedDatabaseConnectionConfig != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `id`                                                                                       | *string*                                                                                   | :heavy_check_mark:                                                                         | The <code>id</code> of the Database Connection to update.                                  |
| `databaseConnectionConfig`                                                                 | [components.DatabaseConnectionConfig](../../models/components/databaseconnectionconfig.md) | :heavy_check_mark:                                                                         | DatabaseConnectionConfig object                                                            |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.UpdateDatabaseConnectionConfigByIDResponse](../../models/operations/updatedatabaseconnectionconfigbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |