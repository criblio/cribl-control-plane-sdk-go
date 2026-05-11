# InputServicenowTableGrantType

ServiceNow OAuth grant type used for token requests

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputServicenowTableGrantTypeClientCredentials

// Open enum: custom values can be created with a direct type cast
custom := components.InputServicenowTableGrantType("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `InputServicenowTableGrantTypeClientCredentials` | client_credentials                               |
| `InputServicenowTableGrantTypePassword`          | password                                         |