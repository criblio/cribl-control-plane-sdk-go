# InputResponseInputServicenowTableAuthenticationType

ServiceNow Table API authentication method

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputResponseInputServicenowTableAuthenticationTypeNone

// Open enum: custom values can be created with a direct type cast
custom := components.InputResponseInputServicenowTableAuthenticationType("custom_value")
```


## Values

| Name                                                             | Value                                                            |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `InputResponseInputServicenowTableAuthenticationTypeNone`        | none                                                             |
| `InputResponseInputServicenowTableAuthenticationTypeBasicSecret` | basicSecret                                                      |
| `InputResponseInputServicenowTableAuthenticationTypeOauthSecret` | oauthSecret                                                      |