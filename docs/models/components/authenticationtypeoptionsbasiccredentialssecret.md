# AuthenticationTypeOptionsBasicCredentialsSecret

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.AuthenticationTypeOptionsBasicCredentialsSecretNone

// Open enum: custom values can be created with a direct type cast
custom := components.AuthenticationTypeOptionsBasicCredentialsSecret("custom_value")
```


## Values

| Name                                                                | Value                                                               |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `AuthenticationTypeOptionsBasicCredentialsSecretNone`               | none                                                                |
| `AuthenticationTypeOptionsBasicCredentialsSecretBasic`              | basic                                                               |
| `AuthenticationTypeOptionsBasicCredentialsSecretCredentialsSecret`  | credentialsSecret                                                   |
| `AuthenticationTypeOptionsBasicCredentialsSecretSslUserCertificate` | sslUserCertificate                                                  |