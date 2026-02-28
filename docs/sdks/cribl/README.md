# System.Settings.Cribl

## Overview

### Available Operations

* [List](#list) - Get Cribl system settings
* [Update](#update) - Update Cribl system settings

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

<!-- UsageSnippet language="go" operationID="updateSystemSettingsConf" method="patch" path="/system/settings/conf" -->
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
            Host: "meaty-spring.biz",
            Port: 2424.38,
        },
        Backups: components.CreateBackupsSettingsUnionBackupsSettings1(
            components.BackupsSettings1{
                BackupPersistence: "<value>",
                BackupsDirectory: "<value>",
            },
        ),
        Pii: components.CreatePiiSettingsUnionPiiSettings1(
            components.PiiSettings1{
                EnablePiiDetection: false,
            },
        ),
        Proxy: components.ProxyTypeSystemSettingsConf{
            UseEnvVars: true,
        },
        Rollback: components.CreateRollbackSettingsUnionRollbackSettings1(
            components.RollbackSettings1{
                RollbackEnabled: false,
                RollbackRetries: criblcontrolplanesdkgo.Pointer[float64](3174.73),
                RollbackTimeout: criblcontrolplanesdkgo.Pointer[float64](1506.54),
            },
        ),
        Shutdown: components.ShutdownTypeSystemSettingsConf{
            DrainTimeout: 3723.75,
        },
        Sni: components.CreateSniSettingsUnionSniSettings1(
            components.SniSettings1{
                DisableSNIRouting: false,
            },
        ),
        System: components.SystemTypeSystemSettingsConf{
            Intercom: false,
            Upgrade: components.UpgradeOptionsSystemSettingsConfSystemAPI,
        },
        TLS: components.CreateTLSSettingsUnionTLSSettings1(
            components.TLSSettings1{
                DefaultCipherList: "<value>",
                DefaultEcdhCurve: "<value>",
                MaxVersion: "<value>",
                MinVersion: "<value>",
                RejectUnauthorized: true,
            },
        ),
        UpgradeGroupSettings: components.CreateUpgradeGroupSettingsUnionUpgradeGroupSettings1(
            components.UpgradeGroupSettings1{
                IsRolling: criblcontrolplanesdkgo.Pointer(false),
                Quantity: criblcontrolplanesdkgo.Pointer[float64](7915.07),
                RetryCount: criblcontrolplanesdkgo.Pointer[float64](4414.66),
                RetryDelay: criblcontrolplanesdkgo.Pointer[float64](4374.4),
            },
        ),
        UpgradeSettings: components.UpgradeSettings{},
        Workers: components.WorkersTypeSystemSettingsConf{
            Count: 2124.14,
            Memory: 20.53,
            Minimum: 6157.83,
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