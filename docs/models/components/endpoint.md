# Endpoint

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.EndpointCloud

// Open enum: custom values can be created with a direct type cast
custom := components.Endpoint("custom_value")
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `EndpointCloud`      | cloud                |
| `EndpointActiveGate` | activeGate           |
| `EndpointManual`     | manual               |