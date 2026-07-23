# ProtocolOptions

Select a transport option for OpenTelemetry

## Example Usage

```go
import (
	"github.com/Cribl-Community/cribl-control-plane-sdk-go/models/components"
)

value := components.ProtocolOptionsGrpc

// Open enum: custom values can be created with a direct type cast
custom := components.ProtocolOptions("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `ProtocolOptionsGrpc` | grpc                  |
| `ProtocolOptionsHTTP` | http                  |