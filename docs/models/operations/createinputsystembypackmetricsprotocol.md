# CreateInputSystemByPackMetricsProtocol

Protocol to use when collecting metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputSystemByPackMetricsProtocolHTTP

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputSystemByPackMetricsProtocol("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `CreateInputSystemByPackMetricsProtocolHTTP`  | http                                          |
| `CreateInputSystemByPackMetricsProtocolHTTPS` | https                                         |