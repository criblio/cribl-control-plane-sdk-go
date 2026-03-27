# EndpointType

Select the type of Dynatrace endpoint configured

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.EndpointTypeSaas

// Open enum: custom values can be created with a direct type cast
custom := components.EndpointType("custom_value")
```


## Values

| Name               | Value              |
| ------------------ | ------------------ |
| `EndpointTypeSaas` | saas               |
| `EndpointTypeAg`   | ag                 |