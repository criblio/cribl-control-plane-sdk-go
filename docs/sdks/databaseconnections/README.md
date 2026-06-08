# DatabaseConnections

## Overview

Actions related to DatabaseConnections

### Available Operations

* [List](#list) - List all Database Connections
* [Create](#create) - Create a Database Connection
* [Get](#get) - Get a Database Connection
* [Update](#update) - Update a Database Connection
* [Delete](#delete) - Delete a Database Connection

## List

Get a list of all Database Connections.

### Example Usage

<!-- UsageSnippet language="go" operationID="getDatabaseConnectionConfig" method="get" path="/lib/database-connections" example="DatabaseConnectionListResponseExamplesDatabaseConnectionList" -->
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

    res, err := s.DatabaseConnections.List(ctx, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DatabaseConnectionResponseEnvelope != nil {
        for {
            // handle items

            res, err = res.Next()

            if err != nil {
                // handle error
            }

            if res == nil {
                break
            }
        }
    }
}
```

### Parameters

| Parameter                                                                                                                                                         | Type                                                                                                                                                              | Required                                                                                                                                                          | Description                                                                                                                                                       |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                                                                                             | :heavy_check_mark:                                                                                                                                                | The context to use for the request.                                                                                                                               |
| `databaseType`                                                                                                                                                    | [*components.DatabaseConnectionType](../../models/components/databaseconnectiontype.md)                                                                           | :heavy_minus_sign:                                                                                                                                                | Filter results by database engine type. Use this parameter to return only Database Connections for the specified engine.                                          |
| `limit`                                                                                                                                                           | `*int64`                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                | Maximum number of Database Connections to return in the response for this request. Use with <code>offset</code> to paginate the response into manageable batches. |
| `offset`                                                                                                                                                          | `*int64`                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                | Starting point from which to retrieve results for this request. Use with <code>limit</code> to paginate the response into manageable batches.                     |
| `opts`                                                                                                                                                            | [][operations.Option](../../models/operations/option.md)                                                                                                          | :heavy_minus_sign:                                                                                                                                                | The options for this request.                                                                                                                                     |

### Response

**[*operations.GetDatabaseConnectionConfigResponse](../../models/operations/getdatabaseconnectionconfigresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Create

Create a new Database Connection.

### Example Usage: DatabaseConnectionBadRequestResponseExamplesInvalidDatabaseConnectionRequest

<!-- UsageSnippet language="go" operationID="createDatabaseConnectionConfig" method="post" path="/lib/database-connections" example="DatabaseConnectionBadRequestResponseExamplesInvalidDatabaseConnectionRequest" -->
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
        AuthType: components.DatabaseConnectionAuthTypeConnectionString,
        ConfigObj: criblcontrolplanesdkgo.Pointer("{\"server\":\"sqlserver.example.com\",\"database\":\"Reporting\",\"user\":\"yourUsername\",\"password\":\"yourPassword\",\"options\":{\"trustServerCertificate\":false,\"connectTimeout\":20000}}"),
        ConnectionString: criblcontrolplanesdkgo.Pointer("mysql://yourUsername:yourPassword@mysql.example.com:3306/production?ssl=true"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[int64](10000),
        CredsSecrets: criblcontrolplanesdkgo.Pointer("oracle-production-credentials"),
        DatabaseType: components.DatabaseConnectionTypePostgres,
        Description: "Production MySQL database for customer data",
        ID: "mysql-prod-db",
        Password: criblcontrolplanesdkgo.Pointer("yourPassword"),
        RequestTimeout: criblcontrolplanesdkgo.Pointer[int64](30000),
        Tags: criblcontrolplanesdkgo.Pointer("production,mysql,customer-data"),
        TextSecret: criblcontrolplanesdkgo.Pointer("mysql-production-connection"),
        User: criblcontrolplanesdkgo.Pointer("yourUsername"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatabaseConnectionResponseEnvelope != nil {
        // handle response
    }
}
```
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
        AuthType: components.DatabaseConnectionAuthTypeConnectionString,
        ConnectionString: criblcontrolplanesdkgo.Pointer("mysql://yourUsername:yourPassword@mysql.example.com:3306/production?ssl=true"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[int64](10000),
        DatabaseType: components.DatabaseConnectionTypeMysql,
        Description: "Production MySQL database for customer data",
        ID: "mysql-prod-db",
        Tags: criblcontrolplanesdkgo.Pointer("production,mysql,customer-data"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatabaseConnectionResponseEnvelope != nil {
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
        AuthType: components.DatabaseConnectionAuthTypeSecret,
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[int64](15000),
        DatabaseType: components.DatabaseConnectionTypeMysql,
        Description: "Analytics MySQL database",
        ID: "mysql-analytics-db",
        Tags: criblcontrolplanesdkgo.Pointer("analytics,mysql"),
        TextSecret: criblcontrolplanesdkgo.Pointer("mysql-analytics-connection"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatabaseConnectionResponseEnvelope != nil {
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
        AuthType: components.DatabaseConnectionAuthTypeConnectionString,
        ConnectionString: criblcontrolplanesdkgo.Pointer("oracle.example.com:1521/ORCL"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[int64](15000),
        DatabaseType: components.DatabaseConnectionTypeOracle,
        Description: "Oracle ERP database",
        ID: "oracle-erp",
        Password: criblcontrolplanesdkgo.Pointer("yourPassword"),
        Tags: criblcontrolplanesdkgo.Pointer("erp,oracle,finance"),
        User: criblcontrolplanesdkgo.Pointer("yourUsername"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatabaseConnectionResponseEnvelope != nil {
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
        AuthType: components.DatabaseConnectionAuthTypeSecrets,
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[int64](15000),
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
    if res.DatabaseConnectionResponseEnvelope != nil {
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
        AuthType: components.DatabaseConnectionAuthTypeSecret,
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[int64](20000),
        DatabaseType: components.DatabaseConnectionTypeOracle,
        Description: "Oracle data warehouse",
        ID: "oracle-warehouse",
        Password: criblcontrolplanesdkgo.Pointer("yourPassword"),
        Tags: criblcontrolplanesdkgo.Pointer("warehouse,oracle,reporting"),
        TextSecret: criblcontrolplanesdkgo.Pointer("oracle-warehouse-connection"),
        User: criblcontrolplanesdkgo.Pointer("yourUsername"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatabaseConnectionResponseEnvelope != nil {
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
        AuthType: components.DatabaseConnectionAuthTypeConnectionString,
        ConnectionString: criblcontrolplanesdkgo.Pointer("postgresql://yourUsername:yourPassword@postgres.example.com:5432/warehouse?sslmode=require"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[int64](15000),
        DatabaseType: components.DatabaseConnectionTypePostgres,
        Description: "Data warehouse PostgreSQL database",
        ID: "postgres-warehouse",
        Tags: criblcontrolplanesdkgo.Pointer("warehouse,postgres,reporting"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatabaseConnectionResponseEnvelope != nil {
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
        AuthType: components.DatabaseConnectionAuthTypeSecret,
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[int64](10000),
        DatabaseType: components.DatabaseConnectionTypePostgres,
        Description: "Logs PostgreSQL database",
        ID: "postgres-logs",
        Tags: criblcontrolplanesdkgo.Pointer("logs,postgres"),
        TextSecret: criblcontrolplanesdkgo.Pointer("postgres-logs-connection"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatabaseConnectionResponseEnvelope != nil {
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
        AuthType: components.DatabaseConnectionAuthTypeConfigObj,
        ConfigObj: criblcontrolplanesdkgo.Pointer("{\"server\":\"sqlserver.example.com\",\"database\":\"Reporting\",\"user\":\"yourUsername\",\"password\":\"yourPassword\",\"options\":{\"encrypt\":true,\"trustServerCertificate\":false,\"connectTimeout\":20000}}"),
        DatabaseType: components.DatabaseConnectionTypeSqlserver,
        Description: "Reporting SQL Server database with custom config",
        ID: "sqlserver-reporting",
        RequestTimeout: criblcontrolplanesdkgo.Pointer[int64](60000),
        Tags: criblcontrolplanesdkgo.Pointer("reporting,sqlserver,analytics"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatabaseConnectionResponseEnvelope != nil {
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
        AuthType: components.DatabaseConnectionAuthTypeConnectionString,
        ConnectionString: criblcontrolplanesdkgo.Pointer("Server=sqlserver.example.com;Database=ERP;User Id=yourUsername;Password=yourPassword;Encrypt=true"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[int64](15000),
        DatabaseType: components.DatabaseConnectionTypeSqlserver,
        Description: "ERP SQL Server database",
        ID: "sqlserver-erp",
        RequestTimeout: criblcontrolplanesdkgo.Pointer[int64](30000),
        Tags: criblcontrolplanesdkgo.Pointer("erp,sqlserver,finance"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatabaseConnectionResponseEnvelope != nil {
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
        AuthType: components.DatabaseConnectionAuthTypeSecret,
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[int64](15000),
        DatabaseType: components.DatabaseConnectionTypeSqlserver,
        Description: "CRM SQL Server database",
        ID: "sqlserver-crm",
        RequestTimeout: criblcontrolplanesdkgo.Pointer[int64](15000),
        Tags: criblcontrolplanesdkgo.Pointer("crm,sqlserver,sales"),
        TextSecret: criblcontrolplanesdkgo.Pointer("sqlserver-crm-connection"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatabaseConnectionResponseEnvelope != nil {
        // handle response
    }
}
```
### Example Usage: DatabaseConnectionResponseExamplesMySQLDatabaseConnection

<!-- UsageSnippet language="go" operationID="createDatabaseConnectionConfig" method="post" path="/lib/database-connections" example="DatabaseConnectionResponseExamplesMySQLDatabaseConnection" -->
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
        AuthType: components.DatabaseConnectionAuthTypeConnectionString,
        ConfigObj: criblcontrolplanesdkgo.Pointer("{\"server\":\"sqlserver.example.com\",\"database\":\"Reporting\",\"user\":\"yourUsername\",\"password\":\"yourPassword\",\"options\":{\"trustServerCertificate\":false,\"connectTimeout\":20000}}"),
        ConnectionString: criblcontrolplanesdkgo.Pointer("mysql://yourUsername:yourPassword@mysql.example.com:3306/production?ssl=true"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[int64](10000),
        CredsSecrets: criblcontrolplanesdkgo.Pointer("oracle-production-credentials"),
        DatabaseType: components.DatabaseConnectionTypePostgres,
        Description: "Production MySQL database for customer data",
        ID: "mysql-prod-db",
        Password: criblcontrolplanesdkgo.Pointer("yourPassword"),
        RequestTimeout: criblcontrolplanesdkgo.Pointer[int64](30000),
        Tags: criblcontrolplanesdkgo.Pointer("production,mysql,customer-data"),
        TextSecret: criblcontrolplanesdkgo.Pointer("mysql-production-connection"),
        User: criblcontrolplanesdkgo.Pointer("yourUsername"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatabaseConnectionResponseEnvelope != nil {
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

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| apierrors.RestAPIJSONError | 400                        | application/json           |
| apierrors.Error            | 500                        | application/json           |
| apierrors.APIError         | 4XX, 5XX                   | \*/\*                      |

## Get

Get the specified Database Connection.

### Example Usage

<!-- UsageSnippet language="go" operationID="getDatabaseConnectionConfigById" method="get" path="/lib/database-connections/{id}" example="DatabaseConnectionResponseExamplesMySQLDatabaseConnection" -->
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
    if res.DatabaseConnectionResponseEnvelope != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | The <code>id</code> of the Database Connection to get.   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetDatabaseConnectionConfigByIDResponse](../../models/operations/getdatabaseconnectionconfigbyidresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| apierrors.RestAPIJSONError | 404                        | application/json           |
| apierrors.Error            | 500                        | application/json           |
| apierrors.APIError         | 4XX, 5XX                   | \*/\*                      |

## Update

Update the specified Database Connection.<br/><br/>Provide a complete representation of the Database Connection that you want to update in the request body. This endpoint does not support partial updates. Cribl removes any omitted fields when updating the Database Connection.<br/><br/>Confirm that the configuration in your request body is correct before sending the request. If the configuration is incorrect, the updated Database Connection might not function as expected.

### Example Usage: DatabaseConnectionBadRequestResponseExamplesInvalidDatabaseConnectionRequest

<!-- UsageSnippet language="go" operationID="updateDatabaseConnectionConfigById" method="patch" path="/lib/database-connections/{id}" example="DatabaseConnectionBadRequestResponseExamplesInvalidDatabaseConnectionRequest" -->
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
        AuthType: components.DatabaseConnectionAuthTypeConfigObj,
        ConfigObj: criblcontrolplanesdkgo.Pointer("{\"server\":\"sqlserver.example.com\",\"database\":\"Reporting\",\"user\":\"yourUsername\",\"password\":\"yourPassword\",\"options\":{\"trustServerCertificate\":false,\"connectTimeout\":20000}}"),
        ConnectionString: criblcontrolplanesdkgo.Pointer("mysql://yourUsername:yourPassword@mysql.example.com:3306/production?ssl=true"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[int64](10000),
        CredsSecrets: criblcontrolplanesdkgo.Pointer("oracle-production-credentials"),
        DatabaseType: components.DatabaseConnectionTypeSqlserver,
        Description: "Production MySQL database for customer data",
        ID: "mysql-prod-db",
        Password: criblcontrolplanesdkgo.Pointer("yourPassword"),
        RequestTimeout: criblcontrolplanesdkgo.Pointer[int64](30000),
        Tags: criblcontrolplanesdkgo.Pointer("production,mysql,customer-data"),
        TextSecret: criblcontrolplanesdkgo.Pointer("mysql-production-connection"),
        User: criblcontrolplanesdkgo.Pointer("yourUsername"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatabaseConnectionResponseEnvelope != nil {
        // handle response
    }
}
```
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
        AuthType: components.DatabaseConnectionAuthTypeConnectionString,
        ConnectionString: criblcontrolplanesdkgo.Pointer("mysql://admin:password123@mysql.example.com:3306/production?ssl=true"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[int64](10000),
        DatabaseType: components.DatabaseConnectionTypeMysql,
        Description: "Production MySQL database for customer data",
        ID: "mysql-prod-db",
        Tags: criblcontrolplanesdkgo.Pointer("production,mysql,customer-data"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatabaseConnectionResponseEnvelope != nil {
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
        AuthType: components.DatabaseConnectionAuthTypeSecret,
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[int64](15000),
        DatabaseType: components.DatabaseConnectionTypeMysql,
        Description: "Analytics MySQL database",
        ID: "mysql-analytics-db",
        Tags: criblcontrolplanesdkgo.Pointer("analytics,mysql"),
        TextSecret: criblcontrolplanesdkgo.Pointer("mysql-analytics-connection"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatabaseConnectionResponseEnvelope != nil {
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
        AuthType: components.DatabaseConnectionAuthTypeConnectionString,
        ConnectionString: criblcontrolplanesdkgo.Pointer("oracle.example.com:1521/ORCL"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[int64](15000),
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
    if res.DatabaseConnectionResponseEnvelope != nil {
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
        AuthType: components.DatabaseConnectionAuthTypeSecrets,
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[int64](15000),
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
    if res.DatabaseConnectionResponseEnvelope != nil {
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
        AuthType: components.DatabaseConnectionAuthTypeSecret,
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[int64](20000),
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
    if res.DatabaseConnectionResponseEnvelope != nil {
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
        AuthType: components.DatabaseConnectionAuthTypeConnectionString,
        ConnectionString: criblcontrolplanesdkgo.Pointer("postgresql://warehouse_user:SecurePass456@postgres.example.com:5432/warehouse?sslmode=require"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[int64](15000),
        DatabaseType: components.DatabaseConnectionTypePostgres,
        Description: "Data warehouse PostgreSQL database",
        ID: "postgres-warehouse",
        Tags: criblcontrolplanesdkgo.Pointer("warehouse,postgres,reporting"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatabaseConnectionResponseEnvelope != nil {
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
        AuthType: components.DatabaseConnectionAuthTypeSecret,
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[int64](10000),
        DatabaseType: components.DatabaseConnectionTypePostgres,
        Description: "Logs PostgreSQL database",
        ID: "postgres-logs",
        Tags: criblcontrolplanesdkgo.Pointer("logs,postgres"),
        TextSecret: criblcontrolplanesdkgo.Pointer("postgres-logs-connection"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatabaseConnectionResponseEnvelope != nil {
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
        AuthType: components.DatabaseConnectionAuthTypeConfigObj,
        ConfigObj: criblcontrolplanesdkgo.Pointer("{\"server\":\"sqlserver.example.com\",\"database\":\"Reporting\",\"user\":\"report_user\",\"password\":\"Report_Pass123!\",\"options\":{\"encrypt\":true,\"trustServerCertificate\":false,\"connectTimeout\":20000}}"),
        DatabaseType: components.DatabaseConnectionTypeSqlserver,
        Description: "Reporting SQL Server database with custom config",
        ID: "sqlserver-reporting",
        RequestTimeout: criblcontrolplanesdkgo.Pointer[int64](60000),
        Tags: criblcontrolplanesdkgo.Pointer("reporting,sqlserver,analytics"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatabaseConnectionResponseEnvelope != nil {
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
        AuthType: components.DatabaseConnectionAuthTypeConnectionString,
        ConnectionString: criblcontrolplanesdkgo.Pointer("Server=sqlserver.example.com;Database=ERP;User Id=erp_admin;Password=ERP_Pass789!;Encrypt=true"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[int64](15000),
        DatabaseType: components.DatabaseConnectionTypeSqlserver,
        Description: "ERP SQL Server database",
        ID: "sqlserver-erp",
        RequestTimeout: criblcontrolplanesdkgo.Pointer[int64](30000),
        Tags: criblcontrolplanesdkgo.Pointer("erp,sqlserver,finance"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatabaseConnectionResponseEnvelope != nil {
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
        AuthType: components.DatabaseConnectionAuthTypeSecret,
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[int64](15000),
        DatabaseType: components.DatabaseConnectionTypeSqlserver,
        Description: "CRM SQL Server database",
        ID: "sqlserver-crm",
        RequestTimeout: criblcontrolplanesdkgo.Pointer[int64](15000),
        Tags: criblcontrolplanesdkgo.Pointer("crm,sqlserver,sales"),
        TextSecret: criblcontrolplanesdkgo.Pointer("sqlserver-crm-connection"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatabaseConnectionResponseEnvelope != nil {
        // handle response
    }
}
```
### Example Usage: DatabaseConnectionResponseExamplesMySQLDatabaseConnection

<!-- UsageSnippet language="go" operationID="updateDatabaseConnectionConfigById" method="patch" path="/lib/database-connections/{id}" example="DatabaseConnectionResponseExamplesMySQLDatabaseConnection" -->
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
        AuthType: components.DatabaseConnectionAuthTypeConfigObj,
        ConfigObj: criblcontrolplanesdkgo.Pointer("{\"server\":\"sqlserver.example.com\",\"database\":\"Reporting\",\"user\":\"yourUsername\",\"password\":\"yourPassword\",\"options\":{\"trustServerCertificate\":false,\"connectTimeout\":20000}}"),
        ConnectionString: criblcontrolplanesdkgo.Pointer("mysql://yourUsername:yourPassword@mysql.example.com:3306/production?ssl=true"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[int64](10000),
        CredsSecrets: criblcontrolplanesdkgo.Pointer("oracle-production-credentials"),
        DatabaseType: components.DatabaseConnectionTypeSqlserver,
        Description: "Production MySQL database for customer data",
        ID: "mysql-prod-db",
        Password: criblcontrolplanesdkgo.Pointer("yourPassword"),
        RequestTimeout: criblcontrolplanesdkgo.Pointer[int64](30000),
        Tags: criblcontrolplanesdkgo.Pointer("production,mysql,customer-data"),
        TextSecret: criblcontrolplanesdkgo.Pointer("mysql-production-connection"),
        User: criblcontrolplanesdkgo.Pointer("yourUsername"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatabaseConnectionResponseEnvelope != nil {
        // handle response
    }
}
```
### Example Usage: UpdateDatabaseConnectionExamplesMySQLWithConnectionString

<!-- UsageSnippet language="go" operationID="updateDatabaseConnectionConfigById" method="patch" path="/lib/database-connections/{id}" example="UpdateDatabaseConnectionExamplesMySQLWithConnectionString" -->
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
        AuthType: components.DatabaseConnectionAuthTypeConnectionString,
        ConnectionString: criblcontrolplanesdkgo.Pointer("mysql://admin:password123@mysql.example.com:3306/production?ssl=true"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[int64](10000),
        DatabaseType: components.DatabaseConnectionTypeMysql,
        Description: "Production MySQL database for customer data",
        ID: "mysql-prod-db",
        Tags: criblcontrolplanesdkgo.Pointer("production,mysql,customer-data"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatabaseConnectionResponseEnvelope != nil {
        // handle response
    }
}
```
### Example Usage: UpdateDatabaseConnectionExamplesMySQLWithSecret

<!-- UsageSnippet language="go" operationID="updateDatabaseConnectionConfigById" method="patch" path="/lib/database-connections/{id}" example="UpdateDatabaseConnectionExamplesMySQLWithSecret" -->
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
        AuthType: components.DatabaseConnectionAuthTypeSecret,
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[int64](15000),
        DatabaseType: components.DatabaseConnectionTypeMysql,
        Description: "Analytics MySQL database",
        ID: "mysql-analytics-db",
        Tags: criblcontrolplanesdkgo.Pointer("analytics,mysql"),
        TextSecret: criblcontrolplanesdkgo.Pointer("mysql-analytics-connection"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatabaseConnectionResponseEnvelope != nil {
        // handle response
    }
}
```
### Example Usage: UpdateDatabaseConnectionExamplesOracleWithConnectionString

<!-- UsageSnippet language="go" operationID="updateDatabaseConnectionConfigById" method="patch" path="/lib/database-connections/{id}" example="UpdateDatabaseConnectionExamplesOracleWithConnectionString" -->
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
        AuthType: components.DatabaseConnectionAuthTypeConnectionString,
        ConnectionString: criblcontrolplanesdkgo.Pointer("oracle.example.com:1521/ORCL"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[int64](15000),
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
    if res.DatabaseConnectionResponseEnvelope != nil {
        // handle response
    }
}
```
### Example Usage: UpdateDatabaseConnectionExamplesOracleWithCredentialsSecrets

<!-- UsageSnippet language="go" operationID="updateDatabaseConnectionConfigById" method="patch" path="/lib/database-connections/{id}" example="UpdateDatabaseConnectionExamplesOracleWithCredentialsSecrets" -->
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
        AuthType: components.DatabaseConnectionAuthTypeSecrets,
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[int64](15000),
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
    if res.DatabaseConnectionResponseEnvelope != nil {
        // handle response
    }
}
```
### Example Usage: UpdateDatabaseConnectionExamplesOracleWithSecret

<!-- UsageSnippet language="go" operationID="updateDatabaseConnectionConfigById" method="patch" path="/lib/database-connections/{id}" example="UpdateDatabaseConnectionExamplesOracleWithSecret" -->
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
        AuthType: components.DatabaseConnectionAuthTypeSecret,
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[int64](20000),
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
    if res.DatabaseConnectionResponseEnvelope != nil {
        // handle response
    }
}
```
### Example Usage: UpdateDatabaseConnectionExamplesPostgreSQLWithConnectionString

<!-- UsageSnippet language="go" operationID="updateDatabaseConnectionConfigById" method="patch" path="/lib/database-connections/{id}" example="UpdateDatabaseConnectionExamplesPostgreSQLWithConnectionString" -->
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
        AuthType: components.DatabaseConnectionAuthTypeConnectionString,
        ConnectionString: criblcontrolplanesdkgo.Pointer("postgresql://warehouse_user:SecurePass456@postgres.example.com:5432/warehouse?sslmode=require"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[int64](15000),
        DatabaseType: components.DatabaseConnectionTypePostgres,
        Description: "Data warehouse PostgreSQL database",
        ID: "postgres-warehouse",
        Tags: criblcontrolplanesdkgo.Pointer("warehouse,postgres,reporting"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatabaseConnectionResponseEnvelope != nil {
        // handle response
    }
}
```
### Example Usage: UpdateDatabaseConnectionExamplesPostgreSQLWithSecret

<!-- UsageSnippet language="go" operationID="updateDatabaseConnectionConfigById" method="patch" path="/lib/database-connections/{id}" example="UpdateDatabaseConnectionExamplesPostgreSQLWithSecret" -->
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
        AuthType: components.DatabaseConnectionAuthTypeSecret,
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[int64](10000),
        DatabaseType: components.DatabaseConnectionTypePostgres,
        Description: "Logs PostgreSQL database",
        ID: "postgres-logs",
        Tags: criblcontrolplanesdkgo.Pointer("logs,postgres"),
        TextSecret: criblcontrolplanesdkgo.Pointer("postgres-logs-connection"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatabaseConnectionResponseEnvelope != nil {
        // handle response
    }
}
```
### Example Usage: UpdateDatabaseConnectionExamplesSQLServerWithConfigObject

<!-- UsageSnippet language="go" operationID="updateDatabaseConnectionConfigById" method="patch" path="/lib/database-connections/{id}" example="UpdateDatabaseConnectionExamplesSQLServerWithConfigObject" -->
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
        AuthType: components.DatabaseConnectionAuthTypeConfigObj,
        ConfigObj: criblcontrolplanesdkgo.Pointer("{\"server\":\"sqlserver.example.com\",\"database\":\"Reporting\",\"user\":\"report_user\",\"password\":\"Report_Pass123!\",\"options\":{\"encrypt\":true,\"trustServerCertificate\":false,\"connectTimeout\":20000}}"),
        DatabaseType: components.DatabaseConnectionTypeSqlserver,
        Description: "Reporting SQL Server database with custom config",
        ID: "sqlserver-reporting",
        RequestTimeout: criblcontrolplanesdkgo.Pointer[int64](60000),
        Tags: criblcontrolplanesdkgo.Pointer("reporting,sqlserver,analytics"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatabaseConnectionResponseEnvelope != nil {
        // handle response
    }
}
```
### Example Usage: UpdateDatabaseConnectionExamplesSQLServerWithConnectionString

<!-- UsageSnippet language="go" operationID="updateDatabaseConnectionConfigById" method="patch" path="/lib/database-connections/{id}" example="UpdateDatabaseConnectionExamplesSQLServerWithConnectionString" -->
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
        AuthType: components.DatabaseConnectionAuthTypeConnectionString,
        ConnectionString: criblcontrolplanesdkgo.Pointer("Server=sqlserver.example.com;Database=ERP;User Id=erp_admin;Password=ERP_Pass789!;Encrypt=true"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[int64](15000),
        DatabaseType: components.DatabaseConnectionTypeSqlserver,
        Description: "ERP SQL Server database",
        ID: "sqlserver-erp",
        RequestTimeout: criblcontrolplanesdkgo.Pointer[int64](30000),
        Tags: criblcontrolplanesdkgo.Pointer("erp,sqlserver,finance"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatabaseConnectionResponseEnvelope != nil {
        // handle response
    }
}
```
### Example Usage: UpdateDatabaseConnectionExamplesSQLServerWithSecret

<!-- UsageSnippet language="go" operationID="updateDatabaseConnectionConfigById" method="patch" path="/lib/database-connections/{id}" example="UpdateDatabaseConnectionExamplesSQLServerWithSecret" -->
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
        AuthType: components.DatabaseConnectionAuthTypeSecret,
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[int64](15000),
        DatabaseType: components.DatabaseConnectionTypeSqlserver,
        Description: "CRM SQL Server database",
        ID: "sqlserver-crm",
        RequestTimeout: criblcontrolplanesdkgo.Pointer[int64](15000),
        Tags: criblcontrolplanesdkgo.Pointer("crm,sqlserver,sales"),
        TextSecret: criblcontrolplanesdkgo.Pointer("sqlserver-crm-connection"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatabaseConnectionResponseEnvelope != nil {
        // handle response
    }
}
```
### Example Usage: UpdateDatabaseConnectionExamplesUpdateMySQLDatabaseConnectionWithConnectionString

<!-- UsageSnippet language="go" operationID="updateDatabaseConnectionConfigById" method="patch" path="/lib/database-connections/{id}" example="UpdateDatabaseConnectionExamplesUpdateMySQLDatabaseConnectionWithConnectionString" -->
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
        AuthType: components.DatabaseConnectionAuthTypeConnectionString,
        ConnectionString: criblcontrolplanesdkgo.Pointer("mysql://yourUsername:yourPassword@mysql.example.com:3306/production?ssl=true"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[int64](10000),
        DatabaseType: components.DatabaseConnectionTypeMysql,
        Description: "Production MySQL database for customer data",
        ID: "mysql-prod-db",
        Tags: criblcontrolplanesdkgo.Pointer("production,mysql,customer-data"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatabaseConnectionResponseEnvelope != nil {
        // handle response
    }
}
```
### Example Usage: UpdateDatabaseConnectionExamplesUpdateMySQLDatabaseConnectionWithSecret

<!-- UsageSnippet language="go" operationID="updateDatabaseConnectionConfigById" method="patch" path="/lib/database-connections/{id}" example="UpdateDatabaseConnectionExamplesUpdateMySQLDatabaseConnectionWithSecret" -->
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
        AuthType: components.DatabaseConnectionAuthTypeSecret,
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[int64](15000),
        DatabaseType: components.DatabaseConnectionTypeMysql,
        Description: "Analytics MySQL database",
        ID: "mysql-analytics-db",
        Tags: criblcontrolplanesdkgo.Pointer("analytics,mysql"),
        TextSecret: criblcontrolplanesdkgo.Pointer("mysql-analytics-connection"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatabaseConnectionResponseEnvelope != nil {
        // handle response
    }
}
```
### Example Usage: UpdateDatabaseConnectionExamplesUpdateOracleDatabaseConnectionWithConnectionString

<!-- UsageSnippet language="go" operationID="updateDatabaseConnectionConfigById" method="patch" path="/lib/database-connections/{id}" example="UpdateDatabaseConnectionExamplesUpdateOracleDatabaseConnectionWithConnectionString" -->
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
        AuthType: components.DatabaseConnectionAuthTypeConnectionString,
        ConnectionString: criblcontrolplanesdkgo.Pointer("oracle.example.com:1521/ORCL"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[int64](15000),
        DatabaseType: components.DatabaseConnectionTypeOracle,
        Description: "Oracle ERP database",
        ID: "oracle-erp",
        Password: criblcontrolplanesdkgo.Pointer("yourPassword"),
        Tags: criblcontrolplanesdkgo.Pointer("erp,oracle,finance"),
        User: criblcontrolplanesdkgo.Pointer("yourUsername"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatabaseConnectionResponseEnvelope != nil {
        // handle response
    }
}
```
### Example Usage: UpdateDatabaseConnectionExamplesUpdateOracleDatabaseConnectionWithCredentialsSecrets

<!-- UsageSnippet language="go" operationID="updateDatabaseConnectionConfigById" method="patch" path="/lib/database-connections/{id}" example="UpdateDatabaseConnectionExamplesUpdateOracleDatabaseConnectionWithCredentialsSecrets" -->
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
        AuthType: components.DatabaseConnectionAuthTypeSecrets,
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[int64](15000),
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
    if res.DatabaseConnectionResponseEnvelope != nil {
        // handle response
    }
}
```
### Example Usage: UpdateDatabaseConnectionExamplesUpdateOracleDatabaseConnectionWithSecret

<!-- UsageSnippet language="go" operationID="updateDatabaseConnectionConfigById" method="patch" path="/lib/database-connections/{id}" example="UpdateDatabaseConnectionExamplesUpdateOracleDatabaseConnectionWithSecret" -->
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
        AuthType: components.DatabaseConnectionAuthTypeSecret,
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[int64](20000),
        DatabaseType: components.DatabaseConnectionTypeOracle,
        Description: "Oracle data warehouse",
        ID: "oracle-warehouse",
        Password: criblcontrolplanesdkgo.Pointer("yourPassword"),
        Tags: criblcontrolplanesdkgo.Pointer("warehouse,oracle,reporting"),
        TextSecret: criblcontrolplanesdkgo.Pointer("oracle-warehouse-connection"),
        User: criblcontrolplanesdkgo.Pointer("yourUsername"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatabaseConnectionResponseEnvelope != nil {
        // handle response
    }
}
```
### Example Usage: UpdateDatabaseConnectionExamplesUpdatePostgreSQLDatabaseConnectionWithConnectionString

<!-- UsageSnippet language="go" operationID="updateDatabaseConnectionConfigById" method="patch" path="/lib/database-connections/{id}" example="UpdateDatabaseConnectionExamplesUpdatePostgreSQLDatabaseConnectionWithConnectionString" -->
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
        AuthType: components.DatabaseConnectionAuthTypeConnectionString,
        ConnectionString: criblcontrolplanesdkgo.Pointer("postgresql://yourUsername:yourPassword@postgres.example.com:5432/warehouse?sslmode=require"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[int64](15000),
        DatabaseType: components.DatabaseConnectionTypePostgres,
        Description: "Data warehouse PostgreSQL database",
        ID: "postgres-warehouse",
        Tags: criblcontrolplanesdkgo.Pointer("warehouse,postgres,reporting"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatabaseConnectionResponseEnvelope != nil {
        // handle response
    }
}
```
### Example Usage: UpdateDatabaseConnectionExamplesUpdatePostgreSQLDatabaseConnectionWithSecret

<!-- UsageSnippet language="go" operationID="updateDatabaseConnectionConfigById" method="patch" path="/lib/database-connections/{id}" example="UpdateDatabaseConnectionExamplesUpdatePostgreSQLDatabaseConnectionWithSecret" -->
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
        AuthType: components.DatabaseConnectionAuthTypeSecret,
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[int64](10000),
        DatabaseType: components.DatabaseConnectionTypePostgres,
        Description: "Logs PostgreSQL database",
        ID: "postgres-logs",
        Tags: criblcontrolplanesdkgo.Pointer("logs,postgres"),
        TextSecret: criblcontrolplanesdkgo.Pointer("postgres-logs-connection"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatabaseConnectionResponseEnvelope != nil {
        // handle response
    }
}
```
### Example Usage: UpdateDatabaseConnectionExamplesUpdateSQLServerDatabaseConnectionWithConfigObject

<!-- UsageSnippet language="go" operationID="updateDatabaseConnectionConfigById" method="patch" path="/lib/database-connections/{id}" example="UpdateDatabaseConnectionExamplesUpdateSQLServerDatabaseConnectionWithConfigObject" -->
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
        AuthType: components.DatabaseConnectionAuthTypeConfigObj,
        ConfigObj: criblcontrolplanesdkgo.Pointer("{\"server\":\"sqlserver.example.com\",\"database\":\"Reporting\",\"user\":\"yourUsername\",\"password\":\"yourPassword\",\"options\":{\"encrypt\":true,\"trustServerCertificate\":false,\"connectTimeout\":20000}}"),
        DatabaseType: components.DatabaseConnectionTypeSqlserver,
        Description: "Reporting SQL Server database with custom config",
        ID: "sqlserver-reporting",
        RequestTimeout: criblcontrolplanesdkgo.Pointer[int64](60000),
        Tags: criblcontrolplanesdkgo.Pointer("reporting,sqlserver,analytics"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatabaseConnectionResponseEnvelope != nil {
        // handle response
    }
}
```
### Example Usage: UpdateDatabaseConnectionExamplesUpdateSQLServerDatabaseConnectionWithConnectionString

<!-- UsageSnippet language="go" operationID="updateDatabaseConnectionConfigById" method="patch" path="/lib/database-connections/{id}" example="UpdateDatabaseConnectionExamplesUpdateSQLServerDatabaseConnectionWithConnectionString" -->
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
        AuthType: components.DatabaseConnectionAuthTypeConnectionString,
        ConnectionString: criblcontrolplanesdkgo.Pointer("Server=sqlserver.example.com;Database=ERP;User Id=yourUsername;Password=yourPassword;Encrypt=true"),
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[int64](15000),
        DatabaseType: components.DatabaseConnectionTypeSqlserver,
        Description: "ERP SQL Server database",
        ID: "sqlserver-erp",
        RequestTimeout: criblcontrolplanesdkgo.Pointer[int64](30000),
        Tags: criblcontrolplanesdkgo.Pointer("erp,sqlserver,finance"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatabaseConnectionResponseEnvelope != nil {
        // handle response
    }
}
```
### Example Usage: UpdateDatabaseConnectionExamplesUpdateSQLServerDatabaseConnectionWithSecret

<!-- UsageSnippet language="go" operationID="updateDatabaseConnectionConfigById" method="patch" path="/lib/database-connections/{id}" example="UpdateDatabaseConnectionExamplesUpdateSQLServerDatabaseConnectionWithSecret" -->
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
        AuthType: components.DatabaseConnectionAuthTypeSecret,
        ConnectionTimeout: criblcontrolplanesdkgo.Pointer[int64](15000),
        DatabaseType: components.DatabaseConnectionTypeSqlserver,
        Description: "CRM SQL Server database",
        ID: "sqlserver-crm",
        RequestTimeout: criblcontrolplanesdkgo.Pointer[int64](15000),
        Tags: criblcontrolplanesdkgo.Pointer("crm,sqlserver,sales"),
        TextSecret: criblcontrolplanesdkgo.Pointer("sqlserver-crm-connection"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatabaseConnectionResponseEnvelope != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `id`                                                                                       | `string`                                                                                   | :heavy_check_mark:                                                                         | The <code>id</code> of the Database Connection to update.                                  |
| `databaseConnectionConfig`                                                                 | [components.DatabaseConnectionConfig](../../models/components/databaseconnectionconfig.md) | :heavy_check_mark:                                                                         | DatabaseConnectionConfig object.                                                           |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.UpdateDatabaseConnectionConfigByIDResponse](../../models/operations/updatedatabaseconnectionconfigbyidresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| apierrors.RestAPIJSONError | 400, 404                   | application/json           |
| apierrors.Error            | 500                        | application/json           |
| apierrors.APIError         | 4XX, 5XX                   | \*/\*                      |

## Delete

Delete the specified Database Connection.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteDatabaseConnectionConfigById" method="delete" path="/lib/database-connections/{id}" example="DatabaseConnectionResponseExamplesMySQLDatabaseConnection" -->
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
    if res.DatabaseConnectionResponseEnvelope != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                 | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `ctx`                                                     | [context.Context](https://pkg.go.dev/context#Context)     | :heavy_check_mark:                                        | The context to use for the request.                       |
| `id`                                                      | `string`                                                  | :heavy_check_mark:                                        | The <code>id</code> of the Database Connection to delete. |
| `opts`                                                    | [][operations.Option](../../models/operations/option.md)  | :heavy_minus_sign:                                        | The options for this request.                             |

### Response

**[*operations.DeleteDatabaseConnectionConfigByIDResponse](../../models/operations/deletedatabaseconnectionconfigbyidresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| apierrors.RestAPIJSONError | 404                        | application/json           |
| apierrors.Error            | 500                        | application/json           |
| apierrors.APIError         | 4XX, 5XX                   | \*/\*                      |