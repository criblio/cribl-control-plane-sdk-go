# EndpointConfiguration

Enter the data collection endpoint URL or the individual ID

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.EndpointConfigurationURL

// Open enum: custom values can be created with a direct type cast
custom := components.EndpointConfiguration("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `EndpointConfigurationURL` | url                        |
| `EndpointConfigurationID`  | ID                         |