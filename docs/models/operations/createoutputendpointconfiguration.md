# CreateOutputEndpointConfiguration

Enter the data collection endpoint URL or the individual ID

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputEndpointConfigurationURL

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputEndpointConfiguration("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `CreateOutputEndpointConfigurationURL` | url                                    |
| `CreateOutputEndpointConfigurationID`  | ID                                     |