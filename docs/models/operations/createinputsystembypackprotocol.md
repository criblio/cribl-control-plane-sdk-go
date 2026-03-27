# CreateInputSystemByPackProtocol

Select whether to leverage gRPC or HTTP for OpenTelemetry

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputSystemByPackProtocolGrpc

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputSystemByPackProtocol("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `CreateInputSystemByPackProtocolGrpc` | grpc                                  |
| `CreateInputSystemByPackProtocolHTTP` | http                                  |