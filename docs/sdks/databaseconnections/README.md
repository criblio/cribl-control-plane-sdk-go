# DatabaseConnections

## Overview

### Available Operations

* [Create](#create) - Create Database Connection
* [Delete](#delete) - Delete a Database Connection
* [Get](#get) - Get a Database Connection
* [Update](#update) - Update a Database Connection

## Create

Create a new Database Connection.

### Example Usage: DatabaseConnectionExamplesMySQLWithConnectionString

<!-- UsageSnippet language="go" operationID="createDatabaseConnectionConfig" method="post" path="/lib/database-connections" example="DatabaseConnectionExamplesMySQLWithConnectionString" -->
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
### Example Usage: DatabaseConnectionExamplesMySQLWithSecret

<!-- UsageSnippet language="go" operationID="createDatabaseConnectionConfig" method="post" path="/lib/database-connections" example="DatabaseConnectionExamplesMySQLWithSecret" -->
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
        AuthType: "secret",
        ConfigObj: criblcontrolplanesdkgo.Pointer("<value>"),
        ConnectionString: criblcontrolplanesdkgo.Pointer("<value>"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](15000),
        CredsSecrets: criblcontrolplanesdkgo.Pointer("<value>"),
        DatabaseType: components.DatabaseConnectionTypeMysql,
        Description: "Analytics MySQL database",
        ID: "mysql-analytics-db",
        Password: criblcontrolplanesdkgo.Pointer("QdEG1YZma4X6Q0d"),
        RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](8016.38),
        Tags: criblcontrolplanesdkgo.Pointer("analytics,mysql"),
        User: criblcontrolplanesdkgo.Pointer("Deon.Zulauf40"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedDatabaseConnectionConfig != nil {
        // handle response
    }
}
```
### Example Usage: DatabaseConnectionExamplesOracleWithConnectionString

<!-- UsageSnippet language="go" operationID="createDatabaseConnectionConfig" method="post" path="/lib/database-connections" example="DatabaseConnectionExamplesOracleWithConnectionString" -->
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
        ConnectionString: criblcontrolplanesdkgo.Pointer("oracle.example.com:1521/ORCL"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](15000),
        CredsSecrets: criblcontrolplanesdkgo.Pointer("<value>"),
        DatabaseType: components.DatabaseConnectionTypeOracle,
        Description: "Oracle ERP database",
        ID: "oracle-erp",
        Password: criblcontrolplanesdkgo.Pointer("Oracle_Pass456!"),
        RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](8432.44),
        Tags: criblcontrolplanesdkgo.Pointer("erp,oracle,finance"),
        User: criblcontrolplanesdkgo.Pointer("erp_user"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedDatabaseConnectionConfig != nil {
        // handle response
    }
}
```
### Example Usage: DatabaseConnectionExamplesOracleWithCredentialsSecrets

<!-- UsageSnippet language="go" operationID="createDatabaseConnectionConfig" method="post" path="/lib/database-connections" example="DatabaseConnectionExamplesOracleWithCredentialsSecrets" -->
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
        AuthType: "secrets",
        ConfigObj: criblcontrolplanesdkgo.Pointer("<value>"),
        ConnectionString: criblcontrolplanesdkgo.Pointer("<value>"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](15000),
        CredsSecrets: criblcontrolplanesdkgo.Pointer("oracle-secure-credentials"),
        DatabaseType: components.DatabaseConnectionTypeOracle,
        Description: "High-security Oracle database with credential secrets",
        ID: "oracle-secure-db",
        Password: criblcontrolplanesdkgo.Pointer("O0ma5h_gwqFSiXO"),
        RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](5552.27),
        Tags: criblcontrolplanesdkgo.Pointer("secure,oracle,sensitive-data"),
        User: criblcontrolplanesdkgo.Pointer("Alek_Dibbert"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedDatabaseConnectionConfig != nil {
        // handle response
    }
}
```
### Example Usage: DatabaseConnectionExamplesOracleWithSecret

<!-- UsageSnippet language="go" operationID="createDatabaseConnectionConfig" method="post" path="/lib/database-connections" example="DatabaseConnectionExamplesOracleWithSecret" -->
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
        AuthType: "secret",
        ConfigObj: criblcontrolplanesdkgo.Pointer("<value>"),
        ConnectionString: criblcontrolplanesdkgo.Pointer("<value>"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](20000),
        CredsSecrets: criblcontrolplanesdkgo.Pointer("<value>"),
        DatabaseType: components.DatabaseConnectionTypeOracle,
        Description: "Oracle data warehouse",
        ID: "oracle-warehouse",
        Password: criblcontrolplanesdkgo.Pointer("Warehouse_Pass789!"),
        RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](1474.54),
        Tags: criblcontrolplanesdkgo.Pointer("warehouse,oracle,reporting"),
        User: criblcontrolplanesdkgo.Pointer("warehouse_user"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedDatabaseConnectionConfig != nil {
        // handle response
    }
}
```
### Example Usage: DatabaseConnectionExamplesPostgreSQLWithConnectionString

<!-- UsageSnippet language="go" operationID="createDatabaseConnectionConfig" method="post" path="/lib/database-connections" example="DatabaseConnectionExamplesPostgreSQLWithConnectionString" -->
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
        ConnectionString: criblcontrolplanesdkgo.Pointer("postgresql://warehouse_user:SecurePass456@postgres.example.com:5432/warehouse?sslmode=require"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](15000),
        CredsSecrets: criblcontrolplanesdkgo.Pointer("<value>"),
        DatabaseType: components.DatabaseConnectionTypePostgres,
        Description: "Data warehouse PostgreSQL database",
        ID: "postgres-warehouse",
        Password: criblcontrolplanesdkgo.Pointer("iRpyPXc98_8DDCo"),
        RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](6489.43),
        Tags: criblcontrolplanesdkgo.Pointer("warehouse,postgres,reporting"),
        User: criblcontrolplanesdkgo.Pointer("Ashlynn_Cassin"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedDatabaseConnectionConfig != nil {
        // handle response
    }
}
```
### Example Usage: DatabaseConnectionExamplesPostgreSQLWithSecret

<!-- UsageSnippet language="go" operationID="createDatabaseConnectionConfig" method="post" path="/lib/database-connections" example="DatabaseConnectionExamplesPostgreSQLWithSecret" -->
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
        AuthType: "secret",
        ConfigObj: criblcontrolplanesdkgo.Pointer("<value>"),
        ConnectionString: criblcontrolplanesdkgo.Pointer("<value>"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
        CredsSecrets: criblcontrolplanesdkgo.Pointer("<value>"),
        DatabaseType: components.DatabaseConnectionTypePostgres,
        Description: "Logs PostgreSQL database",
        ID: "postgres-logs",
        Password: criblcontrolplanesdkgo.Pointer("S_R5PzDp3_vEzJr"),
        RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](3972.05),
        Tags: criblcontrolplanesdkgo.Pointer("logs,postgres"),
        User: criblcontrolplanesdkgo.Pointer("Eldon53"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedDatabaseConnectionConfig != nil {
        // handle response
    }
}
```
### Example Usage: DatabaseConnectionExamplesSQLServerWithConfigObject

<!-- UsageSnippet language="go" operationID="createDatabaseConnectionConfig" method="post" path="/lib/database-connections" example="DatabaseConnectionExamplesSQLServerWithConfigObject" -->
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
        AuthType: "configObj",
        ConfigObj: criblcontrolplanesdkgo.Pointer("{\"server\":\"sqlserver.example.com\",\"database\":\"Reporting\",\"user\":\"report_user\",\"password\":\"Report_Pass123!\",\"options\":{\"encrypt\":true,\"trustServerCertificate\":false,\"connectTimeout\":20000}}"),
        ConnectionString: criblcontrolplanesdkgo.Pointer("<value>"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](4499.78),
        CredsSecrets: criblcontrolplanesdkgo.Pointer("<value>"),
        DatabaseType: components.DatabaseConnectionTypeSqlserver,
        Description: "Reporting SQL Server database with custom config",
        ID: "sqlserver-reporting",
        Password: criblcontrolplanesdkgo.Pointer("g9Dv7PtDzLx0pmQ"),
        RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](60000),
        Tags: criblcontrolplanesdkgo.Pointer("reporting,sqlserver,analytics"),
        User: criblcontrolplanesdkgo.Pointer("Aubree_Herman32"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedDatabaseConnectionConfig != nil {
        // handle response
    }
}
```
### Example Usage: DatabaseConnectionExamplesSQLServerWithConnectionString

<!-- UsageSnippet language="go" operationID="createDatabaseConnectionConfig" method="post" path="/lib/database-connections" example="DatabaseConnectionExamplesSQLServerWithConnectionString" -->
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
        ConnectionString: criblcontrolplanesdkgo.Pointer("Server=sqlserver.example.com;Database=ERP;User Id=erp_admin;Password=ERP_Pass789!;Encrypt=true"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](15000),
        CredsSecrets: criblcontrolplanesdkgo.Pointer("<value>"),
        DatabaseType: components.DatabaseConnectionTypeSqlserver,
        Description: "ERP SQL Server database",
        ID: "sqlserver-erp",
        Password: criblcontrolplanesdkgo.Pointer("KdnOen49IaSDMY2"),
        RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](30000),
        Tags: criblcontrolplanesdkgo.Pointer("erp,sqlserver,finance"),
        User: criblcontrolplanesdkgo.Pointer("Jessika_Smitham"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedDatabaseConnectionConfig != nil {
        // handle response
    }
}
```
### Example Usage: DatabaseConnectionExamplesSQLServerWithSecret

