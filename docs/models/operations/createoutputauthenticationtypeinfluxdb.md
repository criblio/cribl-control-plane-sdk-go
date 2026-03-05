# CreateOutputAuthenticationTypeInfluxdb

InfluxDB authentication type

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputAuthenticationTypeInfluxdbNone

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputAuthenticationTypeInfluxdb("custom_value")
```


## Values

| Name                                                      | Value                                                     |
| --------------------------------------------------------- | --------------------------------------------------------- |
| `CreateOutputAuthenticationTypeInfluxdbNone`              | none                                                      |
| `CreateOutputAuthenticationTypeInfluxdbBasic`             | basic                                                     |
| `CreateOutputAuthenticationTypeInfluxdbCredentialsSecret` | credentialsSecret                                         |
| `CreateOutputAuthenticationTypeInfluxdbToken`             | token                                                     |
| `CreateOutputAuthenticationTypeInfluxdbTextSecret`        | textSecret                                                |