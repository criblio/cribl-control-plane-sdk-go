# CreateOutputEndpoint

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputEndpointCloud

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputEndpoint("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `CreateOutputEndpointCloud`      | cloud                            |
| `CreateOutputEndpointActiveGate` | activeGate                       |
| `CreateOutputEndpointManual`     | manual                           |