<!-- UsageSnippet language="go" operationID="createDatabaseConnectionConfig" method="post" path="/lib/database-connections" example="DatabaseConnectionExamplesSQLServerWithSecret" -->
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
        AuthType: "secret",
        ConfigObj: criblcontrolplanesdkgo.Pointer("<value>"),
        ConnectionString: criblcontrolplanesdkgo.Pointer("<value>"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](15000),
        CredsSecrets: criblcontrolplanesdkgo.Pointer("<value>"),
        DatabaseType: components.DatabaseConnectionTypeSqlserver,
        Description: "CRM SQL Server database",
        ID: "sqlserver-crm",
        Password: criblcontrolplanesdkgo.Pointer("Ua8rfhO1uYnrBPC"),
        RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](15000),
        Tags: criblcontrolplanesdkgo.Pointer("crm,sqlserver,sales"),
        User: criblcontrolplanesdkgo.Pointer("Genevieve.Douglas13"),
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

### Example Usage: DatabaseConnectionExamplesMySQLWithConnectionString

<!-- UsageSnippet language="go" operationID="updateDatabaseConnectionConfigById" method="patch" path="/lib/database-connections/{id}" example="DatabaseConnectionExamplesMySQLWithConnectionString" -->
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
### Example Usage: DatabaseConnectionExamplesMySQLWithSecret

