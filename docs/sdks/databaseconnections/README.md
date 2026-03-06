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
        ConnectionString: criblcontrolplanesdkgo.Pointer("mysql://admin:password123@mysql.example.com:3306/production?ssl=true"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
        DatabaseType: components.DatabaseConnectionTypeMysql,
        Description: "Production MySQL database for customer data",
        ID: "mysql-prod-db",
        Tags: criblcontrolplanesdkgo.Pointer("production,mysql,customer-data"),
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
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](15000),
        DatabaseType: components.DatabaseConnectionTypeMysql,
        Description: "Analytics MySQL database",
        ID: "mysql-analytics-db",
        Tags: criblcontrolplanesdkgo.Pointer("analytics,mysql"),
        TextSecret: criblcontrolplanesdkgo.Pointer("mysql-analytics-connection"),
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
        ConnectionString: criblcontrolplanesdkgo.Pointer("oracle.example.com:1521/ORCL"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](15000),
        DatabaseType: components.DatabaseConnectionTypeOracle,
        Description: "Oracle ERP database",
        ID: "oracle-erp",
        Password: criblcontrolplanesdkgo.Pointer("Oracle_Pass456!"),
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
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](15000),
        CredsSecrets: criblcontrolplanesdkgo.Pointer("oracle-secure-credentials"),
        DatabaseType: components.DatabaseConnectionTypeOracle,
        Description: "High-security Oracle database with credential secrets",
        ID: "oracle-secure-db",
        Tags: criblcontrolplanesdkgo.Pointer("secure,oracle,sensitive-data"),
        TextSecret: criblcontrolplanesdkgo.Pointer("oracle-secure-connection"),
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
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](20000),
        DatabaseType: components.DatabaseConnectionTypeOracle,
        Description: "Oracle data warehouse",
        ID: "oracle-warehouse",
        Password: criblcontrolplanesdkgo.Pointer("Warehouse_Pass789!"),
        Tags: criblcontrolplanesdkgo.Pointer("warehouse,oracle,reporting"),
        TextSecret: criblcontrolplanesdkgo.Pointer("oracle-warehouse-connection"),
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
        ConnectionString: criblcontrolplanesdkgo.Pointer("postgresql://warehouse_user:SecurePass456@postgres.example.com:5432/warehouse?sslmode=require"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](15000),
        DatabaseType: components.DatabaseConnectionTypePostgres,
        Description: "Data warehouse PostgreSQL database",
        ID: "postgres-warehouse",
        Tags: criblcontrolplanesdkgo.Pointer("warehouse,postgres,reporting"),
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
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
        DatabaseType: components.DatabaseConnectionTypePostgres,
        Description: "Logs PostgreSQL database",
        ID: "postgres-logs",
        Tags: criblcontrolplanesdkgo.Pointer("logs,postgres"),
        TextSecret: criblcontrolplanesdkgo.Pointer("postgres-logs-connection"),
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
        DatabaseType: components.DatabaseConnectionTypeSqlserver,
        Description: "Reporting SQL Server database with custom config",
        ID: "sqlserver-reporting",
        RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](60000),
        Tags: criblcontrolplanesdkgo.Pointer("reporting,sqlserver,analytics"),
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
        ConnectionString: criblcontrolplanesdkgo.Pointer("Server=sqlserver.example.com;Database=ERP;User Id=erp_admin;Password=ERP_Pass789!;Encrypt=true"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](15000),
        DatabaseType: components.DatabaseConnectionTypeSqlserver,
        Description: "ERP SQL Server database",
        ID: "sqlserver-erp",
        RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](30000),
        Tags: criblcontrolplanesdkgo.Pointer("erp,sqlserver,finance"),
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
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](15000),
        DatabaseType: components.DatabaseConnectionTypeSqlserver,
        Description: "CRM SQL Server database",
        ID: "sqlserver-crm",
        RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](15000),
        Tags: criblcontrolplanesdkgo.Pointer("crm,sqlserver,sales"),
        TextSecret: criblcontrolplanesdkgo.Pointer("sqlserver-crm-connection"),
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
        ConnectionString: criblcontrolplanesdkgo.Pointer("mysql://admin:password123@mysql.example.com:3306/production?ssl=true"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
        DatabaseType: components.DatabaseConnectionTypeMysql,
        Description: "Production MySQL database for customer data",
        ID: "mysql-prod-db",
        Tags: criblcontrolplanesdkgo.Pointer("production,mysql,customer-data"),
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
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](15000),
        DatabaseType: components.DatabaseConnectionTypeMysql,
        Description: "Analytics MySQL database",
        ID: "mysql-analytics-db",
        Tags: criblcontrolplanesdkgo.Pointer("analytics,mysql"),
        TextSecret: criblcontrolplanesdkgo.Pointer("mysql-analytics-connection"),
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
        ConnectionString: criblcontrolplanesdkgo.Pointer("oracle.example.com:1521/ORCL"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](15000),
        DatabaseType: components.DatabaseConnectionTypeOracle,
        Description: "Oracle ERP database",
        ID: "oracle-erp",
        Password: criblcontrolplanesdkgo.Pointer("Oracle_Pass456!"),
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
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](15000),
        CredsSecrets: criblcontrolplanesdkgo.Pointer("oracle-secure-credentials"),
        DatabaseType: components.DatabaseConnectionTypeOracle,
        Description: "High-security Oracle database with credential secrets",
        ID: "oracle-secure-db",
        Tags: criblcontrolplanesdkgo.Pointer("secure,oracle,sensitive-data"),
        TextSecret: criblcontrolplanesdkgo.Pointer("oracle-secure-connection"),
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
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](20000),
        DatabaseType: components.DatabaseConnectionTypeOracle,
        Description: "Oracle data warehouse",
        ID: "oracle-warehouse",
        Password: criblcontrolplanesdkgo.Pointer("Warehouse_Pass789!"),
        Tags: criblcontrolplanesdkgo.Pointer("warehouse,oracle,reporting"),
        TextSecret: criblcontrolplanesdkgo.Pointer("oracle-warehouse-connection"),
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
        ConnectionString: criblcontrolplanesdkgo.Pointer("postgresql://warehouse_user:SecurePass456@postgres.example.com:5432/warehouse?sslmode=require"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](15000),
        DatabaseType: components.DatabaseConnectionTypePostgres,
        Description: "Data warehouse PostgreSQL database",
        ID: "postgres-warehouse",
        Tags: criblcontrolplanesdkgo.Pointer("warehouse,postgres,reporting"),
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
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
        DatabaseType: components.DatabaseConnectionTypePostgres,
        Description: "Logs PostgreSQL database",
        ID: "postgres-logs",
        Tags: criblcontrolplanesdkgo.Pointer("logs,postgres"),
        TextSecret: criblcontrolplanesdkgo.Pointer("postgres-logs-connection"),
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
        DatabaseType: components.DatabaseConnectionTypeSqlserver,
        Description: "Reporting SQL Server database with custom config",
        ID: "sqlserver-reporting",
        RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](60000),
        Tags: criblcontrolplanesdkgo.Pointer("reporting,sqlserver,analytics"),
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
        ConnectionString: criblcontrolplanesdkgo.Pointer("Server=sqlserver.example.com;Database=ERP;User Id=erp_admin;Password=ERP_Pass789!;Encrypt=true"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](15000),
        DatabaseType: components.DatabaseConnectionTypeSqlserver,
        Description: "ERP SQL Server database",
        ID: "sqlserver-erp",
        RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](30000),
        Tags: criblcontrolplanesdkgo.Pointer("erp,sqlserver,finance"),
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
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](15000),
        DatabaseType: components.DatabaseConnectionTypeSqlserver,
        Description: "CRM SQL Server database",
        ID: "sqlserver-crm",
        RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](15000),
        Tags: criblcontrolplanesdkgo.Pointer("crm,sqlserver,sales"),
        TextSecret: criblcontrolplanesdkgo.Pointer("sqlserver-crm-connection"),
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