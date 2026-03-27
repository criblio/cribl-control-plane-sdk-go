# DatadogSite

Datadog site to which events should be sent

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.DatadogSiteUs

// Open enum: custom values can be created with a direct type cast
custom := components.DatadogSite("custom_value")
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `DatadogSiteUs`     | us                  |
| `DatadogSiteUs3`    | us3                 |
| `DatadogSiteUs5`    | us5                 |
| `DatadogSiteEu`     | eu                  |
| `DatadogSiteFed1`   | fed1                |
| `DatadogSiteAp1`    | ap1                 |
| `DatadogSiteCustom` | custom              |