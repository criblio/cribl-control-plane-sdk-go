# InputResponseGrantType

ServiceNow OAuth grant type used for token requests

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputResponseGrantTypeClientCredentials

// Open enum: custom values can be created with a direct type cast
custom := components.InputResponseGrantType("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `InputResponseGrantTypeClientCredentials` | client_credentials                        |
| `InputResponseGrantTypePassword`          | password                                  |