# InputServicenowTableAuthenticationType

ServiceNow Table API authentication method

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputServicenowTableAuthenticationTypeNone

// Open enum: custom values can be created with a direct type cast
custom := components.InputServicenowTableAuthenticationType("custom_value")
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `InputServicenowTableAuthenticationTypeNone`        | none                                                |
| `InputServicenowTableAuthenticationTypeBasic`       | basic                                               |
| `InputServicenowTableAuthenticationTypeBasicSecret` | basicSecret                                         |
| `InputServicenowTableAuthenticationTypeOauth`       | oauth                                               |
| `InputServicenowTableAuthenticationTypeOauthSecret` | oauthSecret                                         |