# CreateOutputSystemByPackOutputWebhookAuthenticationType1

Authentication method to use for the HTTP request

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputSystemByPackOutputWebhookAuthenticationType1None

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputSystemByPackOutputWebhookAuthenticationType1("custom_value")
```


## Values

| Name                                                                        | Value                                                                       |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `CreateOutputSystemByPackOutputWebhookAuthenticationType1None`              | none                                                                        |
| `CreateOutputSystemByPackOutputWebhookAuthenticationType1Basic`             | basic                                                                       |
| `CreateOutputSystemByPackOutputWebhookAuthenticationType1CredentialsSecret` | credentialsSecret                                                           |
| `CreateOutputSystemByPackOutputWebhookAuthenticationType1Token`             | token                                                                       |
| `CreateOutputSystemByPackOutputWebhookAuthenticationType1TextSecret`        | textSecret                                                                  |
| `CreateOutputSystemByPackOutputWebhookAuthenticationType1Oauth`             | oauth                                                                       |