# CreateInputSystemByPackGrantType

ServiceNow OAuth grant type used for token requests

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputSystemByPackGrantTypeClientCredentials

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputSystemByPackGrantType("custom_value")
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `CreateInputSystemByPackGrantTypeClientCredentials` | client_credentials                                  |
| `CreateInputSystemByPackGrantTypePassword`          | password                                            |