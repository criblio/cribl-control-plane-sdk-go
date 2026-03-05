# OutputWebhookAuthenticationType

Authentication method to use for the HTTP request

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputWebhookAuthenticationTypeNone

// Open enum: custom values can be created with a direct type cast
custom := components.OutputWebhookAuthenticationType("custom_value")
```


## Values

| Name                                               | Value                                              |
| -------------------------------------------------- | -------------------------------------------------- |
| `OutputWebhookAuthenticationTypeNone`              | none                                               |
| `OutputWebhookAuthenticationTypeBasic`             | basic                                              |
| `OutputWebhookAuthenticationTypeCredentialsSecret` | credentialsSecret                                  |
| `OutputWebhookAuthenticationTypeToken`             | token                                              |
| `OutputWebhookAuthenticationTypeTextSecret`        | textSecret                                         |
| `OutputWebhookAuthenticationTypeOauth`             | oauth                                              |