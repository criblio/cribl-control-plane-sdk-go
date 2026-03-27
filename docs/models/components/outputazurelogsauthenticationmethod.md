# OutputAzureLogsAuthenticationMethod

Enter workspace ID and workspace key directly, or select a stored secret

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputAzureLogsAuthenticationMethodManual

// Open enum: custom values can be created with a direct type cast
custom := components.OutputAzureLogsAuthenticationMethod("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `OutputAzureLogsAuthenticationMethodManual` | manual                                      |
| `OutputAzureLogsAuthenticationMethodSecret` | secret                                      |