# CreateOutputSystemByPackAuthenticationTypeInfluxdb

InfluxDB authentication type

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputSystemByPackAuthenticationTypeInfluxdbNone

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputSystemByPackAuthenticationTypeInfluxdb("custom_value")
```


## Values

| Name                                                                  | Value                                                                 |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `CreateOutputSystemByPackAuthenticationTypeInfluxdbNone`              | none                                                                  |
| `CreateOutputSystemByPackAuthenticationTypeInfluxdbBasic`             | basic                                                                 |
| `CreateOutputSystemByPackAuthenticationTypeInfluxdbCredentialsSecret` | credentialsSecret                                                     |
| `CreateOutputSystemByPackAuthenticationTypeInfluxdbToken`             | token                                                                 |
| `CreateOutputSystemByPackAuthenticationTypeInfluxdbTextSecret`        | textSecret                                                            |