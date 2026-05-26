# AuthenticationMethodOffice365MsgTrace

Select authentication method.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.AuthenticationMethodOffice365MsgTraceManual

// Open enum: custom values can be created with a direct type cast
custom := components.AuthenticationMethodOffice365MsgTrace("custom_value")
```


## Values

| Name                                               | Value                                              |
| -------------------------------------------------- | -------------------------------------------------- |
| `AuthenticationMethodOffice365MsgTraceManual`      | manual                                             |
| `AuthenticationMethodOffice365MsgTraceSecret`      | secret                                             |
| `AuthenticationMethodOffice365MsgTraceOauth`       | oauth                                              |
| `AuthenticationMethodOffice365MsgTraceOauthSecret` | oauthSecret                                        |
| `AuthenticationMethodOffice365MsgTraceOauthCert`   | oauthCert                                          |