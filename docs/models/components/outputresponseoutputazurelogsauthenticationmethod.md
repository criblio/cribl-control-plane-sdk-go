# OutputResponseOutputAzureLogsAuthenticationMethod

Enter workspace ID and workspace key directly, or select a stored secret

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputResponseOutputAzureLogsAuthenticationMethodManual

// Open enum: custom values can be created with a direct type cast
custom := components.OutputResponseOutputAzureLogsAuthenticationMethod("custom_value")
```


## Values

| Name                                                      | Value                                                     |
| --------------------------------------------------------- | --------------------------------------------------------- |
| `OutputResponseOutputAzureLogsAuthenticationMethodManual` | manual                                                    |
| `OutputResponseOutputAzureLogsAuthenticationMethodSecret` | secret                                                    |