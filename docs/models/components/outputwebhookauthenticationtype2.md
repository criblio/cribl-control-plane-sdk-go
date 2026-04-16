# OutputWebhookAuthenticationType2

Authentication method to use for the HTTP request

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputWebhookAuthenticationType2None

// Open enum: custom values can be created with a direct type cast
custom := components.OutputWebhookAuthenticationType2("custom_value")
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `OutputWebhookAuthenticationType2None`              | none                                                |
| `OutputWebhookAuthenticationType2Basic`             | basic                                               |
| `OutputWebhookAuthenticationType2CredentialsSecret` | credentialsSecret                                   |
| `OutputWebhookAuthenticationType2Token`             | token                                               |
| `OutputWebhookAuthenticationType2TextSecret`        | textSecret                                          |
| `OutputWebhookAuthenticationType2Oauth`             | oauth                                               |