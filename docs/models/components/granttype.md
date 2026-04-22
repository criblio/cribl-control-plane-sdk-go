# GrantType

ServiceNow OAuth grant type used for token requests

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.GrantTypeClientCredentials

// Open enum: custom values can be created with a direct type cast
custom := components.GrantType("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `GrantTypeClientCredentials` | client_credentials           |
| `GrantTypePassword`          | password                     |