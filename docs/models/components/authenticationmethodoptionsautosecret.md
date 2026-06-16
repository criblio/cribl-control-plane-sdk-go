# AuthenticationMethodOptionsAutoSecret

AWS authentication method. Choose Auto to use IAM roles.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.AuthenticationMethodOptionsAutoSecretAuto

// Open enum: custom values can be created with a direct type cast
custom := components.AuthenticationMethodOptionsAutoSecret("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `AuthenticationMethodOptionsAutoSecretAuto`   | auto                                          |
| `AuthenticationMethodOptionsAutoSecretSecret` | secret                                        |