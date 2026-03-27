# InputOpenTelemetryProtocol

Select whether to leverage gRPC or HTTP for OpenTelemetry

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputOpenTelemetryProtocolGrpc

// Open enum: custom values can be created with a direct type cast
custom := components.InputOpenTelemetryProtocol("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `InputOpenTelemetryProtocolGrpc` | grpc                             |
| `InputOpenTelemetryProtocolHTTP` | http                             |