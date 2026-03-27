# ConnectionProtocol

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.ConnectionProtocolTCP

// Open enum: custom values can be created with a direct type cast
custom := components.ConnectionProtocol("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `ConnectionProtocolTCP`   | tcp                       |
| `ConnectionProtocolTLS`   | tls                       |
| `ConnectionProtocolHttp2` | http2                     |