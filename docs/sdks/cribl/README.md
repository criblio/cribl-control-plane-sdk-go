# System.Settings.Cribl

## Overview

### Available Operations

* [List](#list) - Get system settings
* [Update](#update) - Update system settings

## List

Get the current Cribl system settings.

### Example Usage

<!-- UsageSnippet language="go" operationID="getSystemSettingsConf" method="get" path="/system/settings/conf" example="GetSystemSettingsConfExamplesDefault" -->
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

    res, err := s.System.Settings.Cribl.List(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedSystemSettingsConf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetSystemSettingsConfResponse](../../models/operations/getsystemsettingsconfresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 401                | application/json   |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update the specified Cribl system settings.<br/><br/>This endpoint supports partial updates — provide only the top-level sections (<code>api</code>, <code>workers</code>, <code>tls</code>, <code>proxy</code>, etc.) that you want to change. Omitted top-level sections are preserved unchanged.<br/><br/><b>Important:</b> while top-level sections are optional, nested objects within a section must be complete. For example, if you include <code>api</code>, you must provide its required fields (<code>host</code> and <code>port</code>).

### Example Usage: UpdateSystemSettingsExamplesUpdateApiSettings

<!-- UsageSnippet language="go" operationID="updateSystemSettingsConf" method="patch" path="/system/settings/conf" example="UpdateSystemSettingsExamplesUpdateApiSettings" -->
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

    res, err := s.System.Settings.Cribl.Update(ctx, components.SystemSettingsConfUpdate{
        API: &components.APITypeSystemSettingsConf{
            Disabled: false,
            Host: "0.0.0.0",
            Port: 9000,
            Ssl: &components.SslTypeSystemSettingsConfAPI{
                CertPath: "/opt/cribl/local/cribl/auth/cribl.crt",
                Disabled: false,
                Passphrase: "",
                PrivKeyPath: "/opt/cribl/local/cribl/auth/cribl.key",
            },
        },
        Backups: criblcontrolplanesdkgo.Pointer(components.CreateBackupsSettingsUnionBackupsSettings(
            components.BackupsSettings{
                BackupPersistence: "24h",
                BackupsDirectory: "$CRIBL_STATE_DIR/backups",
            },
        )),
        Pii: criblcontrolplanesdkgo.Pointer(components.CreatePiiSettingsUnionPiiSettings(
            components.PiiSettings{
                EnablePiiDetection: false,
            },
        )),
        Proxy: &components.ProxyTypeSystemSettingsConf{
            UseEnvVars: false,
        },
        Rollback: criblcontrolplanesdkgo.Pointer(components.CreateRollbackSettingsUnionRollbackSettings(
            components.RollbackSettings{
                RollbackEnabled: true,
            },
        )),
        Shutdown: &components.ShutdownTypeSystemSettingsConf{
            DrainTimeout: 10000,
        },
        Sni: criblcontrolplanesdkgo.Pointer(components.CreateSniSettingsUnionSniSettings(
            components.SniSettings{
                DisableSNIRouting: false,
            },
        )),
        System: &components.SystemTypeSystemSettingsConf{
            Intercom: true,
            Upgrade: components.UpgradeOptionsSystemSettingsConfSystemAPI,
        },
        TLS: criblcontrolplanesdkgo.Pointer(components.CreateTLSSettingsUnionTLSSettings(
            components.TLSSettings{
                DefaultCipherList: "DEFAULT",
                DefaultEcdhCurve: "auto",
                MaxVersion: "TLSv1.3",
                MinVersion: "TLSv1.2",
                RejectUnauthorized: true,
            },
        )),
        UpgradeGroupSettings: &components.UpgradeGroupSettings{
            IsRolling: criblcontrolplanesdkgo.Pointer(true),
            Quantity: criblcontrolplanesdkgo.Pointer[int64](100),
            RetryCount: criblcontrolplanesdkgo.Pointer[int64](5),
            RetryDelay: criblcontrolplanesdkgo.Pointer[int64](1000),
        },
        UpgradeSettings: &components.UpgradeSettings{},
        Workers: &components.WorkersTypeSystemSettingsConf{
            Count: 0,
            Memory: 0,
            Minimum: 1,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedSystemSettingsConf != nil {
        // handle response
    }
}
```
### Example Usage: UpdateSystemSettingsResponseExamplesUpdateApiSettings

<!-- UsageSnippet language="go" operationID="updateSystemSettingsConf" method="patch" path="/system/settings/conf" example="UpdateSystemSettingsResponseExamplesUpdateApiSettings" -->
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

    res, err := s.System.Settings.Cribl.Update(ctx, components.SystemSettingsConfUpdate{
        Backups: criblcontrolplanesdkgo.Pointer(components.CreateBackupsSettingsUnionEmptyObject(
            components.EmptyObject{},
        )),
        Pii: criblcontrolplanesdkgo.Pointer(components.CreatePiiSettingsUnionEmptyObject(
            components.EmptyObject{},
        )),
        Rollback: criblcontrolplanesdkgo.Pointer(components.CreateRollbackSettingsUnionEmptyObject(
            components.EmptyObject{},
        )),
        Sni: criblcontrolplanesdkgo.Pointer(components.CreateSniSettingsUnionEmptyObject(
            components.EmptyObject{},
        )),
        TLS: criblcontrolplanesdkgo.Pointer(components.CreateTLSSettingsUnionEmptyObject(
            components.EmptyObject{},
        )),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedSystemSettingsConf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [components.SystemSettingsConfUpdate](../../models/components/systemsettingsconfupdate.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.UpdateSystemSettingsConfResponse](../../models/operations/updatesystemsettingsconfresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 401                | application/json   |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |