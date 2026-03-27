# MetricsProtocol

Protocol to use when collecting metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.MetricsProtocolHTTP

// Open enum: custom values can be created with a direct type cast
custom := components.MetricsProtocol("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `MetricsProtocolHTTP`  | http                   |
| `MetricsProtocolHTTPS` | https                  |