<!-- UsageSnippet language="go" operationID="updateDatabaseConnectionConfigById" method="patch" path="/lib/database-connections/{id}" example="DatabaseConnectionExamplesMySQLWithSecret" -->
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
        AuthType: "secret",
        ConfigObj: criblcontrolplanesdkgo.Pointer("<value>"),
        ConnectionString: criblcontrolplanesdkgo.Pointer("<value>"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](15000),
        CredsSecrets: criblcontrolplanesdkgo.Pointer("<value>"),
        DatabaseType: components.DatabaseConnectionTypeMysql,
        Description: "Analytics MySQL database",
        ID: "mysql-analytics-db",
        Password: criblcontrolplanesdkgo.Pointer("I4rBabWGYRpCqjN"),
        RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](3077.93),
        Tags: criblcontrolplanesdkgo.Pointer("analytics,mysql"),
        User: criblcontrolplanesdkgo.Pointer("Gunnar_Hickle-Davis"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedDatabaseConnectionConfig != nil {
        // handle response
    }
}
```
### Example Usage: DatabaseConnectionExamplesOracleWithConnectionString

<!-- UsageSnippet language="go" operationID="updateDatabaseConnectionConfigById" method="patch" path="/lib/database-connections/{id}" example="DatabaseConnectionExamplesOracleWithConnectionString" -->
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
        ConnectionString: criblcontrolplanesdkgo.Pointer("oracle.example.com:1521/ORCL"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](15000),
        CredsSecrets: criblcontrolplanesdkgo.Pointer("<value>"),
        DatabaseType: components.DatabaseConnectionTypeOracle,
        Description: "Oracle ERP database",
        ID: "oracle-erp",
        Password: criblcontrolplanesdkgo.Pointer("Oracle_Pass456!"),
        RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](4812.37),
        Tags: criblcontrolplanesdkgo.Pointer("erp,oracle,finance"),
        User: criblcontrolplanesdkgo.Pointer("erp_user"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedDatabaseConnectionConfig != nil {
        // handle response
    }
}
```
### Example Usage: DatabaseConnectionExamplesOracleWithCredentialsSecrets

