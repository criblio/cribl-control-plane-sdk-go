# OutputInfluxdbAuthenticationType

InfluxDB authentication type

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputInfluxdbAuthenticationTypeNone

// Open enum: custom values can be created with a direct type cast
custom := components.OutputInfluxdbAuthenticationType("custom_value")
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `OutputInfluxdbAuthenticationTypeNone`              | none                                                |
| `OutputInfluxdbAuthenticationTypeBasic`             | basic                                               |
| `OutputInfluxdbAuthenticationTypeCredentialsSecret` | credentialsSecret                                   |
| `OutputInfluxdbAuthenticationTypeToken`             | token                                               |
| `OutputInfluxdbAuthenticationTypeTextSecret`        | textSecret                                          |