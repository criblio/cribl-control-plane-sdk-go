# AuthenticationMethodOptions2

Select authentication method.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.AuthenticationMethodOptions2Manual

// Open enum: custom values can be created with a direct type cast
custom := components.AuthenticationMethodOptions2("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `AuthenticationMethodOptions2Manual`      | manual                                    |
| `AuthenticationMethodOptions2Secret`      | secret                                    |
| `AuthenticationMethodOptions2Oauth`       | oauth                                     |
| `AuthenticationMethodOptions2OauthSecret` | oauthSecret                               |
| `AuthenticationMethodOptions2OauthCert`   | oauthCert                                 |