# OutputResponseEndpointType

Select the type of Dynatrace endpoint configured

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputResponseEndpointTypeSaas

// Open enum: custom values can be created with a direct type cast
custom := components.OutputResponseEndpointType("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `OutputResponseEndpointTypeSaas` | saas                             |
| `OutputResponseEndpointTypeAg`   | ag                               |