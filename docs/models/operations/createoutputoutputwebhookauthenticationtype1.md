# CreateOutputOutputWebhookAuthenticationType1

Authentication method to use for the HTTP request

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputOutputWebhookAuthenticationType1None

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputOutputWebhookAuthenticationType1("custom_value")
```


## Values

| Name                                                            | Value                                                           |
| --------------------------------------------------------------- | --------------------------------------------------------------- |
| `CreateOutputOutputWebhookAuthenticationType1None`              | none                                                            |
| `CreateOutputOutputWebhookAuthenticationType1Basic`             | basic                                                           |
| `CreateOutputOutputWebhookAuthenticationType1CredentialsSecret` | credentialsSecret                                               |
| `CreateOutputOutputWebhookAuthenticationType1Token`             | token                                                           |
| `CreateOutputOutputWebhookAuthenticationType1TextSecret`        | textSecret                                                      |
| `CreateOutputOutputWebhookAuthenticationType1Oauth`             | oauth                                                           |