<!-- UsageSnippet language="go" operationID="updateDatabaseConnectionConfigById" method="patch" path="/lib/database-connections/{id}" example="DatabaseConnectionExamplesOracleWithCredentialsSecrets" -->
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
        AuthType: "secrets",
        ConfigObj: criblcontrolplanesdkgo.Pointer("<value>"),
        ConnectionString: criblcontrolplanesdkgo.Pointer("<value>"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](15000),
        CredsSecrets: criblcontrolplanesdkgo.Pointer("oracle-secure-credentials"),
        DatabaseType: components.DatabaseConnectionTypeOracle,
        Description: "High-security Oracle database with credential secrets",
        ID: "oracle-secure-db",
        Password: criblcontrolplanesdkgo.Pointer("h2R8xU7uFdKI0S6"),
        RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](9629.77),
        Tags: criblcontrolplanesdkgo.Pointer("secure,oracle,sensitive-data"),
        User: criblcontrolplanesdkgo.Pointer("Kay83"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedDatabaseConnectionConfig != nil {
        // handle response
    }
}
```
### Example Usage: DatabaseConnectionExamplesOracleWithSecret

<!-- UsageSnippet language="go" operationID="updateDatabaseConnectionConfigById" method="patch" path="/lib/database-connections/{id}" example="DatabaseConnectionExamplesOracleWithSecret" -->
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
        AuthType: "secret",
        ConfigObj: criblcontrolplanesdkgo.Pointer("<value>"),
        ConnectionString: criblcontrolplanesdkgo.Pointer("<value>"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](20000),
        CredsSecrets: criblcontrolplanesdkgo.Pointer("<value>"),
        DatabaseType: components.DatabaseConnectionTypeOracle,
        Description: "Oracle data warehouse",
        ID: "oracle-warehouse",
        Password: criblcontrolplanesdkgo.Pointer("Warehouse_Pass789!"),
        RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](2147.18),
        Tags: criblcontrolplanesdkgo.Pointer("warehouse,oracle,reporting"),
        User: criblcontrolplanesdkgo.Pointer("warehouse_user"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedDatabaseConnectionConfig != nil {
        // handle response
    }
}
```
### Example Usage: DatabaseConnectionExamplesPostgreSQLWithConnectionString

<!-- UsageSnippet language="go" operationID="updateDatabaseConnectionConfigById" method="patch" path="/lib/database-connections/{id}" example="DatabaseConnectionExamplesPostgreSQLWithConnectionString" -->
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
        ConnectionString: criblcontrolplanesdkgo.Pointer("postgresql://warehouse_user:SecurePass456@postgres.example.com:5432/warehouse?sslmode=require"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](15000),
        CredsSecrets: criblcontrolplanesdkgo.Pointer("<value>"),
        DatabaseType: components.DatabaseConnectionTypePostgres,
        Description: "Data warehouse PostgreSQL database",
        ID: "postgres-warehouse",
        Password: criblcontrolplanesdkgo.Pointer("SJ185mzGNnaospV"),
        RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](4073.17),
        Tags: criblcontrolplanesdkgo.Pointer("warehouse,postgres,reporting"),
        User: criblcontrolplanesdkgo.Pointer("Charlene27"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedDatabaseConnectionConfig != nil {
        // handle response
    }
}
```
### Example Usage: DatabaseConnectionExamplesPostgreSQLWithSecret

<!-- UsageSnippet language="go" operationID="updateDatabaseConnectionConfigById" method="patch" path="/lib/database-connections/{id}" example="DatabaseConnectionExamplesPostgreSQLWithSecret" -->
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
        AuthType: "secret",
        ConfigObj: criblcontrolplanesdkgo.Pointer("<value>"),
        ConnectionString: criblcontrolplanesdkgo.Pointer("<value>"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
        CredsSecrets: criblcontrolplanesdkgo.Pointer("<value>"),
        DatabaseType: components.DatabaseConnectionTypePostgres,
        Description: "Logs PostgreSQL database",
        ID: "postgres-logs",
        Password: criblcontrolplanesdkgo.Pointer("bJaM8xGlX7QHy5k"),
        RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](145.27),
        Tags: criblcontrolplanesdkgo.Pointer("logs,postgres"),
        User: criblcontrolplanesdkgo.Pointer("Ulices_Morar85"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedDatabaseConnectionConfig != nil {
        // handle response
    }
}
```
### Example Usage: DatabaseConnectionExamplesSQLServerWithConfigObject

