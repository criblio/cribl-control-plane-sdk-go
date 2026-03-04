# CreateOutputEndpointType

Select the type of Dynatrace endpoint configured

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputEndpointTypeSaas

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputEndpointType("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `CreateOutputEndpointTypeSaas` | saas                           |
| `CreateOutputEndpointTypeAg`   | ag                             |