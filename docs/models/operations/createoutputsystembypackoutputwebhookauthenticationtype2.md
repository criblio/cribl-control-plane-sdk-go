# CreateOutputSystemByPackOutputWebhookAuthenticationType2

Authentication method to use for the HTTP request

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputSystemByPackOutputWebhookAuthenticationType2None

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputSystemByPackOutputWebhookAuthenticationType2("custom_value")
```


## Values

| Name                                                                        | Value                                                                       |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `CreateOutputSystemByPackOutputWebhookAuthenticationType2None`              | none                                                                        |
| `CreateOutputSystemByPackOutputWebhookAuthenticationType2Basic`             | basic                                                                       |
| `CreateOutputSystemByPackOutputWebhookAuthenticationType2CredentialsSecret` | credentialsSecret                                                           |
| `CreateOutputSystemByPackOutputWebhookAuthenticationType2Token`             | token                                                                       |
| `CreateOutputSystemByPackOutputWebhookAuthenticationType2TextSecret`        | textSecret                                                                  |
| `CreateOutputSystemByPackOutputWebhookAuthenticationType2Oauth`             | oauth                                                                       |