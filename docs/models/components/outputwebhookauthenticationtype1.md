# OutputWebhookAuthenticationType1

Authentication method to use for the HTTP request

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputWebhookAuthenticationType1None

// Open enum: custom values can be created with a direct type cast
custom := components.OutputWebhookAuthenticationType1("custom_value")
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `OutputWebhookAuthenticationType1None`              | none                                                |
| `OutputWebhookAuthenticationType1Basic`             | basic                                               |
| `OutputWebhookAuthenticationType1CredentialsSecret` | credentialsSecret                                   |
| `OutputWebhookAuthenticationType1Token`             | token                                               |
| `OutputWebhookAuthenticationType1TextSecret`        | textSecret                                          |
| `OutputWebhookAuthenticationType1Oauth`             | oauth                                               |