# InputSplunkSearchAuthenticationType

Splunk Search authentication type

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputSplunkSearchAuthenticationTypeNone

// Open enum: custom values can be created with a direct type cast
custom := components.InputSplunkSearchAuthenticationType("custom_value")
```


## Values

| Name                                                   | Value                                                  |
| ------------------------------------------------------ | ------------------------------------------------------ |
| `InputSplunkSearchAuthenticationTypeNone`              | none                                                   |
| `InputSplunkSearchAuthenticationTypeBasic`             | basic                                                  |
| `InputSplunkSearchAuthenticationTypeCredentialsSecret` | credentialsSecret                                      |
| `InputSplunkSearchAuthenticationTypeToken`             | token                                                  |
| `InputSplunkSearchAuthenticationTypeTextSecret`        | textSecret                                             |