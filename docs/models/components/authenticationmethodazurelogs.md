# AuthenticationMethodAzureLogs

Enter workspace ID and workspace key directly, or select a stored secret

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.AuthenticationMethodAzureLogsManual

// Open enum: custom values can be created with a direct type cast
custom := components.AuthenticationMethodAzureLogs("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `AuthenticationMethodAzureLogsManual` | manual                                |
| `AuthenticationMethodAzureLogsSecret` | secret                                |