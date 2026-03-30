# AuthenticationTypeOptions1

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.AuthenticationTypeOptions1None

// Open enum: custom values can be created with a direct type cast
custom := components.AuthenticationTypeOptions1("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `AuthenticationTypeOptions1None`               | none                                           |
| `AuthenticationTypeOptions1Basic`              | basic                                          |
| `AuthenticationTypeOptions1CredentialsSecret`  | credentialsSecret                              |
| `AuthenticationTypeOptions1SslUserCertificate` | sslUserCertificate                             |