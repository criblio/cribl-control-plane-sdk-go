# OutputXsiamAuthenticationMethod

Enter a token directly, or provide a secret referencing a token

## Example Usage

```go
import (
	"github.com/Cribl-Community/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputXsiamAuthenticationMethodToken

// Open enum: custom values can be created with a direct type cast
custom := components.OutputXsiamAuthenticationMethod("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `OutputXsiamAuthenticationMethodToken`  | token                                   |
| `OutputXsiamAuthenticationMethodSecret` | secret                                  |