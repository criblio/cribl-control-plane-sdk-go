# CreateOutputSystemByPackEndpointType

Select the type of Dynatrace endpoint configured

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputSystemByPackEndpointTypeSaas

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputSystemByPackEndpointType("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `CreateOutputSystemByPackEndpointTypeSaas` | saas                                       |
| `CreateOutputSystemByPackEndpointTypeAg`   | ag                                         |