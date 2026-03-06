# AuthenticationMethodOptionsManualOauth

Select authentication method.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.AuthenticationMethodOptionsManualOauthManual

// Open enum: custom values can be created with a direct type cast
custom := components.AuthenticationMethodOptionsManualOauth("custom_value")
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `AuthenticationMethodOptionsManualOauthManual`      | manual                                              |
| `AuthenticationMethodOptionsManualOauthSecret`      | secret                                              |
| `AuthenticationMethodOptionsManualOauthOauth`       | oauth                                               |
| `AuthenticationMethodOptionsManualOauthOauthSecret` | oauthSecret                                         |
| `AuthenticationMethodOptionsManualOauthOauthCert`   | oauthCert                                           |