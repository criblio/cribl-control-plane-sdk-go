# CreateOutputOutputWebhookAuthenticationType2

Authentication method to use for the HTTP request

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputOutputWebhookAuthenticationType2None

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputOutputWebhookAuthenticationType2("custom_value")
```


## Values

| Name                                                            | Value                                                           |
| --------------------------------------------------------------- | --------------------------------------------------------------- |
| `CreateOutputOutputWebhookAuthenticationType2None`              | none                                                            |
| `CreateOutputOutputWebhookAuthenticationType2Basic`             | basic                                                           |
| `CreateOutputOutputWebhookAuthenticationType2CredentialsSecret` | credentialsSecret                                               |
| `CreateOutputOutputWebhookAuthenticationType2Token`             | token                                                           |
| `CreateOutputOutputWebhookAuthenticationType2TextSecret`        | textSecret                                                      |
| `CreateOutputOutputWebhookAuthenticationType2Oauth`             | oauth                                                           |