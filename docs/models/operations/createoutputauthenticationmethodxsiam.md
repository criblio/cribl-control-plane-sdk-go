# CreateOutputAuthenticationMethodXsiam

Enter a token directly, or provide a secret referencing a token

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputAuthenticationMethodXsiamToken

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputAuthenticationMethodXsiam("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `CreateOutputAuthenticationMethodXsiamToken`  | token                                         |
| `CreateOutputAuthenticationMethodXsiamSecret` | secret                                        |