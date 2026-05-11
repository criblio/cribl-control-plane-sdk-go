# OutputResponseOutputWebhookAuthenticationType1

Authentication method to use for the HTTP request

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputResponseOutputWebhookAuthenticationType1None

// Open enum: custom values can be created with a direct type cast
custom := components.OutputResponseOutputWebhookAuthenticationType1("custom_value")
```


## Values

| Name                                                              | Value                                                             |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `OutputResponseOutputWebhookAuthenticationType1None`              | none                                                              |
| `OutputResponseOutputWebhookAuthenticationType1Basic`             | basic                                                             |
| `OutputResponseOutputWebhookAuthenticationType1CredentialsSecret` | credentialsSecret                                                 |
| `OutputResponseOutputWebhookAuthenticationType1Token`             | token                                                             |
| `OutputResponseOutputWebhookAuthenticationType1TextSecret`        | textSecret                                                        |
| `OutputResponseOutputWebhookAuthenticationType1Oauth`             | oauth                                                             |