<!-- UsageSnippet language="go" operationID="updateDatabaseConnectionConfigById" method="patch" path="/lib/database-connections/{id}" example="DatabaseConnectionExamplesSQLServerWithConfigObject" -->
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
        AuthType: "configObj",
        ConfigObj: criblcontrolplanesdkgo.Pointer("{\"server\":\"sqlserver.example.com\",\"database\":\"Reporting\",\"user\":\"report_user\",\"password\":\"Report_Pass123!\",\"options\":{\"encrypt\":true,\"trustServerCertificate\":false,\"connectTimeout\":20000}}"),
        ConnectionString: criblcontrolplanesdkgo.Pointer("<value>"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](1192.73),
        CredsSecrets: criblcontrolplanesdkgo.Pointer("<value>"),
        DatabaseType: components.DatabaseConnectionTypeSqlserver,
        Description: "Reporting SQL Server database with custom config",
        ID: "sqlserver-reporting",
        Password: criblcontrolplanesdkgo.Pointer("uj_P5dL_oZDsm0P"),
        RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](60000),
        Tags: criblcontrolplanesdkgo.Pointer("reporting,sqlserver,analytics"),
        User: criblcontrolplanesdkgo.Pointer("Retta_Welch"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedDatabaseConnectionConfig != nil {
        // handle response
    }
}
```
### Example Usage: DatabaseConnectionExamplesSQLServerWithConnectionString

<!-- UsageSnippet language="go" operationID="updateDatabaseConnectionConfigById" method="patch" path="/lib/database-connections/{id}" example="DatabaseConnectionExamplesSQLServerWithConnectionString" -->
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
        ConnectionString: criblcontrolplanesdkgo.Pointer("Server=sqlserver.example.com;Database=ERP;User Id=erp_admin;Password=ERP_Pass789!;Encrypt=true"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](15000),
        CredsSecrets: criblcontrolplanesdkgo.Pointer("<value>"),
        DatabaseType: components.DatabaseConnectionTypeSqlserver,
        Description: "ERP SQL Server database",
        ID: "sqlserver-erp",
        Password: criblcontrolplanesdkgo.Pointer("gAICi7pHE7tEZAQ"),
        RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](30000),
        Tags: criblcontrolplanesdkgo.Pointer("erp,sqlserver,finance"),
        User: criblcontrolplanesdkgo.Pointer("Jennifer60"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedDatabaseConnectionConfig != nil {
        // handle response
    }
}
```
### Example Usage: DatabaseConnectionExamplesSQLServerWithSecret

<!-- UsageSnippet language="go" operationID="updateDatabaseConnectionConfigById" method="patch" path="/lib/database-connections/{id}" example="DatabaseConnectionExamplesSQLServerWithSecret" -->
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
        AuthType: "secret",
        ConfigObj: criblcontrolplanesdkgo.Pointer("<value>"),
        ConnectionString: criblcontrolplanesdkgo.Pointer("<value>"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](15000),
        CredsSecrets: criblcontrolplanesdkgo.Pointer("<value>"),
        DatabaseType: components.DatabaseConnectionTypeSqlserver,
        Description: "CRM SQL Server database",
        ID: "sqlserver-crm",
        Password: criblcontrolplanesdkgo.Pointer("uplhN8ZaDm8LRt_"),
        RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](15000),
        Tags: criblcontrolplanesdkgo.Pointer("crm,sqlserver,sales"),
        User: criblcontrolplanesdkgo.Pointer("Millie.Dickens93"),
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