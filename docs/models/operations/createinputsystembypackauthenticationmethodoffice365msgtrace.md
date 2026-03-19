# CreateInputSystemByPackAuthenticationMethodOffice365MsgTrace

Select authentication method.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputSystemByPackAuthenticationMethodOffice365MsgTraceManual

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputSystemByPackAuthenticationMethodOffice365MsgTrace("custom_value")
```


## Values

| Name                                                                      | Value                                                                     |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `CreateInputSystemByPackAuthenticationMethodOffice365MsgTraceManual`      | manual                                                                    |
| `CreateInputSystemByPackAuthenticationMethodOffice365MsgTraceSecret`      | secret                                                                    |
| `CreateInputSystemByPackAuthenticationMethodOffice365MsgTraceOauth`       | oauth                                                                     |
| `CreateInputSystemByPackAuthenticationMethodOffice365MsgTraceOauthSecret` | oauthSecret                                                               |
| `CreateInputSystemByPackAuthenticationMethodOffice365MsgTraceOauthCert`   | oauthCert                                                                 |