# System.Settings.Cribl

## Overview

### Available Operations

* [List](#list) - Get Cribl system settings
* [Update](#update) - Update system settings

## List

Get Cribl system settings.

### Example Usage

<!-- UsageSnippet language="go" operationID="getSystemSettingsConf" method="get" path="/system/settings/conf" -->
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
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update Cribl system settings.

### Example Usage

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

    res, err := s.System.Settings.Cribl.Update(ctx, components.SystemSettingsConf{
        API: components.APITypeSystemSettingsConf{
            Disabled: false,
            Host: "0.0.0.0",
            Port: 9000.0,
            Ssl: &components.SslTypeSystemSettingsConfAPI{
                CertPath: "/opt/cribl/local/cribl/auth/cribl.crt",
                Disabled: false,
                Passphrase: "",
                PrivKeyPath: "/opt/cribl/local/cribl/auth/cribl.key",
            },
        },
        Backups: components.CreateBackupsSettingsUnionBackupsSettings1(
            components.BackupsSettings1{
                BackupPersistence: "24h",
                BackupsDirectory: "$CRIBL_STATE_DIR/backups",
            },
        ),
        Pii: components.CreatePiiSettingsUnionPiiSettings1(
            components.PiiSettings1{
                EnablePiiDetection: false,
            },
        ),
        Proxy: components.ProxyTypeSystemSettingsConf{
            UseEnvVars: false,
        },
        Rollback: components.CreateRollbackSettingsUnionRollbackSettings1(
            components.RollbackSettings1{
                RollbackEnabled: true,
            },
        ),
        Shutdown: components.ShutdownTypeSystemSettingsConf{
            DrainTimeout: 10000.0,
        },
        Sni: components.CreateSniSettingsUnionSniSettings1(
            components.SniSettings1{
                DisableSNIRouting: false,
            },
        ),
        System: components.SystemTypeSystemSettingsConf{
            Intercom: true,
            Upgrade: components.UpgradeOptionsSystemSettingsConfSystemAPI,
        },
        TLS: components.CreateTLSSettingsUnionTLSSettings1(
            components.TLSSettings1{
                DefaultCipherList: "DEFAULT",
                DefaultEcdhCurve: "auto",
                MaxVersion: "TLSv1.3",
                MinVersion: "TLSv1.2",
                RejectUnauthorized: true,
            },
        ),
        UpgradeGroupSettings: components.CreateUpgradeGroupSettingsUnionUpgradeGroupSettings1(
            components.UpgradeGroupSettings1{
                IsRolling: criblcontrolplanesdkgo.Pointer(true),
                Quantity: criblcontrolplanesdkgo.Pointer[float64](100.0),
                RetryCount: criblcontrolplanesdkgo.Pointer[float64](5.0),
                RetryDelay: criblcontrolplanesdkgo.Pointer[float64](1000.0),
            },
        ),
        UpgradeSettings: components.UpgradeSettings{},
        Workers: components.WorkersTypeSystemSettingsConf{
            Count: 0.0,
            Memory: 0.0,
            Minimum: 1.0,
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

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.SystemSettingsConf](../../models/components/systemsettingsconf.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.UpdateSystemSettingsConfResponse](../../models/operations/updatesystemsettingsconfresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |