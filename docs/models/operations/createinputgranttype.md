# CreateInputGrantType

ServiceNow OAuth grant type used for token requests

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputGrantTypeClientCredentials

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputGrantType("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `CreateInputGrantTypeClientCredentials` | client_credentials                      |
| `CreateInputGrantTypePassword`          | password                                |