# CreateOutputSystemByPackAuthenticationMethodAzureLogs

Enter workspace ID and workspace key directly, or select a stored secret

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputSystemByPackAuthenticationMethodAzureLogsManual

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputSystemByPackAuthenticationMethodAzureLogs("custom_value")
```


## Values

| Name                                                          | Value                                                         |
| ------------------------------------------------------------- | ------------------------------------------------------------- |
| `CreateOutputSystemByPackAuthenticationMethodAzureLogsManual` | manual                                                        |
| `CreateOutputSystemByPackAuthenticationMethodAzureLogsSecret` | secret                                                        |