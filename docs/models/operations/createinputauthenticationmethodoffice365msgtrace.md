# CreateInputAuthenticationMethodOffice365MsgTrace

Select authentication method.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputAuthenticationMethodOffice365MsgTraceManual

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputAuthenticationMethodOffice365MsgTrace("custom_value")
```


## Values

| Name                                                          | Value                                                         |
| ------------------------------------------------------------- | ------------------------------------------------------------- |
| `CreateInputAuthenticationMethodOffice365MsgTraceManual`      | manual                                                        |
| `CreateInputAuthenticationMethodOffice365MsgTraceSecret`      | secret                                                        |
| `CreateInputAuthenticationMethodOffice365MsgTraceOauth`       | oauth                                                         |
| `CreateInputAuthenticationMethodOffice365MsgTraceOauthSecret` | oauthSecret                                                   |
| `CreateInputAuthenticationMethodOffice365MsgTraceOauthCert`   | oauthCert                                                     |