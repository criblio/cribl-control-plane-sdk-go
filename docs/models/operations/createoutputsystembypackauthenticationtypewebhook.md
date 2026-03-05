# CreateOutputSystemByPackAuthenticationTypeWebhook

Authentication method to use for the HTTP request

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputSystemByPackAuthenticationTypeWebhookNone

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputSystemByPackAuthenticationTypeWebhook("custom_value")
```


## Values

| Name                                                                 | Value                                                                |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `CreateOutputSystemByPackAuthenticationTypeWebhookNone`              | none                                                                 |
| `CreateOutputSystemByPackAuthenticationTypeWebhookBasic`             | basic                                                                |
| `CreateOutputSystemByPackAuthenticationTypeWebhookCredentialsSecret` | credentialsSecret                                                    |
| `CreateOutputSystemByPackAuthenticationTypeWebhookToken`             | token                                                                |
| `CreateOutputSystemByPackAuthenticationTypeWebhookTextSecret`        | textSecret                                                           |
| `CreateOutputSystemByPackAuthenticationTypeWebhookOauth`             | oauth                                                                |