# DestinationProtocolOptions

Protocol to use when communicating with the destination.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.DestinationProtocolOptionsUDP

// Open enum: custom values can be created with a direct type cast
custom := components.DestinationProtocolOptions("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `DestinationProtocolOptionsUDP` | udp                             |
| `DestinationProtocolOptionsTCP` | tcp                             |