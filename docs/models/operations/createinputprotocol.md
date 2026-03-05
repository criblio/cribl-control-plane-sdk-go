# CreateInputProtocol

Select whether to leverage gRPC or HTTP for OpenTelemetry

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputProtocolGrpc

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputProtocol("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `CreateInputProtocolGrpc` | grpc                      |
| `CreateInputProtocolHTTP` | http                      |