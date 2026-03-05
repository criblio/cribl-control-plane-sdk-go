# CreateOutputDatadogSite

Datadog site to which events should be sent

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputDatadogSiteUs

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputDatadogSite("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `CreateOutputDatadogSiteUs`     | us                              |
| `CreateOutputDatadogSiteUs3`    | us3                             |
| `CreateOutputDatadogSiteUs5`    | us5                             |
| `CreateOutputDatadogSiteEu`     | eu                              |
| `CreateOutputDatadogSiteFed1`   | fed1                            |
| `CreateOutputDatadogSiteAp1`    | ap1                             |
| `CreateOutputDatadogSiteCustom` | custom                          |