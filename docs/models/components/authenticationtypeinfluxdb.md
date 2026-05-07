# AuthenticationTypeInfluxdb

InfluxDB authentication type

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.AuthenticationTypeInfluxdbNone

// Open enum: custom values can be created with a direct type cast
custom := components.AuthenticationTypeInfluxdb("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `AuthenticationTypeInfluxdbNone`              | none                                          |
| `AuthenticationTypeInfluxdbBasic`             | basic                                         |
| `AuthenticationTypeInfluxdbCredentialsSecret` | credentialsSecret                             |
| `AuthenticationTypeInfluxdbToken`             | token                                         |
| `AuthenticationTypeInfluxdbTextSecret`        | textSecret                                    |