# Cribl
(*System.Settings.Cribl*)

## Overview

### Available Operations

* [List](#list) - Get Cribl system settings
* [Update](#update) - Update Cribl system settings

## List

Get Cribl system settings

### Example Usage

<!-- UsageSnippet language="go" operationID="getSystemSettingsConf" method="get" path="/system/settings/conf" -->
```go
package main

import(
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.String(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.System.Settings.Cribl.List(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
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

Update Cribl system settings

### Example Usage

<!-- UsageSnippet language="go" operationID="updateSystemSettingsConf" method="patch" path="/system/settings/conf" -->
```go
package main

import(
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.String(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.System.Settings.Cribl.Update(ctx, components.SystemSettingsConf{
        API: components.SystemSettingsConfAPI{
            BaseURL: criblcontrolplanesdkgo.String("https://both-draw.com/"),
            DisableAPICache: criblcontrolplanesdkgo.Bool(true),
            Disabled: false,
            Headers: &components.SystemSettingsConfHeaders{},
            Host: "meaty-spring.biz",
            IdleSessionTTL: criblcontrolplanesdkgo.Float64(89.27),
            ListenOnPort: criblcontrolplanesdkgo.Bool(true),
            LoginRateLimit: criblcontrolplanesdkgo.String("<value>"),
            Port: 2424.38,
            Protocol: "<value>",
            Scripts: criblcontrolplanesdkgo.Bool(true),
            SensitiveFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Ssl: components.SystemSettingsConfSsl{
                CaPath: criblcontrolplanesdkgo.String("<value>"),
                CertPath: "<value>",
                Disabled: true,
                Passphrase: "<value>",
                PrivKeyPath: "<value>",
            },
            SsoRateLimit: criblcontrolplanesdkgo.String("<value>"),
            WorkerRemoteAccess: true,
        },
        Backups: components.SystemSettingsConfBackups{
            BackupPersistence: "<value>",
            BackupsDirectory: "<value>",
        },
        CustomLogo: components.SystemSettingsConfCustomLogo{
            Enabled: false,
            LogoDescription: "<value>",
            LogoImage: "<value>",
        },
        Pii: components.SystemSettingsConfPii{
            EnablePiiDetection: false,
        },
        Proxy: components.SystemSettingsConfProxy{
            UseEnvVars: true,
        },
        Rollback: components.SystemSettingsConfRollback{
            RollbackEnabled: false,
            RollbackRetries: criblcontrolplanesdkgo.Float64(3174.73),
            RollbackTimeout: criblcontrolplanesdkgo.Float64(1506.54),
        },
        Shutdown: components.SystemSettingsConfShutdown{
            DrainTimeout: 3723.75,
        },
        Sni: components.SystemSettingsConfSni{
            DisableSNIRouting: false,
        },
        Sockets: &components.SystemSettingsConfSockets{
            Directory: criblcontrolplanesdkgo.String("/usr/ports"),
        },
        System: components.SystemSettingsConfSystem{
            Intercom: false,
            Upgrade: components.SystemSettingsConfUpgradeAPI,
        },
        TLS: components.SystemSettingsConfTLS{
            DefaultCipherList: "<value>",
            DefaultEcdhCurve: "<value>",
            MaxVersion: "<value>",
            MinVersion: "<value>",
            RejectUnauthorized: true,
        },
        UpgradeGroupSettings: components.UpgradeGroupSettings{
            IsRolling: criblcontrolplanesdkgo.Bool(false),
            Quantity: criblcontrolplanesdkgo.Float64(7915.07),
            RetryCount: criblcontrolplanesdkgo.Float64(4414.66),
            RetryDelay: criblcontrolplanesdkgo.Float64(4374.4),
        },
        UpgradeSettings: components.UpgradeSettings{
            AutomaticUpgradeCheckPeriod: criblcontrolplanesdkgo.String("<value>"),
            DisableAutomaticUpgrade: false,
            EnableLegacyEdgeUpgrade: false,
            PackageUrls: []components.UpgradePackageUrls{
                components.UpgradePackageUrls{
                    PackageHashURL: criblcontrolplanesdkgo.String("https://thrifty-teammate.net/"),
                    PackageURL: "https://skeletal-dwell.info/",
                },
            },
            UpgradeSource: "<value>",
        },
        Workers: components.SystemSettingsConfWorkers{
            Count: 2124.14,
            EnableHeapSnapshots: criblcontrolplanesdkgo.Bool(true),
            LoadThrottlePerc: criblcontrolplanesdkgo.Float64(2538.71),
            Memory: 20.53,
            Minimum: 6157.83,
            StartupMaxConns: criblcontrolplanesdkgo.Float64(4731.29),
            StartupThrottleTimeout: criblcontrolplanesdkgo.Float64(1613.48),
            V8SingleThread: criblcontrolplanesdkgo.Bool(true),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
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