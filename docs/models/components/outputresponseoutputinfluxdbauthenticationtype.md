# OutputResponseOutputInfluxdbAuthenticationType

InfluxDB authentication type

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputResponseOutputInfluxdbAuthenticationTypeNone

// Open enum: custom values can be created with a direct type cast
custom := components.OutputResponseOutputInfluxdbAuthenticationType("custom_value")
```


## Values

| Name                                                              | Value                                                             |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `OutputResponseOutputInfluxdbAuthenticationTypeNone`              | none                                                              |
| `OutputResponseOutputInfluxdbAuthenticationTypeBasic`             | basic                                                             |
| `OutputResponseOutputInfluxdbAuthenticationTypeCredentialsSecret` | credentialsSecret                                                 |
| `OutputResponseOutputInfluxdbAuthenticationTypeToken`             | token                                                             |
| `OutputResponseOutputInfluxdbAuthenticationTypeTextSecret`        | textSecret                                                        |