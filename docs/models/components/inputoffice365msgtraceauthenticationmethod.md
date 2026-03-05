# InputOffice365MsgTraceAuthenticationMethod

Select authentication method.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputOffice365MsgTraceAuthenticationMethodManual

// Open enum: custom values can be created with a direct type cast
custom := components.InputOffice365MsgTraceAuthenticationMethod("custom_value")
```


## Values

| Name                                                    | Value                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `InputOffice365MsgTraceAuthenticationMethodManual`      | manual                                                  |
| `InputOffice365MsgTraceAuthenticationMethodSecret`      | secret                                                  |
| `InputOffice365MsgTraceAuthenticationMethodOauth`       | oauth                                                   |
| `InputOffice365MsgTraceAuthenticationMethodOauthSecret` | oauthSecret                                             |
| `InputOffice365MsgTraceAuthenticationMethodOauthCert`   | oauthCert                                               |