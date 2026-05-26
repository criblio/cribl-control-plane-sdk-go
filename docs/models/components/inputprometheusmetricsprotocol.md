# InputPrometheusMetricsProtocol

Protocol to use when collecting metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputPrometheusMetricsProtocolHTTP

// Open enum: custom values can be created with a direct type cast
custom := components.InputPrometheusMetricsProtocol("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `InputPrometheusMetricsProtocolHTTP`  | http                                  |
| `InputPrometheusMetricsProtocolHTTPS` | https                                 |