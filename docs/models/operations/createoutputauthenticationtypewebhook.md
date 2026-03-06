# CreateOutputAuthenticationTypeWebhook

Authentication method to use for the HTTP request

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputAuthenticationTypeWebhookNone

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputAuthenticationTypeWebhook("custom_value")
```


## Values

| Name                                                     | Value                                                    |
| -------------------------------------------------------- | -------------------------------------------------------- |
| `CreateOutputAuthenticationTypeWebhookNone`              | none                                                     |
| `CreateOutputAuthenticationTypeWebhookBasic`             | basic                                                    |
| `CreateOutputAuthenticationTypeWebhookCredentialsSecret` | credentialsSecret                                        |
| `CreateOutputAuthenticationTypeWebhookToken`             | token                                                    |
| `CreateOutputAuthenticationTypeWebhookTextSecret`        | textSecret                                               |
| `CreateOutputAuthenticationTypeWebhookOauth`             | oauth                                                    |