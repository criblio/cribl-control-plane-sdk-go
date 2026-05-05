# InputResponseInputOffice365MsgTraceAuthenticationMethod

Select authentication method.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputResponseInputOffice365MsgTraceAuthenticationMethodManual

// Open enum: custom values can be created with a direct type cast
custom := components.InputResponseInputOffice365MsgTraceAuthenticationMethod("custom_value")
```


## Values

| Name                                                                 | Value                                                                |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `InputResponseInputOffice365MsgTraceAuthenticationMethodManual`      | manual                                                               |
| `InputResponseInputOffice365MsgTraceAuthenticationMethodSecret`      | secret                                                               |
| `InputResponseInputOffice365MsgTraceAuthenticationMethodOauth`       | oauth                                                                |
| `InputResponseInputOffice365MsgTraceAuthenticationMethodOauthSecret` | oauthSecret                                                          |
| `InputResponseInputOffice365MsgTraceAuthenticationMethodOauthCert`   | oauthCert                                                            |