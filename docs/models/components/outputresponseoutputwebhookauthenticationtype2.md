# OutputResponseOutputWebhookAuthenticationType2

Authentication method to use for the HTTP request

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputResponseOutputWebhookAuthenticationType2None

// Open enum: custom values can be created with a direct type cast
custom := components.OutputResponseOutputWebhookAuthenticationType2("custom_value")
```


## Values

| Name                                                              | Value                                                             |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `OutputResponseOutputWebhookAuthenticationType2None`              | none                                                              |
| `OutputResponseOutputWebhookAuthenticationType2Basic`             | basic                                                             |
| `OutputResponseOutputWebhookAuthenticationType2CredentialsSecret` | credentialsSecret                                                 |
| `OutputResponseOutputWebhookAuthenticationType2Token`             | token                                                             |
| `OutputResponseOutputWebhookAuthenticationType2TextSecret`        | textSecret                                                        |
| `OutputResponseOutputWebhookAuthenticationType2Oauth`             | oauth                                                             |