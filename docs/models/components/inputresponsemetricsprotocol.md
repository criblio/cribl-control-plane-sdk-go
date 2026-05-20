# InputResponseMetricsProtocol

Protocol to use when collecting metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputResponseMetricsProtocolHTTP

// Open enum: custom values can be created with a direct type cast
custom := components.InputResponseMetricsProtocol("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `InputResponseMetricsProtocolHTTP`  | http                                |
| `InputResponseMetricsProtocolHTTPS` | https                               |