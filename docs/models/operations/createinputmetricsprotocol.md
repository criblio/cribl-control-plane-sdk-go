# CreateInputMetricsProtocol

Protocol to use when collecting metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputMetricsProtocolHTTP

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputMetricsProtocol("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `CreateInputMetricsProtocolHTTP`  | http                              |
| `CreateInputMetricsProtocolHTTPS` | https                             |