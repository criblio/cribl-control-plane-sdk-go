# InputResponseInputSplunkSearchAuthenticationType

Splunk Search authentication type

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputResponseInputSplunkSearchAuthenticationTypeNone

// Open enum: custom values can be created with a direct type cast
custom := components.InputResponseInputSplunkSearchAuthenticationType("custom_value")
```


## Values

| Name                                                                | Value                                                               |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `InputResponseInputSplunkSearchAuthenticationTypeNone`              | none                                                                |
| `InputResponseInputSplunkSearchAuthenticationTypeBasic`             | basic                                                               |
| `InputResponseInputSplunkSearchAuthenticationTypeCredentialsSecret` | credentialsSecret                                                   |
| `InputResponseInputSplunkSearchAuthenticationTypeToken`             | token                                                               |
| `InputResponseInputSplunkSearchAuthenticationTypeTextSecret`        | textSecret                                                          |